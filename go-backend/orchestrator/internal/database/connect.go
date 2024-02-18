package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DBParams struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

// Возвращает новое соединение с бд
func Connect(params DBParams) (conn *pgxpool.Conn) {
	db_url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", params.Username, params.Password,
		params.Host, params.Port, params.DBName)
	// Соединяемся с базой
	pool, err := pgxpool.Connect(context.Background(), db_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	conn, err = pool.Acquire(context.Background())
	if err != nil {
		fmt.Sprintf("Unable to acquire a database connection: %v\n", err)
	}
	return conn
}
