package controllers

import (
    "net/http"
    "encoding/json"
)

type Error struct {
    Name string `json:"name"`
    StackTrace string `json:"stack_trace"`
}

func ErrorsController(response_writter http.ResponseWriter, r *http.Request) {
    error := Error {
        Name: "Application Error",
        StackTrace: "You're application error stack_trace",
    }

    json.NewEncoder(response_writter).Encode(error)
}
