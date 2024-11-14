package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func main() {
	var year int
	fmt.Print("Введите год: ")
	_, err := fmt.Scan(&year)
	if err != nil {
		fmt.Println("Пожалуйста, введите целое число.")
		return
	}

	if isLeapYear(year) {
		fmt.Printf("%d является високосным годом.\n", year)
	} else {
		fmt.Printf("%d не является високосным годом.\n", year)
	}
}
