package routes

import "github.com/gin-gonic/gin"

func InitialiseRoutes(router *gin.Engine) {
	AgentRoutes(router)
	ActionRoutes(router)
	GoalRoutes(router)
	BeliefRoutes(router)
	LocationRoutes(router)
}
