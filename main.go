package main

import (
    "go_error_tracker/config"
    "go_error_tracker/database"
)

func init() {
    config.LoadEnvs()

    database_connection := database.Connect()

    database.Migrate(database_connection)
}

func main() {
    config.StartServer()
}
