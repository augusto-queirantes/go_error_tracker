package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "go_error_tracker/database"
    "go_error_tracker/models"
)

func GetExceptions(context *gin.Context) {
    var exceptions []models.Exception

    database.Database.Find(&exceptions)

    context.JSON(http.StatusOK, gin.H{"data": exceptions})
}

type CreateExceptionInput struct {
    Name string `json:"exception_name" binding:"required"`
    StackTrace string `json:"exception_stack_trace" binding:"required"`
}

func CreateException(context *gin.Context) {
    var input CreateExceptionInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

        return
    }

    exception := models.Exception{Name: input.Name, StackTrace: input.StackTrace}

    database.Database.Create(&exception)

    context.JSON(http.StatusOK, gin.H{"data": exception})
}
