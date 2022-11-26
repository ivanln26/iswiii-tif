package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type VoteDB interface {
	Insert(Vote) (Vote, error)
	Get(string) (Vote, error)
	GetAll() ([]Vote, error)
	GetPercentages() ([]VotePercentage, error)
}

type MapVoteDB map[string]Vote

func (m MapVoteDB) Insert(v Vote) (Vote, error) {
	v.Id = uuid.NewString()
	m[v.Id] = v
	log.Printf("map db: vote %+v inserted\n", v)
	return v, nil
}

func (m MapVoteDB) Get(id string) (Vote, error) {
	v, ok := m[id]
	if !ok {
		return v, fmt.Errorf("map db: vote with id %s not found\n", id)
	}
	return v, nil
}

func (m MapVoteDB) GetAll() ([]Vote, error) {
	votes := make([]Vote, 0, len(m))
	for _, v := range m {
		votes = append(votes, v)
	}
	return votes, nil
}

type SQLDB struct {
	DB *sql.DB
}

func (m MapVoteDB) GetPercentages() ([]VotePercentage, error) {
	percentages := make([]VotePercentage, 0, 2)
	if len(m) == 0 {
		return percentages, fmt.Errorf("map db: could not create percentages\n")
	}
	var countA int
	var countB int
	for _, v := range m {
		if v.Choice == 1 {
			countA++
		}
		if v.Choice == 2 {
			countB++
		}
	}
	percentages = append(percentages, VotePercentage{1, float64(countA) / float64(len(m)) * 100.0})
	percentages = append(percentages, VotePercentage{2, float64(countB) / float64(len(m)) * 100.0})
	return percentages, nil
}

func SQLDBConnect(dsn string) *SQLDB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		log.Fatalln("mysql: bad dsn arguments")
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		log.Fatalln("mysql: could not connect to the database")
	}
	log.Println("mysql: connection succeded")
	return &SQLDB{db}
}

func (db *SQLDB) Insert(v Vote) (Vote, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return v, err
	}
	defer tx.Rollback()

	_, err = tx.Exec("SET @id=UUID()")
	if err != nil {
		return v, err
	}

	_, err = tx.Exec("INSERT INTO `vote` VALUES (UUID_TO_BIN(@id), ?)", v.Choice)
	if err != nil {
		return v, err
	}

	var id string
	tx.QueryRow("SELECT @id").Scan(&id)
	v.Id = id

	if err := tx.Commit(); err != nil {
		return v, err
	}
	log.Printf("mysql: vote %+v inserted\n", v)
	return v, nil
}

func (db SQLDB) Get(id string) (Vote, error) {
	var choice int
	err := db.DB.QueryRow("SELECT `choice` FROM `vote` WHERE `id` = UUID_TO_BIN(?)", id).Scan(&choice)
	v := Vote{id, choice}
	if err != nil {
		return v, err
	}
	return v, nil
}

func (db SQLDB) GetAll() ([]Vote, error) {
	votes := make([]Vote, 0)
	rows, err := db.DB.Query("SELECT BIN_TO_UUID(`id`), `choice` FROM `vote`")
	if err != nil {
		return votes, err
	}
	for rows.Next() {
		var id string
		var choice int
		if err := rows.Scan(&id, &choice); err != nil {
			return votes, err
		}
		votes = append(votes, Vote{id, choice})
	}
	return votes, nil
}

func (db SQLDB) GetPercentages() ([]VotePercentage, error) {
	per := make([]VotePercentage, 0, 2)

	tx, err := db.DB.Begin()
	if err != nil {
		return per, err
	}
	defer tx.Rollback()

	_, err = tx.Exec("SET @count = (SELECT COUNT(*) FROM `vote`)")
	if err != nil {
		return per, err
	}

	rows, err := tx.Query("SELECT `choice`, COUNT(*) / @count * 100 AS porcentaje FROM `vote` GROUP BY `choice`")
	if err != nil {
		return per, err
	}

	for rows.Next() {
		var choice int
		var percentage float64
		if err := rows.Scan(&choice, &percentage); err != nil {
			return per, err
		}
		per = append(per, VotePercentage{choice, percentage})
	}

	return per, nil
}

func DBFactory(dsn string) VoteDB {
	if dsn != "" {
		log.Println("mysql: starting connection")
		return SQLDBConnect(dsn)
	}
	log.Println("map db: created")
	return make(MapVoteDB)
}
