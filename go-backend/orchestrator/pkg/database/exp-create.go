package database

import (
	"context"
	"distributed-calculator/orchestrator/pkg/models"
	"distributed-calculator/orchestrator/postgres"
	"errors"
	"fmt"
	"time"
)

func AddExpressionIntoDB(exp *models.Expression) (int, error) { // Возвращает id добавленного выражения
	conn := postgres.Connect()
	defer conn.Close(context.Background())

	exp.Started_at = time.Now().Unix()

	var insertStmt string = fmt.Sprintf("INSERT INTO expressions(expression, status, started_at, ended_at) VALUES ('%s', '%s', %d, %d)",
		exp.Expression, exp.Status, exp.Started_at, 0)
	_, err := conn.Exec(context.Background(), insertStmt)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Exec for insert expression into table failed: %v\n", err))
	}

	conn = postgres.Connect()
	defer conn.Close(context.Background())
	var exp_id int
	// Получаем id только что добавленного выражения
	id, _ := conn.Query(context.Background(), "SELECT MAX(id) FROM expressions;")
	for id.Next() {
		id.Scan(&exp_id)
	}

	fmt.Println("Expression was succesfully insert")
	return exp_id, nil
}
