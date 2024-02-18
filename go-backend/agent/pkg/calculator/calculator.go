package calculator

import (
	"distributed-calculator/orchestrator/pkg/models"
	"fmt"
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

func Agent(task *models.Task, ch chan *models.Task) {
	channel := make(chan int)

	arg1, _ := strconv.Atoi(task.Operand1)
	arg2, _ := strconv.Atoi(task.Operand2)

	go counter(arg1, arg2, task.Operation, task.Duration, channel)

	select {
	case res := <-channel:
		task.Result = res
		ch <- task
	}
}
