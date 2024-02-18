package request

import (
	"bytes"
	"distributed-calculator/agent/pkg/calculator"
	"distributed-calculator/orchestrator/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Здесь агент запрашивает у оркестратора задачу
func GetTask() {
	var task models.Task
	resp, err := http.Get("http://localhost:8000/waiting-task/") // Запрашиваем у оркестратора задачу
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Задача запрошена")

	if resp.StatusCode == 404 {
		fmt.Println("Нет доступных задач")
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
	fmt.Println("Задача получена и считается")
	// Считаем результат задачи
	ch := make(chan *models.Task)
	go calculator.Agent(&task, ch)

	res := <-ch
	fmt.Println("Task was count succesfully")
	fmt.Println(res.Operand1, res.Operation, res.Operand2, res.Result)
	PostTask(res)
}

// Здесь посчитанная задача отправляется обратно оркестратору
func PostTask(task *models.Task) {
	json_task, err := json.Marshal(task)

	if err != nil {
		fmt.Println(err)
		return
	}
	data := []byte(json_task)
	r := bytes.NewReader(data)
	_, err = http.Post("http://localhost:8000/task/", "application/json", r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Задача отправлена оркестратору")
}
