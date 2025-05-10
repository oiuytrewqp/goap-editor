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

	for i, goal := range allGoals {
		goalBeliefs, err := models.GetGoalBeliefs(goal.ID)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting goal beliefs.", "error": err.Error(), "id": goal.ID})
			return
		}

		if goalBeliefs == nil {
			goalBeliefs = []int64{}
		}

		allGoals[i].Beliefs = goalBeliefs
	}

	if allGoals == nil {
		allGoals = []models.Goal{}
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

	goalBeliefs, err := models.GetGoalBeliefs(goal.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting goal beliefs."})
		return
	}

	if goalBeliefs == nil {
		goalBeliefs = []int64{}
	}

	goal.Beliefs = goalBeliefs

	context.JSON(http.StatusOK, goal)
}

func createGoal(context *gin.Context) {
	var goal models.Goal
	context.ShouldBindJSON(&goal)

	id, err := models.CreateGoal(goal)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving goal."})
		return
	}

	goalBeliefs := []models.GoalBelief{}
	for _, beliefID := range goal.Beliefs {
		goalBeliefs = append(goalBeliefs, models.GoalBelief{GoalID: id, BeliefID: beliefID})
	}

	err = models.AddGoalBeliefs(goalBeliefs)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving goal beliefs."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"id": id})
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

	goalBeliefs := []models.GoalBelief{}
	for _, beliefID := range goal.Beliefs {
		goalBeliefs = append(goalBeliefs, models.GoalBelief{GoalID: int64(id), BeliefID: beliefID})
	}

	err = models.RemoveGoalBeliefs(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error removing goal beliefs."})
		return
	}

	err = models.AddGoalBeliefs(goalBeliefs)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving goal beliefs."})
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

	err = models.RemoveGoalBeliefs(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error removing goal beliefs."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}
