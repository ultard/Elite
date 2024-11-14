package main

import "fmt"

func main() {
	var maxNumber int
	fmt.Scanln(&maxNumber)

	fib := []int{0, 1}
	for i := 2; i < maxNumber; i++ {
		a, b := fib[i-1], fib[i-2]

		if a+b > maxNumber {
			break
		}

		fib = append(fib, a+b)
	}

	fmt.Println(fib)
}
