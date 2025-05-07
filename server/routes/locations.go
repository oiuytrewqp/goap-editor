package routes

import (
	"net/http"
	"oiuytrewqp/goap-server/server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/locations", getLocations)
		api.GET("/locations/:id", getLocation)
		api.POST("/locations", createLocation)
		api.PUT("/locations/:id", updateLocation)
		api.DELETE("/locations/:id", deleteLocation)
	}
}

func getLocations(context *gin.Context) {
	allLocations, err := models.GetLocations()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting locations."})
		return
	}

	context.JSON(http.StatusOK, allLocations)
}

func getLocation(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	location, err := models.GetLocation(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting location."})
		return
	}

	context.JSON(http.StatusOK, location)
}

func createLocation(context *gin.Context) {
	var location models.Location
	context.ShouldBindJSON(&location)

	err := models.CreateLocation(location)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving location."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

func updateLocation(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	var location models.Location
	context.ShouldBindJSON(&location)

	err = models.UpdateLocation(id, location)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating location."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

func deleteLocation(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	err = models.DeleteLocation(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting location."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}
