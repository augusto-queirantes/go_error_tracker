package database

import (
    "fmt"
    "log"
    "os"

    "go_error_tracker/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type Database struct {
    Postgres *gorm.DB
}

func Connect() *Database {
    var err error
    var database_client *gorm.DB

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
        os.Getenv("POSTGRES_HOST"),
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DATABASE"),
        os.Getenv("POSTGRES_PORT"),
    )

    database_client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to the Database")
    }

    fmt.Println("Database successfully connected")

    return &Database{
        Postgres: database_client,
    }
}

func Migrate(database_connection *Database) {
    err := database_connection.Postgres.AutoMigrate(&models.Exception{}, &models.Application{})

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Database successfully migrated")
}
