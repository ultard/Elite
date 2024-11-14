package main

import (
	"fmt"
	"math"
)

func main() {
	var a1, a2, b1, b2, c1, c2 float64
	fmt.Scanf("%f %f %f %f %f %f", &a1, &a2, &b1, &b2, &c1, &c2)

	start := math.Max(math.Max(a1, b1), c1)
	end := math.Min(math.Min(a2, b2), c2)

	if start > end {
		fmt.Println("Пересечений нет")
	} else {
		fmt.Println(start, end)
	}
}
