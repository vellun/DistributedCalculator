package main

import (
	"distributed-calculator/cmd/orchesrtator"
	// "distributed-calculator/internal/config"
)

func main() {
	// config.Сreate_all_tables_in_db()
	orchesrtator.Orchestrator("3+4-6*5")
}
