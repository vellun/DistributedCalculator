package agent

import (
	"time"

	"github.com/spf13/viper"
)

var (
	interval time.Duration = time.Duration(viper.GetInt64("agent.get_task_time")) * time.Second // Интервал времени между запросами задач
)

var ticker = time.NewTicker(interval) // Тикер с заданным интервалом

func RunAgentManager() {  // Менеджер запускает всех агентов
	for _, agent := range Resources.Agents { // Запускаем всех агентов
		go agent.RunAgent()
	}
	go RunHealthChecker() // Запускаем проверку активности агентов
}
