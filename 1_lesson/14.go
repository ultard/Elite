package main

import (
	"fmt"
	"math"
)

func main() {
	arr1 := []int{1, 3, 5, 7, 9}
	arr2 := []int{0, 2, 4, 6, 8, 10, 1922}
	mergedArr := []int{}

	for i := 0; i < int(math.Max(float64(len(arr1)), float64(len(arr2)))); i++ {
		if i >= len(arr1) && i < len(arr2) {
			mergedArr = append(mergedArr, arr2[i])
			continue
		} else if i < len(arr1) && i >= len(arr2) {
			mergedArr = append(mergedArr, arr1[i])
			continue
		}

		if arr1[i] > arr2[i] {
			mergedArr = append(mergedArr, arr2[i], arr1[i])
		} else {
			mergedArr = append(mergedArr, arr1[i], arr2[i])
		}
	}

	fmt.Println(mergedArr)
}
