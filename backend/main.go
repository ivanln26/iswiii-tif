package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v9"
)

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
	ctx := context.Background()

	r := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	res, err := r.Ping(ctx).Result()
	if err != nil {
		panic("redis not connected")
	}
	fmt.Println(res)
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
	http.ListenAndServe(":8000", nil)
}
