package main

import (
	"fmt"
	"sort"
)

func sortByAbsoluteValue(arr []int) []int {
	// Используем пользовательскую сортировку
	sort.Slice(arr, func(i, j int) bool {
		return abs(arr[i]) < abs(arr[j])
	})
	return arr
}

// Функция для получения абсолютного значения
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	arr := []int{-5, 3, -2, 8, -1, 4}

	fmt.Println("Исходный массив:", arr)
	sortedArr := sortByAbsoluteValue(arr)
	fmt.Println("Отсортированный массив:", sortedArr)
}
