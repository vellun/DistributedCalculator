package agent

import (
	"time"
)

const (
	interval time.Duration = 10 * time.Second // Интервал времени между запросами задач
)

var ticker = time.NewTicker(interval)

func RunAgentManager() {
	for _, agent := range Resources.Agents { // Запускаем всех агентов
		go agent.RunAgent()
	}
	go RunHealthChecker() // Запускаем проверку активности агентов
}
