package main

import (
	"fmt"
	"log"
	"strconv"
)

func getInput(prompt string) float64 {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatalf("Ошибка при вводе данных: %v", err)
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatalf("Ошибка при преобразовании ввода в число: %v", err)
	}
	return value
}

func main() {
	initialVelocity := getInput("Введите начальную скорость (v0) в м/с: ")
	acceleration := getInput("Введите ускорение (a) в м/с²: ")
	time := getInput("Введите время (t) в секундах: ")

	if time < 0 {
		log.Fatalf("Время не может быть отрицательным.")
	}

	finalPosition := initialVelocity*time + 0.5*acceleration*time*time
	fmt.Printf("Конечное положение тела: %.2f м\n", finalPosition)
}
