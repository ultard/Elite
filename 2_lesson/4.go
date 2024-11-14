package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scanf("%f %f %f", &a, &b, &c)
	x1, x2 := findRoots(a, b, c)

	fmt.Printf("Корни уравнения: %.2f и %.2f\n", x1, x2)
}

func findRoots(a, b, c float64) (complex128, complex128) {
	if a == 0 {
		fmt.Println("Коэффициент a не может быть равным 0.")
		return 0, 0
	}

	d := b*b - 4*a*c

	if d > 0 {
		root1 := (-b + math.Sqrt(d)) / (2 * a)
		root2 := (-b - math.Sqrt(d)) / (2 * a)
		return complex(root1, 0), complex(root2, 0)
	} else if d == 0 {
		root := -b / (2 * a)
		return complex(root, 0), complex(root, 0)
	} else {
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(-d) / (2 * a)
		root1 := complex(realPart, imaginaryPart)
		root2 := complex(realPart, -imaginaryPart)
		return root1, root2
	}
}
