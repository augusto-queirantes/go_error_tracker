package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go_error_tracker/database"
	"go_error_tracker/repositories"
)

func GetExceptions(context *gin.Context) {
    database := database.Connect()
    exception_repository  := repositories.ExceptionRepository{Client: database}
    exceptions := exception_repository.FindAll()

    context.JSON(http.StatusOK, gin.H{"data": exceptions})
}

type CreateExceptionInput struct {
    Name string `json:"name" binding:"required"`
    StackTrace string `json:"stack_trace" binding:"required"`
    ApplicationName string `json:"application_name" binding:"required"`
}

func CreateException(context *gin.Context) {
    var input CreateExceptionInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

        return
    }

    database := database.Connect()
    exception_repository := repositories.ExceptionRepository{Client: database}
    exception := exception_repository.Create(input.Name, input.StackTrace, input.ApplicationName)

    context.JSON(http.StatusOK, gin.H{"data": exception})
}
