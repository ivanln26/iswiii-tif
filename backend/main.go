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
	Choice int `json:"choice"`
}

var votes = make([]Vote, 0)

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

func ListVotes(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"application/json"}
	if r.Method != http.MethodGet {
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
	redisHost := GetEnv("REDIS_HOST", "localhost")
	redisPort := GetEnv("REDIS_PORT", "6379")
	log.Printf("Redis URI: %s:%s", redisHost, redisPort)

	ctx := context.Background()

	r := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: "",
		DB:       0,
	})
	res, err := r.Ping(ctx).Result()
	if err != nil {
		panic("redis not connected")
	}
	log.Println(res)
	pubsub := r.Subscribe(ctx, "votes")
	ch := pubsub.Channel()

	go func() {
		for msg := range ch {
			var vote Vote
			err := json.Unmarshal([]byte(msg.Payload), &vote)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%+v\n", vote)
			votes = append(votes, vote)
		}
	}()

	http.HandleFunc("/", Index)
	http.HandleFunc("/votes", ListVotes)
	log.Printf("Application running on port: %s", port)
	http.ListenAndServe(":"+port, nil)
}
