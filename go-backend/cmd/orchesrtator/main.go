package orchesrtator

import (
	"distributed-calculator/pkg/parser"
	"fmt"
)

func Orchestrator(exp string) {
	_, err := parser.ParseExpression(exp)
	fmt.Println(err)

}
