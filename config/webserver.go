package config

import (
    "github.com/gin-gonic/gin"

    "go_error_tracker/controllers"
)

func StartServer() {
    router := gin.Default()

    router.GET("/exceptions", controllers.GetExceptions)
    router.POST("/exceptions", controllers.CreateException)

    router.Run()
}
