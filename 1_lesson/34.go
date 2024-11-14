package main

import "fmt"

func main() {
	var text string
	fmt.Scanf("%s", &text)

	runeText := []rune(text)
	fmt.Println(string(reverseRunes(runeText)))
}

func reverseRunes(r []rune) []rune {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return r
}
