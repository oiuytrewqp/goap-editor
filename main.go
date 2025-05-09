package main

import (
	"oiuytrewqp/goap-server/server/models"
	"oiuytrewqp/goap-server/server/routes"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	models.InitialiseDatabase()

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./client/dist", true)))
	router.Use(static.Serve("/agents", static.LocalFile("./client/dist", true)))
	router.Use(static.Serve("/beliefs", static.LocalFile("./client/dist", true)))
	router.Use(static.Serve("/goals", static.LocalFile("./client/dist", true)))
	router.Use(static.Serve("/actions", static.LocalFile("./client/dist", true)))
	router.Use(static.Serve("/locations", static.LocalFile("./client/dist", true)))

	routes.InitialiseRoutes(router)

	router.Run(":8080")

	defer models.Database.Close()
}
