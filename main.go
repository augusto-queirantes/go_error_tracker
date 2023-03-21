package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
    Name string `json:"name"`
    StackTrace string `json:"stack_trace"`
}

func ErrorIndex(response_writter http.ResponseWriter, r *http.Request) {
    error := Error {
        Name: "Application Error",
        StackTrace: "You're application error stack_trace",
    }

    json.NewEncoder(response_writter).Encode(error)
}

func main() {
    http.HandleFunc("/", ErrorIndex)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
