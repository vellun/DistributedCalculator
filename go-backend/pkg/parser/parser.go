package parser

import (
	"fmt"
)

var prefix_string string

func ParseExpression(exp string) (string, error) {
	token_list, err := ValidateExpression(exp) // Функция находится в файле validator.go
	if err != nil {
		fmt.Println("Выражение некорректно")
		return "", err
	}

	tokens := &Tokens{List: token_list}
	tree, err := tokens.getExpTree() // Получаем дерево

	if err != nil {
		fmt.Println("Ошибка при парсинге")
		return "", err
	}

	get_prefix_string(tree) // Получаем выражение в префиксной форме
	// Например выражение 1 + 2 * 3 примет вид + 1 * 2 3 (сначала корень, потом дочерние узлы)

	return prefix_string, nil
}

// Функция рекурсивно обходит дерево
func get_prefix_string(tree *Node) {
	if tree == nil {
		return
	}
	prefix_string += tree.Value
	get_prefix_string(tree.Left)
	get_prefix_string(tree.Right)
}
