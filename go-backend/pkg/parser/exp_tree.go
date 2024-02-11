package parser

import (
	"errors"
	"fmt"
	"strconv"
)

// Создание дерева выражений

// Узел дерева выражений
type Node struct {
	Value string
	Left  *Node
	Right *Node
}

type Tokens struct {
	List []string
}

func NewNode(value string, left, right *Node) *Node {
	return &Node{Value: value, Left: left, Right: right}
}

// Функция сравнивает первый токен в списке с ожидаемыми значениями и удаляет из списка
func (tokens *Tokens) popExpectedToken(expected1, expected2 string) (string, bool) {
	top := tokens.List[0]
	if top == expected1 || top == expected2 {
		tokens.List = tokens.List[1:] // Удаляем ожидаемый токен из списка
		return top, true
	}
	return "", false
}

// Возвращает узел, содержащий число или скобочное выражение
func (tokens *Tokens) getNumberNode() (*Node, error) {
	// Если встретилось выражение в скобках
	_, is_expected := tokens.popExpectedToken("(", "(")
	if is_expected {
		node, _ := tokens.getExpTree()                  // Получаем узел с выражением в скобках
		_, is_expected := tokens.popExpectedToken(")", ")") // Пробуем получить пару для первой скобки
		if !is_expected {
			return nil, errors.New("Неверно указаны скобки")
		}
		return node, nil
	}
	// Если токен на очереди не скобка, получаем узел с числом
	val := tokens.List[0]
	if _, err := strconv.Atoi(val); err != nil {
		fmt.Println(err)
		return nil, errors.New("В выражениии недопустимые символы")
	}
	tokens.List = tokens.List[1:] // Удаляем полученное из списка число
	return NewNode(val, nil, nil), nil
}

// Функция возвращает дерево умножения или деления
func (tokens *Tokens) getProductDivTree() (*Node, error) {
	val1, err := tokens.getNumberNode() // Узел с числом или подвыражением в скобках
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	token, is_expected := tokens.popExpectedToken("*", "/")
	if is_expected {
		// Рекурсивно получаем второе значение для умножения или деления
		val2, _ := tokens.getProductDivTree()
		return NewNode(token, val1, val2), nil
	}
	return val1, nil
}

// Функция собирает дерево выражений
func (tokens *Tokens) getExpTree() (*Node, error) {
	// Пробуем получить узел с произведением или делением
	// Если операция на очереди не умножение или деление, вернется узел с числом
	val1, err := tokens.getProductDivTree()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	token, is_expected := tokens.popExpectedToken("+", "-")
	// Если операция сложение или вычитание, рекурсивно получаем второе значение
	if is_expected {
		val2, _ := tokens.getExpTree()
		return NewNode(token, val1, val2), nil
	}

	return val1, nil
}
