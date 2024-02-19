package agent

import (
	"fmt"
	"time"
)

type Agent struct {
	Id          int      `json:"id"`
	Status      string   `json:"status"` //running/missing/dead
	Last_active int64      `json:"last_active"`
	Goroutines  int      // Количество горутин, которые сейчас задействует агент
	Stop        chan int // Канал для остановки работы агента(нужен, если агент завис или связь с ним потеряна)
}

func (ag *Agent) RunAgent() { // Функция запуска агента
	fmt.Printf("Agent %d is running\n", ag.Id)
	for {
		go func(id int) {
			select {
			case <-ag.Stop:
				break
			case <-ticker.C:
				GetTask(ag) // Запрашиваем задачу у оркестратора
			}
		}(ag.Id)
		time.Sleep(interval + 2)
	}
}

func NewAgent(id int) *Agent {
	return &Agent{Id: id, Status: "running", Goroutines: 0}
}
