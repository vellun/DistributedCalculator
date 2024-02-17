package parser

import (
	"distributed-calculator/orchestrator/pkg/database"
	"distributed-calculator/orchestrator/pkg/models"
	"fmt"
)

type postfix_slice struct{
	slice []string
}

func ParseExpression(exp string) ([]string, int, error) {
	token_list, exp, err := ValidateExpression(exp) // Функция находится в файле validator.go
	if err != nil {
		fmt.Println("Выражение некорректно")
		return nil, 0, err
	}

	tokens := &Tokens{List: token_list}
	tree, err := tokens.getExpTree() // Получаем дерево

	if err != nil {
		fmt.Println("Ошибка при парсинге")
		return nil, 0, err
	}

	// Если выражение прошло все проверки, добавляем его в бд
	exp_id, err := database.AddExpressionIntoDB(&models.Expression{Expression: exp, Status: "process"})
	if err != nil {
		return nil, 0, err
	}

	ps := &postfix_slice{slice: []string{}}
	ps.get_postfix_string(tree) // Получаем выражение в постфиксной форме
	// Например выражение 1 + 2 * 3 примет вид 123*+

	return ps.slice, exp_id, nil
}

// Функция рекурсивно обходит дерево
func (ps *postfix_slice) get_postfix_string(tree *Node) {
	if tree == nil {
		return
	}
	ps.get_postfix_string(tree.Left)
	ps.get_postfix_string(tree.Right)
	ps.slice = append(ps.slice, tree.Value)
}
