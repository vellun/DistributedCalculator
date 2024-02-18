package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetPool() *pgxpool.Pool {
	params, err := GetDBParams()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get database params: %v\n", err)
		os.Exit(1)
	}
	db_url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", params.Username, params.Password,
		params.Host, params.Port, params.DBName)

	pool, err := pgxpool.Connect(context.Background(), db_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return pool
}
