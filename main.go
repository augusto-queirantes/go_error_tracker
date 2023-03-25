package main

import (
    "go_error_tracker/config"
)

func init() {
    config.LoadEnvs()
    config.Connect()
    config.Migrate()
}

func main() {
    config.StartServer()
}
