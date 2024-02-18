package agent

import (
	"time"
)

const (
	interval time.Duration = 10 * time.Second // Интервал времени между запросами задач
)

var ticker = time.NewTicker(interval)

func RunAgentManager(agents *CompResources) {
	for _, agent := range agents.Agents { // Запускаем всех агентов
		go agent.RunAgent()
	}
}
