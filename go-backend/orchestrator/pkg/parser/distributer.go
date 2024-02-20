package parser

import (
	"distributed-calculator/orchestrator/pkg/database"
	"distributed-calculator/orchestrator/pkg/models"
	"fmt"
	"strconv"
)

func DistributeTask(exp string) error {
	postfix_slice, exp_id, err := ParseExpression(exp)  // Получаем постфиксную запись
	if err != nil {
		fmt.Println(err)
		return err
	}

	var cnt int
	for {
		cnt++
		if len(postfix_slice) == 1 {
			break
		}
		l, _, r, task := GetTask(postfix_slice)

		task.Id = cnt
		task.Operation = postfix_slice[r]
		task.Exp_id = exp_id                // Устанавливаем для подвыражения id выражения
		err := database.AddTaskIntoDB(task) // Добавляем подвыражение в бд
		if err != nil {
			return err
		}

		// В списке заменяем членов выражения и оператора на task{номер подвыражения}
		postfix_slice = append(append(postfix_slice[:l], "task"+strconv.Itoa(task.Id)), postfix_slice[r+1:]...)
	}

	return nil
}

// Функция ищет в постфиксной записи подвыражения
func GetTask(postfix_slice []string) (int, int, int, *models.Task) {
	var (
		l int = 0
		m int = 1
		r int = 2
	)

	for i := 0; i < len(postfix_slice); i++ {
		_, err1 := strconv.Atoi(postfix_slice[l])
		_, err2 := strconv.Atoi(postfix_slice[m])
		p := postfix_slice[r]
		if p == "+" || p == "-" || p == "*" || p == "/" { // Если правый указатель попал на оператора
			if err1 == nil && err2 == nil { // Если два токена перед оператором - числа
				task := &models.Task{Operand1: postfix_slice[l],
					Operand2: postfix_slice[m]}
				return l, m, r, task

				// Если первый токен - подвыражение, а второй - число
			} else if string(postfix_slice[l][:len(postfix_slice[l])-1]) == "task" && err2 == nil {
				token, _ := strconv.Atoi(string(postfix_slice[l][len(postfix_slice[l])-1]))
				task := &models.Task{Task_id1: token,
					Operand2: postfix_slice[m]}
				return l, m, r, task

				// Если оба токена - подвыражения
			} else if string(postfix_slice[l][:len(postfix_slice[l])-1]) == "task" &&
				string(postfix_slice[m][:len(postfix_slice[m])-1]) == "task" {
				token1, _ := strconv.Atoi(string(postfix_slice[l][len(postfix_slice[l])-1]))
				token2, _ := strconv.Atoi(string(postfix_slice[m][len(postfix_slice[m])-1]))
				task := &models.Task{Task_id1: token1, Task_id2: token2}
				return l, m, r, task

				// Если первый токен - число, а второй - подвыражение
			} else if err1 == nil && string(postfix_slice[m][:len(postfix_slice[m])-1]) == "task" {
				token, _ := strconv.Atoi(string(postfix_slice[m][len(postfix_slice[m])-1]))
				fmt.Println("token", token)
				task := &models.Task{Operand1: postfix_slice[l],
					Task_id2: token}
				return l, m, r, task
			}

		}
		l++
		m++
		r++
	}
	return l, m, r, nil
}
