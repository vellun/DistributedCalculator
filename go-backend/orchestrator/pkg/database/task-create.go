package database

import (
	"context"
	"distributed-calculator/orchestrator/internal/config"
	"distributed-calculator/orchestrator/internal/database"
	"distributed-calculator/orchestrator/pkg/models"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func AddTaskIntoDB(task *models.Task) error {
	fmt.Println(task.Operand1, task.Operand2, task.Operation)
	DBParams, err := config.GetDBParams()
	if err != nil {
		return errors.New("Cannont connect to database. Params are wrong")
	}
	conn := database.Connect(DBParams)

	// Если в качестве одного из членов подвыражения должна быть ссылка на другое подвыражение
	// Нужно получить id этого подвыражения из бд(их может быть сразу два)
	if task.Task_id1 != 0 || task.Task_id2 != 0 {
		fmt.Println(task.Task_id1, task.Task_id2)
		task.Task_id1, task.Task_id2 = GetTasksId(task, conn)
		fmt.Println("aa", task.Task_id1, task.Task_id2)
	}

	var insertStmt string

	if task.Task_id1 != 0 && task.Task_id2 != 0 {
		insertStmt = fmt.Sprintf("INSERT INTO tasks(expression_id, status, task_id1, task_id2, seq_number, operation_id) VALUES (%d, '%s', %d, %d, %d, %d)",
			task.Exp_id, "process", task.Task_id1, task.Task_id2, task.Id, GetOperationId(task, conn))
	} else if task.Task_id1 == 0 && task.Task_id2 == 0 {
		insertStmt = fmt.Sprintf("INSERT INTO tasks(expression_id, operand1, operand2, status, seq_number, operation_id) VALUES (%d, %s, %s, '%s', %d, %d)",
			task.Exp_id, task.Operand1, task.Operand2, "process", task.Id, GetOperationId(task, conn))
	} else if task.Task_id1 == 0 {
		insertStmt = fmt.Sprintf("INSERT INTO tasks(expression_id, operand1, status, task_id2, seq_number, operation_id) VALUES (%d, %s, '%s', %d, %d, %d)",
			task.Exp_id, task.Operand1, "process", task.Task_id2, task.Id, GetOperationId(task, conn))
	} else {
		insertStmt = fmt.Sprintf("INSERT INTO tasks(expression_id, operand2, status, task_id1, seq_number, operation_id) VALUES (%d, %s, '%s', %d, %d, %d)",
			task.Exp_id, task.Operand2, "process", task.Task_id1, task.Id, GetOperationId(task, conn))
	}
	_, err = conn.Exec(context.Background(), insertStmt)
	if err != nil {
		fmt.Printf("Exec for insert task into table failed: %v\n", err)
		return err
	}
	fmt.Println("Task was succesfully insert")
	return nil
}

func GetTasksId(task *models.Task, conn *pgxpool.Conn) (int, int) {
	var (
		task_stmt1 int
		task_stmt2 int
	)
	if task.Task_id1 != 0 {
		task_stmt1 = task.Task_id1
	}
	if task.Task_id2 != 0 {
		task_stmt2 = task.Task_id2
	}
	var selectStmt1 string = fmt.Sprintf("SELECT id FROM tasks WHERE seq_number=%d AND expression_id=%d;", task_stmt1, task.Exp_id)
	var selectStmt2 string = fmt.Sprintf("SELECT id FROM tasks WHERE seq_number=%d AND expression_id=%d;", task_stmt2, task.Exp_id)

	var (
		task_id1 int
		task_id2 int
	)

	// Получаем id нужного подвыражения
	id1, _ := conn.Query(context.Background(), selectStmt1)
	for id1.Next() {
		id1.Scan(&task_id1)
	}

	id2, _ := conn.Query(context.Background(), selectStmt2)
	for id2.Next() {
		id2.Scan(&task_id2)
	}
	return task_id1, task_id2
}

func GetOperationId(task *models.Task, conn *pgxpool.Conn) int {
	var id int
	stmt := fmt.Sprintf("SELECT id FROM operations WHERE name='%s';", task.Operation)
	op_id, _ := conn.Query(context.Background(), stmt)
	for op_id.Next() {
		op_id.Scan(&id)
	}
	return id
}
