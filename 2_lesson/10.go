package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)
	// Вычисление НОД
	result := gcd(a, b)
	fmt.Printf("Наибольший общий делитель чисел %d и %d: %d\n", a, b, result)
}
