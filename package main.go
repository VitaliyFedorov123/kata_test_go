package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var romanToArabicMap = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var arabicToRomanMap = []struct {
	Value  int
	Symbol string
}{
	{10, "X"},
	{9, "IX"},
	{8, "VIII"},
	{7, "VII"},
	{6, "VI"},
	{5, "V"},
	{4, "IV"},
	{3, "III"},
	{2, "II"},
	{1, "I"},
}

func romanToArabic(roman string) (int, error) {
	roman = strings.ToUpper(roman)
	total := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		char := string(roman[i])
		value, exists := romanToArabicMap[char]
		if !exists {
			return 0, errors.New("недопустимая римская цифра")
		}

		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}

	return total, nil
}

func arabicToRoman(num int) (string, error) {
	if num <= 0 {
		return "", errors.New("цифра больжна быть больше 0")
	}

	result := ""
	for _, entry := range arabicToRomanMap {
		for num >= entry.Value {
			result += entry.Symbol
			num -= entry.Value
		}
	}
	return result, nil
}

func isArabic(input string) bool {
	for _, char := range input {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func parseInput(input string) (int, string, error) {
	if isArabic(input) {
		var num int
		_, err := fmt.Sscanf(input, "%d", &num)
		if err != nil {
			return 0, "", errors.New("ошибка при обработке арабского числа")
		}
		return num, "arabic", nil
	} else {
		num, err := romanToArabic(input)
		if err != nil {
			return 0, "", errors.New("ошибка при обработке римского числа")
		}
		return num, "roman", nil
	}
}

func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль невозможно")
		}
		return a / b, nil
	default:
		return 0, errors.New("неизвестный оператор")
	}
}

func main() {
	var input1, input2, operator string

	fmt.Println("Введите первое число (арабское или римское):")
	fmt.Scanln(&input1)
	fmt.Println("Введите оператор (+, -, *, /):")
	fmt.Scanln(&operator)
	fmt.Println("Введите второе число (арабское или римское):")
	fmt.Scanln(&input2)

	// Определяем тип и значение первого числа
	num1, type1, err := parseInput(input1)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	num2, type2, err := parseInput(input2)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if type1 != type2 {
		fmt.Println("Ошибка: оба числа должны быть либо арабскими, либо римскими")
		return
	}

	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if type1 == "roman" {
		romanResult, err := arabicToRoman(result)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Printf("Результат: %s\n", romanResult)
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}
