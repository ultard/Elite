package main

import (
	"fmt"
)

func toBinary(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""
	for n > 0 {
		binary = fmt.Sprintf("%d", n%2) + binary
		n /= 2
	}

	return binary
}

func main() {
	var num int
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Ошибка ввода. Пожалуйста, введите целое число.")
		return
	}

	if num < 0 {
		fmt.Printf("Двоичное представление числа %d: 1,%s\n", num, toBinary(-num))
	} else {
		fmt.Printf("Двоичное представление числа %d: 0,%s\n", num, toBinary(num))
	}
}
