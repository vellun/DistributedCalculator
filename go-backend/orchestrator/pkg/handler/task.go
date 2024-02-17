package handler

import (
	"distributed-calculator/orchestrator/pkg/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Отдаем подвыражение агенту, если у нас есть что посчитать
func GetWaitingTaskHandler(c *gin.Context) {
	task, err := database.GetWaitingTask()
	fmt.Println(task)
	if err != nil || task == nil {
		fmt.Println("здесть ошибка")
		fmt.Println(task)
		fmt.Println(err)
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, task)

}
