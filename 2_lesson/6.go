package main

import (
	"fmt"
)

func fibonacciIterative(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

func fibonacciRecursive(n int, memo map[int]int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if val, exists := memo[n]; exists {
		return val
	}

	memo[n] = fibonacciRecursive(n-1, memo) + fibonacciRecursive(n-2, memo)
	return memo[n]
}

func main() {
	var N int

	fmt.Print("Введите количество чисел Фибоначчи для вывода: ")
	_, err := fmt.Scan(&N)
	if err != nil || N < 0 {
		fmt.Println("Некорректный ввод. Пожалуйста, введите неотрицательное целое число.")
		return
	}

	fmt.Printf("Первые %d чисел Фибоначчи (итеративный метод):\n", N)
	iterativeFib := fibonacciIterative(N)
	for _, num := range iterativeFib {
		fmt.Print(num, " ")
	}
	fmt.Println()

	// Рекурсивный метод
	fmt.Printf("Первые %d чисел Фибоначчи (рекурсивный метод):\n", N)
	memo := make(map[int]int)
	for i := 0; i < N; i++ {
		fmt.Print(fibonacciRecursive(i, memo), " ")
	}
	fmt.Println()
}
