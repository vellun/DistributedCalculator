package router

import (
	"distributed-calculator/orchestrator/pkg/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine{
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-type", "Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/expressions/", handler.GetExpressionsHandler)
	router.GET("/waiting-task/", handler.GetWaitingTaskHandler)
	router.POST("/expression/", handler.PostExpressionHandler)

	return router
}