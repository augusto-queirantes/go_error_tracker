package config

import (
	"log"

	"github.com/gin-gonic/gin"

	"go_error_tracker/controllers"
)

func StartServer() {
    router := gin.Default()

    router.GET("/exceptions", controllers.GetExceptions)
    router.POST("/exceptions", controllers.CreateException)
    router.GET("/applications", controllers.GetApplications)
    router.POST("/applications", controllers.CreateApplication)

    err := router.Run()

    if err != nil {
        log.Fatal(err)
    }
}
