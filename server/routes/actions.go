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

	for i, action := range allActions {
		actionPrerequisites, err := models.GetActionPrerequisites(action.ID)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting action prerequisites.", "error": err.Error(), "id": action.ID})
			return
		}

		if actionPrerequisites == nil {
			actionPrerequisites = []int64{}
		}

		allActions[i].Prerequisites = actionPrerequisites

		actionOutcomes, err := models.GetActionOutcomes(action.ID)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting action outcomes.", "error": err.Error(), "id": action.ID})
			return
		}

		if actionOutcomes == nil {
			actionOutcomes = []int64{}
		}

		allActions[i].Outcomes = actionOutcomes
	}

	if allActions == nil {
		allActions = []models.Action{}
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

	actionPrerequisites, err := models.GetActionPrerequisites(action.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting action prerequisites.", "error": err.Error(), "id": action.ID})
		return
	}

	if actionPrerequisites == nil {
		actionPrerequisites = []int64{}
	}

	action.Prerequisites = actionPrerequisites

	actionOutcomes, err := models.GetActionOutcomes(action.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting action outcomes.", "error": err.Error(), "id": action.ID})
		return
	}

	if actionOutcomes == nil {
		actionOutcomes = []int64{}
	}

	action.Outcomes = actionOutcomes

	context.JSON(http.StatusOK, action)
}

func createAction(context *gin.Context) {
	var action models.Action
	context.ShouldBindJSON(&action)

	id, err := models.CreateAction(action)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving action.", "err": err.Error(), "action": action})
		return
	}

	actionPrerequisites := []models.ActionPrerequisite{}
	for _, prerequisiteID := range action.Prerequisites {
		actionPrerequisites = append(actionPrerequisites, models.ActionPrerequisite{ActionID: id, PrerequisiteID: prerequisiteID})
	}

	err = models.AddActionPrerequisites(actionPrerequisites)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving action prerequisites.", "err": err.Error()})
		return
	}

	actionOutcomes := []models.ActionOutcome{}
	for _, outcomeID := range action.Outcomes {
		actionOutcomes = append(actionOutcomes, models.ActionOutcome{ActionID: id, OutcomeID: outcomeID})
	}

	err = models.AddActionOutcomes(actionOutcomes)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving action outcomes.", "err": err.Error()})
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

	actionPrerequisites := []models.ActionPrerequisite{}
	for _, prerequisiteID := range action.Prerequisites {
		actionPrerequisites = append(actionPrerequisites, models.ActionPrerequisite{ActionID: int64(id), PrerequisiteID: prerequisiteID})
	}

	err = models.RemoveActionPrerequisites(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error removing action prerequisites."})
		return
	}

	err = models.AddActionPrerequisites(actionPrerequisites)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving action prerequisites.", "err": err.Error()})
		return
	}

	actionOutcomes := []models.ActionOutcome{}
	for _, outcomeID := range action.Outcomes {
		actionOutcomes = append(actionOutcomes, models.ActionOutcome{ActionID: int64(id), OutcomeID: outcomeID})
	}

	err = models.RemoveActionOutcomes(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error removing action outcomes."})
		return
	}

	err = models.AddActionOutcomes(actionOutcomes)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving action outcomes.", "err": err.Error()})
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
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting action.", "err": err.Error()})
		return
	}

	err = models.RemoveActionPrerequisites(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error removing action prerequisites."})
		return
	}

	err = models.RemoveActionOutcomes(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error removing action outcomes."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}
