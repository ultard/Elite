package main

import "fmt"

func main() {
	var rows int
	fmt.Scan(&rows)

	arr := [][]int{{1}}
	for i := 1; i < rows; i++ {
		arr = append(arr, []int{1})
		for j := 1; j < i; j++ {
			arr[i] = append(arr[i], arr[i-1][j-1]+arr[i-1][j])
		}
		arr[i] = append(arr[i], 1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < rows-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
