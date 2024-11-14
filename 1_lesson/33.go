package main

import (
	"fmt"
	"math"
)

func isArmstrong(n int) bool {
	sum := 0
	initial := n
	length := len(fmt.Sprint(n))

	for n > 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(length)))
		n /= 10
	}

	return sum == initial
}

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	for i := a; i <= b; i++ {
		if isArmstrong(i) {
			fmt.Println(i)
		}
	}
}
