package main

import "fmt"

func isPrimary(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	for i := a; i <= b; i++ {
		if isPrimary(i) {
			fmt.Println(i)
		}
	}
}
