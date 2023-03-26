package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "go_error_tracker/database"
    "go_error_tracker/models"
)

func GetErrors(context *gin.Context) {
    var errors []models.Error

    database.Database.Find(&errors)

    context.JSON(http.StatusOK, gin.H{"data": errors})
}

type CreateErrorInput struct {
    Name string `json:"error_name" binding:"required"`
    StackTrace string `json:"error_stack_trace" binding:"required"`
}

func CreateError(context *gin.Context) {
    var input CreateErrorInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := models.Error{Name: input.Name, StackTrace: input.StackTrace}

    database.Database.Create(&err)

    context.JSON(http.StatusOK, gin.H{"data": err})
}
