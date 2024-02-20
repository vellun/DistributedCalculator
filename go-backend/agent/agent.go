package agent

import (
	"fmt"
	"time"
)

type Agent struct {
	Id          int    `json:"id"`
	Status      string `json:"status"` //running/missing/dead
	Last_active int64  `json:"last_active"`
	Goroutines  int    // Количество горутин, которые сейчас задействует агент
}

func (ag *Agent) RunAgent() { // Функция запуска агента
	fmt.Printf("Agent %d is running\n", ag.Id)
	for {
		if ag.Status == "dead" {
			fmt.Printf("Агент %d: Я умер, не буду запрашивать задачу\n", ag.Id)
			return
		}
		go func(id int) {
			<-ticker.C
			GetTask(ag) // Запрашиваем задачу у оркестратора

		}(ag.Id)
		time.Sleep(interval + 2)
	}

}

func NewAgent(id int) *Agent {
	return &Agent{Id: id, Status: "running"}
}
