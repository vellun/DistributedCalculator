package agent

import (
	"distributed-calculator/orchestrator/pkg/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type agentId struct {
	Id int
}

func DisconnectAgentHandler(c *gin.Context) {
	var recieved_id agentId

	if err := c.BindJSON(&recieved_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if err := database.UpdateStatus(recieved_id.Id, "dead"); err != nil { // Меняем статус в бд
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	for _, agent := range Resources.Agents {
		if agent.Id == recieved_id.Id { // Меняем статус в списке ресурсов
			agent.Status = "dead"
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})

}
