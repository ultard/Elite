package main

import (
	"fmt"
	"strings"
)

func charToInt(c rune) int {
	if '0' <= c && c <= '9' {
		return int(c - '0')
	} else if 'A' <= c && c <= 'Z' {
		return int(c - 'A' + 10)
	} else if 'a' <= c && c <= 'z' {
		return int(c - 'a' + 10)
	}
	return -1
}

func intToChar(i int) rune {
	if 0 <= i && i <= 9 {
		return rune(i + '0')
	} else if 10 <= i && i <= 35 {
		return rune(i - 10 + 'A')
	}
	return '?'
}

func toDecimal(number string, base int) int {
	result := 0
	for _, char := range number {
		result = result*base + charToInt(char)
	}
	return result
}

func fromDecimal(number int, base int) string {
	if number == 0 {
		return "0"
	}
	var result []rune
	for number > 0 {
		result = append([]rune{intToChar(number % base)}, result...)
		number /= base
	}
	return string(result)
}

func convertBase(number string, sourceBase, targetBase int) string {
	decimal := toDecimal(number, sourceBase)
	return fromDecimal(decimal, targetBase)
}

func main() {
	var number string
	var sourceBase, targetBase int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	fmt.Print("Введите исходную систему счисления: ")
	fmt.Scan(&sourceBase)
	fmt.Print("Введите конечную систему счисления: ")
	fmt.Scan(&targetBase)

	if sourceBase < 2 || sourceBase > 36 || targetBase < 2 || targetBase > 36 {
		fmt.Println("Система счисления должна быть в диапазоне от 2 до 36")
		return
	}

	converted := convertBase(strings.ToUpper(number), sourceBase, targetBase)
	fmt.Printf("Число в системе %d: %s\n", targetBase, converted)
}
