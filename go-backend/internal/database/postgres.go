package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type DBParams struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

// Функция для создания всех таблиц в базе данных
func New(params DBParams) {
	db_url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", params.Username, params.Password,
		params.Host, params.Port, params.DBName)
	// Соединяемся с базой
	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var create_tables_stmt = `
		CREATE TABLE IF NOT EXISTS expressions(
			id INT generated always AS IDENTITY PRIMARY KEY,
			expression VARCHAR NOT NULL,
			status VARCHAR NOT NULL,
			started_at TIMESTAMP,
			ended_at TIMESTAMP);
		CREATE TABLE IF NOT EXISTS operations(
			id INT generated always AS IDENTITY PRIMARY KEY,
			name VARCHAR NOT NULL,
			duration TIMESTAMP);
		CREATE TABLE IF NOT EXISTS computing_servers(
			id INT generated always AS IDENTITY PRIMARY KEY,
			last_ping TIMESTAMP,
			status VARCHAR);
		CREATE TABLE IF NOT EXISTS tasks(
			id INT generated always AS IDENTITY PRIMARY KEY,
			task_expression VARCHAR,
			status VARCHAR NOT NULL,
			seq_number INTEGER NOT NULL,
			expression_id int,
			FOREIGN KEY (expression_id) REFERENCES expressions (id))`

	_, err = conn.Exec(context.Background(), create_tables_stmt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec for create tables failed: %v\n", err)
		return
	}
	fmt.Println("All tables succesfully created ;)")

	// поле status в таблице computing_servers(таблица агентов, которые сейчас работают)
	// может принимать 3 значения:
	// 1. running(сервер успешно работает и считает)
	// 2. missing(сервер пропал и какое-то время не выходил на связь, ждем от него ответа)
	// 3. died(связь с сервером окончательно потеряна и он сам удалится через время)

	// таблица tasks - таблица подзадач на которые разбивается выражение
	// столбец expression_id указывает какому выражению принадлежит задача
	// столбец статус нужен чтобы знать если вдруг задача не была решена(она должна быть передана другому серверу)
	// seq_number - порядковый номер подвыражения в выражении, чтобы считать в правильном порядке: первым умножение и тд.
}
