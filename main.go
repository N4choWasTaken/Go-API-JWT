package main

import (
	"github.com/gin-gonic/gin"
	"www.github.com/N4choWasTaken/Go-API-JWT/controllers"
	db "www.github.com/N4choWasTaken/Go-API-JWT/database"
	middlewares "www.github.com/N4choWasTaken/Go-API-JWT/middleware"
)

func main() {
	// Initialize Database
	db.Connect("host=localhost user=postgres password=mysecretpassword dbname=postgres port=5455 sslmode=disable")
	db.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":8081")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
