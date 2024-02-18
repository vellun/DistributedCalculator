package handler

import (
	"distributed-calculator/orchestrator/pkg/database"
	"distributed-calculator/orchestrator/pkg/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Отдаем подвыражение агенту, если у нас есть что посчитать
func GetWaitingTaskHandler(c *gin.Context) {
	task, err := database.GetWaitingTask()
	if err != nil || task == nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, task)
}

// Здесь получаем таску с заполненным полем Result после того как агент посчитал
func PostResultTaskHandler(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	err := database.SetTaskResult(&task)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
