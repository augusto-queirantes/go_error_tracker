package controllers

import (
    "net/http"
    "encoding/json"

    "go_error_tracker/models"
)

func ErrorsController(response_writter http.ResponseWriter, request *http.Request) {
    error := models.Error {
        Name: "Application Error",
        StackTrace: "You're application error stack_trace",
    }

    json.NewEncoder(response_writter).Encode(error)
}
