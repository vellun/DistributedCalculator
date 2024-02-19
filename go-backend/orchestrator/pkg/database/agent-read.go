package database

import (
	"context"
	"distributed-calculator/orchestrator/internal/database"
	"distributed-calculator/orchestrator/pkg/models"
	"errors"
	"fmt"
)

func GetAllAgents() ([]models.Agent, error) {
	conn := database.Connect()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, last_active, status, ind FROM computing_resources ORDER BY id;")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Query for select agents from table failed: %v\n", err))
	}

	agents := []models.Agent{}
	for rows.Next() {
		var agent models.Agent
		err := rows.Scan(&agent.Id, &agent.Last_active, &agent.Status, &agent.Goroutines)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occured while scan agents: %v\n", err))
		}
		agent.Goroutines = 0
		agents = append(agents, agent)
	}
	return agents, nil
}
