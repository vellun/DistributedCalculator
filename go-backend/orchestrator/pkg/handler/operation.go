package handler

import (
	"distributed-calculator/orchestrator/pkg/database"
	"distributed-calculator/orchestrator/pkg/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Получить все операции
func GetOperationsHandler(c *gin.Context) {
	operations, err := database.GetAllOperations()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, operations)
}

func PostOperationDurationHandler(c *gin.Context) {
	var recieved_op models.Operation

	if err := c.BindJSON(&recieved_op); err != nil {
		fmt.Println(err) // Если что-то пошло не так, не возвращаем плохой статус, а просто игнорируем изменения
		c.JSON(http.StatusOK, gin.H{})
	}

	if recieved_op.Duration < 0 {
		c.JSON(http.StatusOK, gin.H{})
	}

	if err := database.UpdateOperationDuration(&recieved_op); err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
