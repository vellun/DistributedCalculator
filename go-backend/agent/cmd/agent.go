package agent

import (
	"distributed-calculator/agent/pkg/request"
	// "distributed-calculator/agent/pkg/router"
	"fmt"

	// "distributed-calculator/agent/pkg/router"
	// "fmt"
	"time"
)

const (
	interval time.Duration = 10 * time.Second // Интервал времени между запросами задач
)

var ticker = time.NewTicker(interval)

func worker() {
	<-ticker.C
	request.GetTask() // Запрашиваем задачу у оркестратора
}

func RunAgentManager() {
	fmt.Println("Agent's goroutine is running")
	for {
		go worker()
		time.Sleep(interval + 2)
	}
}
