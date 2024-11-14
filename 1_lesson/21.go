package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	var operator string

	fmt.Scanf("%d %s %d", &a, &operator, &b)

	switch operator {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	case "%":
		fmt.Println(a % b)
	case "**":
		fmt.Println(math.Pow(float64(a), float64(b)))
	default:
		fmt.Println("Неизвестный оператор")
	}
}
