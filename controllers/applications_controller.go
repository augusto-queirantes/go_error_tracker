package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "go_error_tracker/database"
    "go_error_tracker/repositories"
)

func GetApplications(context *gin.Context) {
    database := database.Connect()
    application_repository := repositories.ApplicationRepository{Client: database}
    applications := application_repository.FindAll()

    context.JSON(http.StatusOK, gin.H{"data": applications})
}

type CreateApplicationInput struct {
    Name string `json:"application_name" binding:"required"`
}

func CreateApplication(context *gin.Context) {
    var input CreateApplicationInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

        return
    }

    database := database.Connect()
    application_repository := repositories.ApplicationRepository{Client: database}

    application := application_repository.Create(input.Name)

    context.JSON(http.StatusOK, gin.H{"data": application})
}
