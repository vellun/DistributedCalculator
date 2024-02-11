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

func Connect(params DBParams) (conn *pgx.Conn) {
	db_url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", params.Username, params.Password,
		params.Host, params.Port, params.DBName)
	// Соединяемся с базой
	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
