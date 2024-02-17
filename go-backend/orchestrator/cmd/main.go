package main

import (
	"distributed-calculator/orchestrator/internal/config"
	"distributed-calculator/orchestrator/internal/router"
	"fmt"
)

const (
	port = ":8080"
)

func init_db() {
	config.Сreate_all_tables_in_db()
}

func main() {
	init_db()
	router := router.NewRouter()

	if err := router.Run(port); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
}
