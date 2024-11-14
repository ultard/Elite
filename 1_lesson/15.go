package main

import "fmt"

func findSubstring(str, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	if len(substr) > len(str) {
		return -1
	}

	runeStr := []rune(str)
	runeSubstr := []rune(substr)

	for i := 0; i <= len(runeStr)-len(runeSubstr); i++ {
		match := true
		for j := 0; j < len(runeSubstr); j++ {
			if runeStr[i+j] != runeSubstr[j] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}

	return -1
}

func main() {
	var str, substr string
	fmt.Scanln(&str)
	fmt.Scanln(&substr)

	result := findSubstring(str, substr)
	if result == -1 {
		fmt.Println("Субстрока не найдена")
	} else {
		fmt.Printf("Субстрока найдена под индексом: %d\n", result)
	}
}
