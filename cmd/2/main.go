package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Функция для определения приоритета операторов
func precedence(operator rune) int {
	switch operator {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0 // для скобок
	}
}

// Функция для преобразования инфиксного выражения в постфиксное
func infixToPostfix(tokens []string) []string {
	var output []string
	var operatorStack []string

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			for len(operatorStack) > 0 && precedence(rune(operatorStack[len(operatorStack)-1][0])) >= precedence(rune(token[0])) {
				output = append(output, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = append(operatorStack, token)
		case "(":
			operatorStack = append(operatorStack, token)
		case ")":
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "(" {
				output = append(output, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = operatorStack[:len(operatorStack)-1] // Убираем "(" из стека
		default:
			output = append(output, token)
		}
	}

	for len(operatorStack) > 0 {
		output = append(output, operatorStack[len(operatorStack)-1])
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return output
}

// Проверка на баланс скобок
func areBracketsBalanced(expression string) bool {
	var stack []rune
	for _, char := range expression {
		switch char {
		case '(':
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// Проверка на валидность выражения
func isValidExpression(expression string) bool {
	if !areBracketsBalanced(expression) {
		return false
	}

	previousChar := ' '
	for i, char := range expression {
		if !unicode.IsDigit(char) && char != '+' && char != '-' && char != '*' && char != '/' && char != '(' && char != ')' && char != ' ' {
			return false
		}

		if char == '+' || char == '*' || char == '/' {
			if i == 0 || previousChar == '+' || previousChar == '-' || previousChar == '*' || previousChar == '/' || previousChar == '(' {
				return false
			}
		}

		if previousChar == '/' && char == '0' {
			return false
		}

		if char != ' ' {
			previousChar = char
		}
	}

	if previousChar == '+' || previousChar == '-' || previousChar == '*' || previousChar == '/' {
		return false
	}

	return true
}

// Функция для разбиения строки на токены
func splitExpression(expression string) ([]string, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	if !isValidExpression(expression) {
		return nil, fmt.Errorf("invalid expression")
	}

	var tokens []string
	token := ""
	for _, char := range expression {
		if char == ' ' {
			continue
		} else if char == '+' || char == '-' || char == '*' || char == '/' || char == '(' || char == ')' {
			if token != "" {
				tokens = append(tokens, token)
				token = ""
			}
			tokens = append(tokens, string(char))
		} else {
			token += string(char)
		}
	}
	if token != "" {
		tokens = append(tokens, token)
	}

	tokens = infixToPostfix(tokens)
	return tokens, nil
}

// Функция для вычисления постфиксного выражения
func computeExpression(tokens []string) (float64, bool) {
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
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		}
	}

	return stack[0], true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите выражение: ")
	scanner.Scan()
	expression := scanner.Text()

	if expression == "" {
		fmt.Println("Пожалуйста, введите выражение.")
	}

	tokens, err := splitExpression(expression)
	if err != nil {
		fmt.Println("Некорректное выражение.")
		return
	}

	result, ok := computeExpression(tokens)
	if !ok {
		fmt.Println("Ошибка: деление на ноль.")
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}
