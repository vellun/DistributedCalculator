package handler

import (
	"distributed-calculator/orchestrator/pkg/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Получить все операции
func GetAgentsHandler(c *gin.Context) {
	agents, err := database.GetAllAgents()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, agents)
}