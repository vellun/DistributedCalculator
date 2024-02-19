package database

import (
	"context"
	"distributed-calculator/orchestrator/internal/database"
	"fmt"
)

// Меняет статус агента в бд
func UpdateStatus(id int, status string) error {
	conn := database.Connect()
	defer conn.Close(context.Background())

	stmt := `UPDATE computing_resources SET status=%s WHERE id=%d`
	_, err := conn.Query(context.Background(), fmt.Sprintf(stmt, status, id))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Обновляет время последней активности агента
func UpdateLastActive(id int, timestamp int64) error {
	conn := database.Connect()
	defer conn.Close(context.Background())

	stmt := `UPDATE computing_resources SET last_active=%d WHERE id=%d`
	_, err := conn.Query(context.Background(), fmt.Sprintf(stmt, timestamp, id))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
