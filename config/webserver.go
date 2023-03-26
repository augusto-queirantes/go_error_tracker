package config

import (
    "github.com/gin-gonic/gin"

    "go_error_tracker/controllers"
)

func StartServer() {
    router := gin.Default()

    router.GET("/errors", controllers.GetErrors)
    router.POST("/errors", controllers.CreateError)

    router.Run()
}
