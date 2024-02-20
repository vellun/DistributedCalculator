package agent

import (
	"bytes"
	"distributed-calculator/orchestrator/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Горутина, считающая субвыражения(вычислитель, воркер)
func counter(arg1, arg2 int, op string, duration int, ch chan int) {
	fmt.Println("длительность", duration, "секунд")
	time.Sleep(time.Second * time.Duration(duration)) // Имитация долгих вычислений
	switch op {
	case "+":
		ch <- arg1 + arg2
	case "-":
		ch <- arg1 - arg2
	case "*":
		ch <- arg1 * arg2
	case "/":
		ch <- arg1 / arg2
	}

}

// Сюда попадает задача от агента
func Calculator(task *models.Task, agent *Agent) {
	channel := make(chan int)

	arg1, _ := strconv.Atoi(task.Operand1)
	arg2, _ := strconv.Atoi(task.Operand2)

	go counter(arg1, arg2, task.Operation, task.Duration, channel) // Отправляем задачу считаться
	time.Sleep(time.Second * time.Duration(task.Duration))

	res := <-channel // Ждем пока задача посчитается
	task.Result = res

	fmt.Println("Task was count succesfully")

	// Отправляем посчитанную задачу оркестратору
	PostTask(task, agent)

}

// Здесь посчитанная задача отправляется обратно оркестратору
func PostTask(task *models.Task, agent *Agent) {
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
	agent.Goroutines-- // Горутина отработала, отнимаем 1 от общего количества работающих горутин у агента
	fmt.Println("Задача отправлена оркестратору агентом", agent.Id)
}
