package database

import (
	"context"
	"distributed-calculator/orchestrator/pkg/models"
	"distributed-calculator/orchestrator/postgres"
	"errors"
	"fmt"
)

func UpdateOperationDuration(operation *models.Operation) error {
	conn := postgres.Connect()
	defer conn.Close(context.Background())
	stmt := `UPDATE operations SET duration=%d WHERE id=%d`
	_, err := conn.Query(context.Background(), fmt.Sprintf(stmt, operation.Duration, operation.Id))
	if err != nil {
		fmt.Println("aa")
		return errors.New(fmt.Sprintf("Query for update operation failed: %v\n", err))
	}
	return nil
}
