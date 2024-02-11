package parser

import (
	"errors"
	"strconv"
	"strings"
)

// Функция проверяет выражение на корректность
func ValidateExpression(exp string) ([]string, error) {
	var list []string

	exp = strings.Join(strings.Fields(exp), "")

	// Если в начале выражения не число
	if _, err := strconv.Atoi(string(exp[0])); err != nil {
		return nil, errors.New("Выражение должно начинаться с числа")
	}

	var num string // Строка, в которой собираются цифры числа

	for i, t := range exp {
		token := string(t)

		switch token {
		case "*", "/", "-", "+", "(", ")":
			// Если дошли до знака, добавляем в список получившееся число
			if num != "" {
				list = append(list, num)
				num = ""
			}
			list = append(list, token)
		default:
			if _, err := strconv.Atoi(token); err != nil {
				return nil, errors.New("Найден недопустимый символ")
			}
			num += token
		}

		// Если дошли до конца выражения, тоже добавляем в список получившееся число
		if i == len(exp)-1 {
			if num != "" {
				list = append(list, num)
				num = ""
			}
		}
	}

	list = append(list, "end") //Обозначим конец выражения(нужно при создании дерева)

	return list, nil

}
