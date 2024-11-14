package main

import (
	"fmt"
	"math/rand"
)

func main() {

	answer := rand.Intn(100)

	for maxAttempts := 10; maxAttempts > 0; maxAttempts -= 1 {
		var guess int
		fmt.Print("Введите вашу догадку: ")
		fmt.Scan(&guess)

		if guess < 0 || guess > 100 {
			fmt.Println("Пожалуйста, введите число от 0 до 100.")
			continue
		}

		if guess < answer {
			fmt.Println("Загаданное число больше.")
		} else if guess > answer {
			fmt.Println("Загаданное число меньше.")
		} else {
			fmt.Printf("Поздравляем! Вы угадали число с %d попытки.\n", maxAttempts)
			return
		}
	}

	fmt.Printf("К сожалению, попытки закончились. Загаданное число было: %d\n", answer)
}
