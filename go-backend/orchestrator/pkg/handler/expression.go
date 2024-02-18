package handler

import (
	"distributed-calculator/orchestrator/pkg/database"
	"distributed-calculator/orchestrator/pkg/parser"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type expression struct {
	Exp string
}

// Получить все выражения
func GetExpressionsHandler(c *gin.Context) {
	exps, err := database.GetAllExpressions()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, exps)
}

// Получить введенное пользователем выражение
func PostExpressionHandler(c *gin.Context) {
	var recieved_exp expression

	if err := c.BindJSON(&recieved_exp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not valid expression"})
		return
	}

	if err := parser.DistributeTask(recieved_exp.Exp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not valid expression"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": ""})
}
