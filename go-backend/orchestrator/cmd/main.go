package main

import (
	agent "distributed-calculator/agent/cmd"
	"distributed-calculator/orchestrator/internal/config"
	"distributed-calculator/orchestrator/internal/router"
	"fmt"
)

const (
	port = ":8000"
)

func init_db() {
	config.Сreate_all_tables_in_db()
}

func main() {
	// init_db()
	router := router.NewRouter()

	go agent.RunAgentManager() // Запускаем горутину менеджера агента
	// Она будет запрашивать у оркестратора задачи для решения через установленные промежутки времени
	// И отправлять на решение агенту

	if err := router.Run(port); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
}
