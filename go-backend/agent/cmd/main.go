package main

import (
	"distributed-calculator/agent/pkg/request"
	// "distributed-calculator/agent/pkg/router"
	// "fmt"
	"time"
)

var ticker = time.NewTicker(10 * time.Second)

func worker() {
	for {
		select {
		case <-ticker.C:
			request.GetTask()
		}

		// ch <- fmt.Sprintf("worker %d: завершил задачу", id)
	}
}

func main() {  // Агент запрашивает у оркестратора задачу раз в 10 секунд
	go worker()

	for {
	}
}
