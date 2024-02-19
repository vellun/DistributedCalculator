package agent

import (
	"distributed-calculator/orchestrator/pkg/database"
	"distributed-calculator/orchestrator/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Здесь агент запрашивает у оркестратора задачу и отправляет на вычисление
func GetTask(agent *Agent) {
	var task models.Task
	resp, err := http.Get("http://localhost:8000/waiting-task/") // Запрашиваем у оркестратора задачу
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Задача запрошена агентом", agent.Id)

	agent.Last_active = time.Now().Unix()                        // Агент запросил задачу, а значит надо отметить последнюю активность
	err = database.UpdateLastActive(agent.Id, agent.Last_active) // Также отмечаем это в бд
	if agent.Status == "missing"{
	agent.Status = "running"
	database.UpdateStatus(agent.Id, "running") // И меняем статус на случай если он был missing
	}

	if err != nil {
		fmt.Println("Не удалось обновить время активности агента\n", agent.Id)
	} else {
		fmt.Printf("Агент %d: Время последней активности обновлено\n", agent.Id)
	}

	if resp.StatusCode == 404 {
		fmt.Printf("Агент %d: Нет доступных задач\n", agent.Id)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if body == nil { // Если не получили задачу
		return
	}

	err = json.Unmarshal(body, &task)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Задача получена агентом", agent.Id)

	fmt.Printf("У агента %d %d горутин\n", agent.Id, agent.Goroutines)

	if agent.Goroutines < 5 { // Если действующих горутин у агента < 5
		// Отправляем задачу считаться
		agent.Goroutines++
		go Calculator(&task, agent)
	} else {
		fmt.Printf("Агент %d не смог принять задачу: все горутины заняты\n", agent.Id)
	}
}
