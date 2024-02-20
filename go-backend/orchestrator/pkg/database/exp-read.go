package database

import (
	"context"
	"distributed-calculator/orchestrator/pkg/models"
	"distributed-calculator/orchestrator/postgres"
	"errors"
	"fmt"
)

// Получение всех выражений для страницы истории
func GetAllExpressions() ([]models.Expression, error) {
	conn := postgres.Connect()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, expression, status, started_at, ended_at, COALESCE(result, '') FROM expressions ORDER BY id DESC;")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Query for select expressions from table failed: %v\n", err))
	}

	expressions := []models.Expression{}
	for rows.Next() {
		var exp models.Expression
		err := rows.Scan(&exp.Id, &exp.Expression, &exp.Status, &exp.Started_at, &exp.Ended_at, &exp.Result)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occured while scan expressions: %v\n", err))
		}
		expressions = append(expressions, exp)
	}
	return expressions, nil
}
