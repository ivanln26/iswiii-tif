package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v9"
)

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

	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}
