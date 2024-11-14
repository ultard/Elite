package main

import "fmt"

func isPalindrome(s string) bool {
	runeS := []rune(s)

	for i := 0; i < len(runeS)/2; i++ {
		if runeS[i] != runeS[len(runeS)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	var text string

	fmt.Scanf("%s", &text)
	fmt.Println(isPalindrome(text))
}
