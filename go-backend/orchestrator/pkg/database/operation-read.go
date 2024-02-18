package database

import (
	"context"
	"distributed-calculator/orchestrator/internal/database"
	"distributed-calculator/orchestrator/pkg/models"
	"errors"
	"fmt"
)

func GetAllOperations() ([]models.Operation, error) {
	conn := database.Connect()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, name, duration FROM operations;")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Query for select operations from table failed: %v\n", err))
	}

	operations := []models.Operation{}
	for rows.Next() {
		var op models.Operation
		err := rows.Scan(&op.Id, &op.Name, &op.Duration)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occured while scan operations: %v\n", err))
		}
		operations = append(operations, op)
	}
	return operations, nil
}
