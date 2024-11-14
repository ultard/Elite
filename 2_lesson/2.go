package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func computeRPN(tokens []string) (float64, bool) {
	var stack []float64

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, operand1+operand2)
			case "-":
				stack = append(stack, operand1-operand2)
			case "*":
				stack = append(stack, operand1*operand2)
			case "/":
				if operand2 == 0 {
					return 0, false
				}
				stack = append(stack, operand1/operand2)
			}
		default:
			num, _ := strconv.ParseFloat(token, 32)
			stack = append(stack, num)
		}
	}

	return stack[0], true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение в обратной польской записи: ")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	tokens := strings.Fields(expression)
	result, ok := computeRPN(tokens)
	if !ok {
		log.Fatalf("Ошибка: некорректное выражение")
	}

	fmt.Printf("Результат: %.2f\n", result)
}
