package agent

import (
	"time"
)

var (
	interval time.Duration = 10 * time.Second // Интервал времени между запросами задач
)

var ticker = time.NewTicker(interval) // Тикер с заданным интервалом

func RunAgentManager() {  // Менеджер запускает всех агентов
	for _, agent := range Resources.Agents { // Запускаем всех агентов
		go agent.RunAgent()
	}
	go RunHealthChecker() // Запускаем проверку активности агентов
}
