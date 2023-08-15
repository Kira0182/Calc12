package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var counter int

func main() {

	fmt.Println("Калькулятор запущен! ")

	numbers := readUsersNumbers()
	firstValueStr, signValueStr, secondValueStr := GetUserString(numbers)
	firstValue := ConvertString(firstValueStr)
	secondValue := ConvertString(secondValueStr)
	Validation(firstValue, secondValue)
	result := Calculate(firstValue, signValueStr, secondValue)
	ValidationResult(result)
}

func readUsersNumbers() string {

	var inputValue string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputValue = scanner.Text()
	return inputValue
}

func GetUserString(str string) (string, string, string) {

	s := strings.Split(str, " ")
	if len(s) == 3 {
		firstValue := s[0]
		signValue := s[1]
		secondValue := s[2]
		return firstValue, signValue, secondValue
	} else {
		panic("Ошибка, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор")
	}
}

func ConvertString(numeral string) int {

	var arabic, err = strconv.Atoi(numeral)
	if err != nil {
		counter++
		var roman int = Decode(numeral)
		return roman
	}
	return arabic
}

func Validation(firstValue int, secondValue int) {
	if firstValue > 10 || firstValue < 0 || secondValue > 10 || secondValue < 0 {
		panic("Ошибка, так как калькулятор принимает числа на вход от 1 до 10")
	}
}

func Decode(roman string) int {

	var decoder = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	first := decoder[rune(roman[0])]
	if first == 0 {
		panic("Ошибка, так как это не является числом/цифрой в интервале от 1 до 10 включительно, или вовсе не является числом/цифрой")
	}
	if len(roman) == 1 {
		return first
	}
	next := decoder[rune(roman[1])]
	if next > first {
		return (next - first) + Decode(roman[2:])
	}
	return first + Decode(roman[1:])
}

func Calculate(firstValue int, sign string, secondValue int) int {

	if sign == "+" {
		result := firstValue + secondValue
		return result
	} else if sign == "-" {
		result := firstValue - secondValue
		return result
	} else if sign == "/" {
		result := firstValue / secondValue
		return result
	} else if sign == "*" {
		result := firstValue * secondValue
		return result
	} else {
		panic("Ошибка, так как введен неизвестный оператор")
		return -1
	}
	// по-хорошему, наверное, генерировать ошибку при помощи errors.New(), но и без него, вроде, нормально.

}

func ValidationResult(result int) {

	if counter == 0 {
		fmt.Println(result)
	}
	if counter == 1 {
		panic("Ошибка, так как используются одновременно разные системы счисления.")
	}
	if counter == 2 {
		fmt.Println(IntegerToRoman(result))
	}
}

func IntegerToRoman(number int) string {

	if number < 0 {
		panic("Ошибка, так как в римской системе нет отрицательных чисел.")
	}
	conversions := []struct {
		value int
		digit string
	}{
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}
	return roman.String()
}

// Очень жду вашего ответа ^_^
