package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Scanf("%d", &size)

	middle := size / 2
	for i := 0; i <= middle; i++ {
		for j := 0; j < middle-i; j++ {
			fmt.Print(" ")
		}
		// Печатаем звезды
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := middle - 1; i >= 0; i-- {
		for j := 0; j < middle-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
