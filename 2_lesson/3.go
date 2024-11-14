package main

import (
	"fmt"
	"log"
)

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func main() {
	var year int
	fmt.Print("Введите год: ")
	_, err := fmt.Scan(&year)
	if err != nil {
		log.Fatalf("Ошибка ввода: %v", err)
	}

	if isLeapYear(year) {
		fmt.Printf("%d год является високосным.\n", year)
	} else {
		fmt.Printf("%d год не является високосным.\n", year)
	}
}
