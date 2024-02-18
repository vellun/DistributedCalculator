package main

import (
	// "context"
	"distributed-calculator/agent"
	"distributed-calculator/orchestrator/internal/database"
	"distributed-calculator/orchestrator/internal/router"
	"fmt"
)

const (
	port = ":8000"
)

var Repo database.Repository = *database.NewRepository(database.Pool)

func main() {
	// Repo.Init(context.Background())
	router := router.NewRouter()

	agents := agent.NewResources()
	agents.Init()

	agent.RunAgentManager(agents) // Запускаем горутину менеджера агентов
	// Она запустит агентов(горутины), которые будут запрашивать у оркестратора задачи для решения через установленные промежутки времени,
	// а затем отправлять на решение одному из своих воркеров(горутина)

	if err := router.Run(port); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
}
