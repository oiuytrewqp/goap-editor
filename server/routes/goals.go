package routes

import (
	"net/http"
	"oiuytrewqp/goap-server/server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GoalRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/goals", getGoals)
		api.GET("/goals/:id", getGoal)
		api.POST("/goals", createGoal)
		api.PUT("/goals/:id", updateGoal)
		api.DELETE("/goals/:id", deleteGoal)
	}
}

func getGoals(context *gin.Context) {
	allGoals, err := models.GetGoals()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting goals."})
		return
	}

	context.JSON(http.StatusOK, allGoals)
}

func getGoal(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	goal, err := models.GetGoal(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting goal."})
		return
	}

	context.JSON(http.StatusOK, goal)
}

func createGoal(context *gin.Context) {
	var goal models.Goal
	context.ShouldBindJSON(&goal)

	err := models.CreateGoal(goal)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving goal."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

func updateGoal(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	var goal models.Goal
	context.ShouldBindJSON(&goal)

	err = models.UpdateGoal(id, goal)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating goal."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

func deleteGoal(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	err = models.DeleteGoal(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting goal."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}
