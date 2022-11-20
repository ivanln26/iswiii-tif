package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v9"
)

func GetEnv(key string, def string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return env
}

type Vote struct {
	Id     string `json:"id"`
	Choice int    `json:"choice"`
}

type msg struct {
	Hello string `json:"hello"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"application/json"}
	if r.Method != http.MethodGet {
		return
	}

	m := msg{"world"}
	b, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "{\"detail\":\"could not parse json\"}")
		return
	}

	w.WriteHeader(http.StatusTeapot)
	fmt.Fprintln(w, string(b))
}

type ListVotesHandler struct {
	db VoteDB
}

func (h ListVotesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"application/json"}
	if r.Method != http.MethodGet {
		return
	}

	votes, err := h.db.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	b, err := json.Marshal(votes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "{\"detail\":\"could not parse json\"}")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(b))
}

func main() {
	port := GetEnv("PORT", "8000")
	dbDSN := GetEnv("DATABASE_DSN", "")
	redisHost := GetEnv("REDIS_HOST", "localhost")
	redisPort := GetEnv("REDIS_PORT", "6379")
	redisPassword := GetEnv("REDIS_PASSWORD", "")
	log.Printf("Redis URI: %s:%s", redisHost, redisPort)

	db := DBFactory(dbDSN)

	ctx := context.Background()

	r := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})
	res, err := r.Ping(ctx).Result()
	if err != nil {
		panic("redis: connection could not be established")
	}
	log.Println("redis: connection established", res)
	pubsub := r.Subscribe(ctx, "votes")
	ch := pubsub.Channel()

	go func() {
		for msg := range ch {
			var vote Vote
			err := json.Unmarshal([]byte(msg.Payload), &vote)
			if err != nil {
				log.Println(err)
			}
			log.Printf("redis: vote %+v arrived\n", vote)
			_, err = db.Insert(vote)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	http.HandleFunc("/", Index)
	http.Handle("/votes", ListVotesHandler{db})
	log.Printf("Application running on port: %s", port)
	http.ListenAndServe(":"+port, nil)
}
