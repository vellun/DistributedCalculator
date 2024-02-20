package database

import (
	"context"
	"distributed-calculator/orchestrator/pkg/models"
	"distributed-calculator/orchestrator/postgres"
	"errors"
	"fmt"
)

// Функция берет из бд первое найденное подвыражение со статусом process
func GetWaitingTask() (*models.Task, error) {
	conn := postgres.Connect()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id FROM expressions WHERE status = 'process' ORDER BY id;")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Query for select expression from table failed: %v\n", err))
	}

	var exp_id int
	for rows.Next() {
		err := rows.Scan(&exp_id) // Получаем id самого старого выражения в бд со статусом process
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occured while scan expression id: %v\n", err))
		}
		break // Получаем одно значение и прерываем цикл
	}

	stmt := fmt.Sprintf(`
		SELECT tasks.id, tasks.operand1, tasks.operand2, tasks.status, 
		tasks.expression_id, operations.name, operations.duration FROM tasks 
		JOIN operations ON tasks.operation_id=operations.id
		WHERE tasks.expression_id = %d 
		AND tasks.status = 'process' 
		AND tasks.operand1 IS NOT NULL
		AND tasks.operand2 IS NOT NULL ORDER BY tasks.id;
	`, exp_id)
	is_null_stmt := fmt.Sprintf(`SELECT COUNT(*) FROM tasks 
			JOIN operations ON tasks.operation_id=operations.id
			WHERE tasks.expression_id = %d 
			AND tasks.status = 'process' 
			AND tasks.operand1 IS NOT NULL
			AND tasks.operand2 IS NOT NULL;`, exp_id)

	var n int

	conn = postgres.Connect()
	defer conn.Close(context.Background())
	num, err := conn.Query(context.Background(), is_null_stmt)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Query for select task from table failed: %v\n", err))
	}

	for num.Next() {
		num.Scan(&n)
		// Если нет подвыражений которые можно подсчитать, возвращаем nil
		if n == 0 {
			return nil, nil
		}
	}

	conn = postgres.Connect()
	defer conn.Close(context.Background())
	rows, err = conn.Query(context.Background(), stmt)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Query for select task from table failed: %v\n", err))
	}

	var task models.Task
	for rows.Next() {
		err := rows.Scan(&task.Id, &task.Operand1, &task.Operand2, &task.Status, &task.Exp_id, &task.Operation, &task.Duration)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error occured while scan task: %v\n", err))
		}
	}
	return &task, nil

}
