package routes

import (
	"net/http"
	"oiuytrewqp/goap-server/server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ActionRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/actions", getActions)
		api.GET("/actions/:id", getAction)
		api.POST("/actions", createAction)
		api.PUT("/actions/:id", updateAction)
		api.DELETE("/actions/:id", deleteAction)
	}
}

func getActions(context *gin.Context) {
	allActions, err := models.GetActions()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting actions."})
		return
	}

	context.JSON(http.StatusOK, allActions)
}

func getAction(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	action, err := models.GetAction(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting action."})
		return
	}

	context.JSON(http.StatusOK, action)
}

func createAction(context *gin.Context) {
	var action models.Action
	context.ShouldBindJSON(&action)

	id, err := models.CreateAction(action)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving action.", "err": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"id": id})
}

func updateAction(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	var action models.Action
	context.ShouldBindJSON(&action)

	err = models.UpdateAction(id, action)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating action."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

func deleteAction(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	err = models.DeleteAction(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting action."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}
