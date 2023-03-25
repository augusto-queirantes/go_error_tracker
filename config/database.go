package config

import (
    "fmt"
    "log"
    "os"

    "go_error_tracker/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
    var err error

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
        os.Getenv("POSTGRES_HOST"),
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DATABASE"),
        os.Getenv("POSTGRES_PORT"),
    )

    Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to the Database")
    }

    fmt.Println("Database successfully connected")
}

func Migrate() {
    Database.AutoMigrate(&models.Error{})

    fmt.Println("Database successfully migrated")
}
