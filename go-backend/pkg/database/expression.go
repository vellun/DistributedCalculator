package database

import (
	"context"
	"distributed-calculator/internal/config"
	"distributed-calculator/internal/database"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Expression struct {
	Expression string
	Status     string
	started_at int64
}

func AddExpressionIntoDB(exp *Expression) (int, error) { // Возвращает id добавленного выражения
	DBParams, err := config.GetDBParams()
	if err != nil {
		return 0, errors.New("Cannont connect to database. Params are wrong")
	}
	conn := database.Connect(DBParams)
	defer conn.Close(context.Background())

	exp.started_at = time.Now().Unix()

	var insertStmt string = fmt.Sprintf("INSERT INTO expressions(expression, status, started_at) VALUES ('%s', '%s', %d)",
		exp.Expression, exp.Status, exp.started_at)
	_, err = conn.Exec(context.Background(), insertStmt)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Exec for insert expression into table failed: %v\n", err))
	}

	var exp_id string
	// Получаем id только что добавленного выражения
	id, _ := conn.Query(context.Background(), "SELECT MAX(id) FROM expressions;")
	for id.Next() {
		id.Scan(&exp_id)
	}
	expression_id, _ := strconv.Atoi(exp_id)
	fmt.Println("Expression was succesfully insert")
	return expression_id, nil
}
