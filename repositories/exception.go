package repositories

import (
	"go_error_tracker/database"
	"go_error_tracker/models"
)

type ExceptionRepository struct {
    Client *database.Database
}

func (exception_repository ExceptionRepository) Create(
    name string,
    stack_trace string,
    application_name string,
) models.Exception {
    application_repository := ApplicationRepository{Client: exception_repository.Client}
    application := application_repository.FindByName(application_name)

    exception := models.Exception{Name: name, StackTrace: stack_trace, Application: application}

    exception_repository.Client.Postgres.Create(&exception)

    return exception
}

func (exception_repository ExceptionRepository) FindAll() []models.Exception {
    var exceptions []models.Exception

    exception_repository.Client.Postgres.Joins("Application").Find(&exceptions)

    return exceptions
}
