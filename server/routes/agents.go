package routes

import (
	"net/http"
	"oiuytrewqp/goap-server/server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AgentRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/agents", getAgents)
		api.GET("/agents/:id", getAgent)
		api.POST("/agents", createAgent)
		api.PUT("/agents/:id", updateAgent)
		api.DELETE("/agents/:id", deleteAgent)
	}
}

func getAgents(context *gin.Context) {
	allAgents, err := models.GetAgents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting agents."})
		return
	}

	context.JSON(http.StatusOK, allAgents)
}

func getAgent(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	agent, err := models.GetAgent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting agent."})
		return
	}
	context.JSON(http.StatusOK, agent)
}

func createAgent(context *gin.Context) {
	var agent models.Agent
	context.ShouldBindJSON(&agent)

	err := models.CreateAgent(agent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving agent."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

func updateAgent(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	var agent models.Agent
	context.ShouldBindJSON(&agent)

	err = models.UpdateAgent(id, agent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating agent."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

func deleteAgent(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	err = models.DeleteAgent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting agent."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}
