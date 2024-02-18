package database

import (
	"context"
	"distributed-calculator/orchestrator/internal/database"
	"distributed-calculator/orchestrator/pkg/models"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// В функцию на вход поступает задача с посчитанным результатом и у следующей задачи, которая зависит от нее
// устанавливается значение результата в поле одного из операндов
func SetTaskResult(task *models.Task) error {
	conn := database.Connect()

	// Получаем из бд задачу, у которой одно из полей ссылается на уже посчитанную задачу
	stmt := `SELECT id, COALESCE(task_id1, 0), COALESCE(task_id2, 0) FROM tasks WHERE task_id1 = %d OR task_id2 = %d`
	rows, err := conn.Query(context.Background(), fmt.Sprintf(stmt, task.Id, task.Id))
	if err != nil {
		return errors.New(fmt.Sprintf("Query for select task from table failed: %v\n", err))
	}

	var model models.Task
	for rows.Next() {
		err := rows.Scan(&model.Id, &model.Task_id1, &model.Task_id2)
		if err != nil {
			return errors.New(fmt.Sprintf("Error occured while scan task: %v\n", err))
		}
	}

	// Если не получили никаких задач, ссылающихся на нашу, значит она последняя и выражение полностью посчитано
	if model.Id == 0 && model.Task_id1 == 0 && model.Task_id2 == 0 {
		conn := database.Connect()
		stmt = `UPDATE expressions SET result=%s WHERE id=%d`
		_, err = conn.Query(context.Background(), fmt.Sprintf(stmt, strconv.Itoa(task.Result), task.Exp_id))
		if err != nil {
			return errors.New(fmt.Sprintf("Query for update expression failed: %v\n", err))
		}

		conn = database.Connect()
		conn.Query(context.Background(), fmt.Sprintf("UPDATE expressions SET status = 'complete' WHERE id = %d", task.Exp_id))
		conn = database.Connect()
		conn.Query(context.Background(), fmt.Sprintf("UPDATE tasks SET status = 'complete' WHERE id = %d", task.Id))
		conn = database.Connect()
		conn.Query(context.Background(), fmt.Sprintf("UPDATE expressions SET ended_at = %d WHERE id = %d", time.Now().Unix(), task.Exp_id))
		return nil
	}

	if model.Task_id1 == task.Id {
		stmt = `UPDATE tasks SET operand1=%s WHERE id=%d`
	}
	if model.Task_id2 == task.Id {
		stmt = `UPDATE tasks SET operand2=%s WHERE id=%d`
	}

	conn = database.Connect()
	_, err = conn.Query(context.Background(), fmt.Sprintf(stmt, strconv.Itoa(task.Result), model.Id))
	if err != nil {
		return errors.New(fmt.Sprintf("Query for update task failed: %v\n", err))
	}

	conn = database.Connect()
	conn.Query(context.Background(), fmt.Sprintf("UPDATE tasks SET status = 'complete' WHERE id = %d", task.Id)) // Меняем статус задачи
	return nil

}
