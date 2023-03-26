package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "go_error_tracker/database"
    "go_error_tracker/models"
)

func GetApplications(context *gin.Context) {
    var applications []models.Application

    database.Database.Find(&applications)

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

    application := models.Application{Name: input.Name}

    database.Database.Create(&application)

    context.JSON(http.StatusOK, gin.H{"data": application})
}
