package config

import (
    "log"
    "net/http"

    controllers "go_error_tracker/controllers"
)

func StartServer() {
    http.HandleFunc("/", controllers.ErrorsController)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
