package main

import (
	// "distributed-calculator/agent"
	// "distributed-calculator/orchestrator/internal/database"
	"distributed-calculator/agent"
	"distributed-calculator/orchestrator/internal/router"
	"fmt"
)

const (
	port = ":8000"
)

func main() {
	// database.InitRepository()
	router := router.NewRouter()

	agent.Resources.Init() // Создаются агенты

	// agent.RunAgentManager(agents) // Запускаем горутину менеджера агентов
	// Она запустит агентов(горутины), которые будут запрашивать у оркестратора задачи для решения через установленные промежутки времени,
	// а затем отправлять на решение одному из своих воркеров(горутина)

	if err := router.Run(port); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
}
