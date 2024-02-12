package parser

import (
	"distributed-calculator/pkg/database"
	"fmt"
)

var postfix_slice []string

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
	exp_id, err := database.AddExpressionIntoDB(&database.Expression{Expression: exp, Status: "process"})
	if err != nil {
		return nil, 0, err
	}

	get_postfix_string(tree) // Получаем выражение в постфиксной форме
	// Например выражение 1 + 2 * 3 примет вид 123*+

	return postfix_slice, exp_id, nil
}

// Функция рекурсивно обходит дерево
func get_postfix_string(tree *Node) {
	if tree == nil {
		return
	}
	get_postfix_string(tree.Left)
	get_postfix_string(tree.Right)
	postfix_slice = append(postfix_slice, tree.Value)
}
