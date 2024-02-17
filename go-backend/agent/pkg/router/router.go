package router

import (
	// "distributed-calculator/orchestrator/pkg/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine{
	router := gin.Default()

	// router.GET("/expressions/", handler.GetExpressionsHandler)
	// router.GET("/waiting-task/", handler.GetWaitingTaskHandler)
	// router.POST("/expression/", handler.PostExpressionHandler)

	return router
}