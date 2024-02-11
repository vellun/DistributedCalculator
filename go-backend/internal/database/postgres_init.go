package database

import (
	"context"
	"fmt"
	"os"
)

// Функция для создания всех таблиц в базе данных
func New(params DBParams) {
	conn := Connect(params)
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
			duration INT);
		CREATE TABLE IF NOT EXISTS computing_servers(
			id INT generated always AS IDENTITY PRIMARY KEY,
			last_ping TIMESTAMP,
			status VARCHAR);
		CREATE TABLE IF NOT EXISTS tasks(
			id INT generated always AS IDENTITY PRIMARY KEY,
			first_operand INT,
			second_operand INT,
			operation_id VARCHAR,
			status VARCHAR,
			seq_number INTEGER NOT NULL,
			expression_id int,
			FOREIGN KEY (expression_id) REFERENCES expressions (id)
			FOREIGN KEY (operation_id) REFERENCES operations (id))`

	_, err := conn.Exec(context.Background(), create_tables_stmt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec for create tables failed: %v\n", err)
		return
	}
	fmt.Println("All tables succesfully created ;)")

	// Некоторые пояснения к структуре базы данных:

	// поле status в таблице computing_servers(таблица агентов, которые сейчас работают)
	// может принимать 3 значения:
	// 1. running(сервер успешно работает и считает)
	// 2. missing(сервер пропал и какое-то время не выходил на связь, ждем от него ответа)
	// 3. died(связь с сервером окончательно потеряна и он сам удалится через время)

	// таблица tasks - таблица подзадач на которые разбивается выражение
	// столбец expression_id указывает какому выражению принадлежит задача
	// столбец статус нужен чтобы знать если вдруг задача не была решена(она должна быть передана другому серверу)
	// seq_number - порядковый номер подвыражения в выражении, чтобы считать в правильном порядке: первым умножение и тд.
	// first_operand, second_operand - члены выражения

	var n string
	// Проверяем пустая ли таблица с операциями
	num, _ := conn.Query(context.Background(), "SELECT COUNT(*) FROM operations;")
	for num.Next() {
		num.Scan(&n)
		// Если не пустая, ничего не делаем и выходим
		if n != "0" {
			return
		}
	}
	// Добавляем доступные операции
	var insertStmt string = `INSERT INTO operations(name, duration) VALUES 
							('plus', 200),
							('minus', 200), 
							('times', 200), 
							('divide', 200)`
	_, err = conn.Exec(context.Background(), insertStmt)
	if err != nil {
		fmt.Printf("Exec for insert operations into table failed: %v\n", err)
	}
	fmt.Println("Succesfully inserted default operations)")
}
