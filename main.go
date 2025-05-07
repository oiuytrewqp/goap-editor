package main

import (
	"net/http"
	"oiuytrewqp/goap-server/server/routes"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./client/dist", true)))

	routes.InitialiseRoutes(router)

	api := router.Group("/api")
	{
		api.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	router.Run(":8080")
}
