package controllers

import (
    "net/http"
    "encoding/json"

    "go_error_tracker/interfaces"
)

func ErrorsController(response_writter http.ResponseWriter, request *http.Request) {
    error := interfaces.Error {
        Name: "Application Error",
        StackTrace: "You're application error stack_trace",
    }

    json.NewEncoder(response_writter).Encode(error)
}
