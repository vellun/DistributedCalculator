package main

import (
	"distributed-calculator/cmd/orchesrtator"
	"distributed-calculator/internal/config"
)

func main() {
	config.Сreate_all_tables_in_db()
	// 3+(4-6) /2 +(4-2)
	// (3 + 5) * 3 / 2 + 423 - 3
	// 3+(4-6) /2 +(4-2)
	orchesrtator.Orchestrator("(3 + 5) * 3")
}
