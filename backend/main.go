package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}
