package repositories

import (
    "go_error_tracker/database"
    "go_error_tracker/models"
)

type ApplicationRepository struct {
    Client *database.Database
}

func (application_repository ApplicationRepository) Create(name string) models.Application {
    application := models.Application{Name: name}

    application_repository.Client.Postgres.Create(&application)

    return application
}

func (application_repository ApplicationRepository) FindAll() []models.Application {
    var applications []models.Application

    application_repository.Client.Postgres.Find(&applications)

    return applications
}
