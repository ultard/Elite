package main

import (
	"fmt"
)

func main() {
	var start, end int

	fmt.Scanf("%d %d", &start, &end)

	fmt.Printf("Таблица умножения от %d до %d:\n", start, end)
	for i := 1; i <= 10; i++ {
		for j := start; j <= end; j++ {
			fmt.Printf("%d * %d = %2d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
