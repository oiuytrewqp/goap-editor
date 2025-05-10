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

	for i, agent := range allAgents {
		agentBeliefs, err := models.GetAgentBeliefs(agent.ID)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting agent beliefs.", "error": err.Error(), "id": agent.ID})
			return
		}

		if agentBeliefs == nil {
			agentBeliefs = []int64{}
		}

		allAgents[i].Beliefs = agentBeliefs

		agentGoals, err := models.GetAgentGoals(agent.ID)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting agent goals.", "error": err.Error(), "id": agent.ID})
			return
		}

		if agentGoals == nil {
			agentGoals = []int64{}
		}

		allAgents[i].Goals = agentGoals

		agentActions, err := models.GetAgentActions(agent.ID)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting agent actions.", "error": err.Error(), "id": agent.ID})
			return
		}

		if agentActions == nil {
			agentActions = []int64{}
		}

		allAgents[i].Actions = agentActions
	}

	if allAgents == nil {
		allAgents = []models.Agent{}
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

	agentBeliefs, err := models.GetAgentBeliefs(agent.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting agent beliefs.", "error": err.Error(), "id": agent.ID})
		return
	}

	if agentBeliefs == nil {
		agentBeliefs = []int64{}
	}

	agent.Beliefs = agentBeliefs

	agentGoals, err := models.GetAgentGoals(agent.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting agent goals.", "error": err.Error(), "id": agent.ID})
		return
	}

	if agentGoals == nil {
		agentGoals = []int64{}
	}

	agent.Goals = agentGoals

	agentActions, err := models.GetAgentActions(agent.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting agent actions.", "error": err.Error(), "id": agent.ID})
		return
	}

	if agentActions == nil {
		agentActions = []int64{}
	}

	agent.Actions = agentActions

	context.JSON(http.StatusOK, agent)
}

func createAgent(context *gin.Context) {
	var agent models.Agent
	context.ShouldBindJSON(&agent)

	id, err := models.CreateAgent(agent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving agent.", "error": err.Error()})
		return
	}

	agentBeliefs := []models.AgentBelief{}
	for _, beliefID := range agent.Beliefs {
		agentBeliefs = append(agentBeliefs, models.AgentBelief{AgentID: id, BeliefID: beliefID})
	}

	err = models.AddAgentBeliefs(agentBeliefs)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving agent beliefs."})
		return
	}

	agentGoals := []models.AgentGoal{}
	for _, goalID := range agent.Goals {
		agentGoals = append(agentGoals, models.AgentGoal{AgentID: id, GoalID: goalID})
	}

	err = models.AddAgentGoals(agentGoals)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving agent goals."})
		return
	}

	agentActions := []models.AgentAction{}
	for _, actionID := range agent.Actions {
		agentActions = append(agentActions, models.AgentAction{AgentID: id, ActionID: actionID})
	}

	err = models.AddAgentActions(agentActions)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving agent actions."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"id": id})
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

	agentBeliefs := []models.AgentBelief{}
	for _, beliefID := range agent.Beliefs {
		agentBeliefs = append(agentBeliefs, models.AgentBelief{AgentID: int64(id), BeliefID: beliefID})
	}

	err = models.RemoveAgentBeliefs(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error removing agent beliefs."})
		return
	}

	err = models.AddAgentBeliefs(agentBeliefs)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving agent beliefs."})
		return
	}

	agentGoals := []models.AgentGoal{}
	for _, goalID := range agent.Goals {
		agentGoals = append(agentGoals, models.AgentGoal{AgentID: int64(id), GoalID: goalID})
	}

	err = models.RemoveAgentGoals(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error removing agent goals."})
		return
	}

	err = models.AddAgentGoals(agentGoals)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving agent goals.", "error": err.Error()})
		return
	}

	agentActions := []models.AgentAction{}
	for _, actionID := range agent.Actions {
		agentActions = append(agentActions, models.AgentAction{AgentID: int64(id), ActionID: actionID})
	}

	err = models.RemoveAgentActions(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error removing agent actions."})
		return
	}

	err = models.AddAgentActions(agentActions)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving agent actions."})
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

	err = models.RemoveAgentBeliefs(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error removing agent beliefs."})
		return
	}

	err = models.RemoveAgentGoals(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error removing agent goals."})
		return
	}

	err = models.RemoveAgentActions(int64(id))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error removing agent actions."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}
