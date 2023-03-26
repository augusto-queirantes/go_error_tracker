package main

import (
    "go_error_tracker/config"
    "go_error_tracker/database"
)

func init() {
    config.LoadEnvs()
    database.Connect()
    database.Migrate()
}

func main() {
    config.StartServer()
}
