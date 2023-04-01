package repositories

import (
    "go_error_tracker/database"
    "go_error_tracker/models"
)

type ExceptionRepository struct {
    Client *database.Database
}

func (exception_repository ExceptionRepository) Create(name string, stack_trace string) models.Exception {
    exception := models.Exception{Name: name, StackTrace: stack_trace}

    exception_repository.Client.Postgres.Create(exception)

    return exception
}

func (exception_repository ExceptionRepository) FindAll() []models.Exception {
    var exceptions []models.Exception

    exception_repository.Client.Postgres.Find(&exceptions)

    return exceptions
}
