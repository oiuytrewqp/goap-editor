package routes

import (
	"net/http"
	"oiuytrewqp/goap-server/server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BeliefRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/beliefs", getBeliefs)
		api.GET("/beliefs/:id", getBelief)
		api.POST("/beliefs", createBelief)
		api.PUT("/beliefs/:id", updateBelief)
		api.DELETE("/beliefs/:id", deleteBelief)
	}
}

func getBeliefs(context *gin.Context) {
	allBelliefs, err := models.GetBeliefs()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting beliefs."})
		return
	}

	context.JSON(http.StatusOK, allBelliefs)
}

func getBelief(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	belief, err := models.GetBelief(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting belief."})
		return
	}

	context.JSON(http.StatusOK, belief)
}

func createBelief(context *gin.Context) {
	var belief models.Belief
	context.ShouldBindJSON(&belief)

	id, err := models.CreateBelief(belief)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving belief."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"id": id})
}

func updateBelief(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	var belief models.Belief
	context.ShouldBindJSON(&belief)

	err = models.UpdateBelief(id, belief)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating belief."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

func deleteBelief(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."})
		return
	}

	err = models.DeleteBelief(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting belief."})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}
