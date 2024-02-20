package main

import (
	"distributed-calculator/agent"
	"distributed-calculator/orchestrator/pkg/router"
	"distributed-calculator/orchestrator/postgres"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// Инициализируем конфиг
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err)
	}

	postgres.InitRepository()  // Создаем все таблицы в бд если их нет
	router := router.NewRouter()  // Роутер запросов

	agent.Resources.Init() // Создаются агенты

	agent.RunAgentManager() // Запускаем менеджера агентов
	// Она запустит агентов(горутины), которые будут запрашивать у оркестратора задачи для решения через установленные промежутки времени,
	// а затем отправлять на решение одному из своих воркеров(горутине)

	if err := router.Run(viper.GetString("port")); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
