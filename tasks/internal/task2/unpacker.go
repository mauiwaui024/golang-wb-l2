package task2

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Unpacker(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("входная строка пуста")
	}
	var sb strings.Builder
	re := regexp.MustCompile("[0-9]+|[a-zA-Z]")
	matches := re.FindAllString(input, -1)

	//"a4 b c2 d5 e" => "aaaabccddddde"
	for i := range matches {
		if isDigit(matches[i]) {
			if i == 0 {
				return "", errors.New("первый символ строки - цифра")
			}
			count, _ := strconv.Atoi(matches[i])
			for j := 0; j < count; j++ {
				sb.WriteString(matches[i-1])
			}
		} else if !isDigit(matches[i]) {
			if i != len(matches)-1 && !isDigit(matches[i+1]) {
				sb.WriteString(matches[i])
			} else if i == len(matches)-1 {
				sb.WriteString(matches[i])
			}
		}
	}
	return sb.String(), nil
}

// Функция для проверки, является ли строка числом
func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

/// парсить цифру легко через стринг билдер, но парсить число чето не понял как, поэтому сделал через слайсы и регулярные

// package task2

// import (
// 	"errors"
// 	"strconv"
// 	"strings"
// )

// func Unpacker(input string) (string, error) {
// 	if len(input) == 0 {
// 		return "", errors.New("входная строка пуста")
// 	}

// 	var builder strings.Builder
// 	repeatChar := byte(0)

// 	for i := 0; i < len(input); i++ {
// 		char := input[i]
// 		// "a4 b c2 d5 e" => "aaaabccddddde"
// 		//a3 b2 c => aaabbc
// 		if char >= '0' && char <= '9' {
// 			if repeatChar == 0 {
// 				return "", errors.New("первый символ строки - цифра")
// 			}
// 			count, _ := strconv.Atoi(string(char))
// 			if count != 1 {
// 				count -= 1
// 			}
// 			for j := 0; j < count; j++ {
// 				builder.WriteByte(repeatChar)
// 			}
// 			repeatChar = 0
// 		} else {
// 			builder.WriteByte(char)
// 			repeatChar = char
// 		}
// 	}

// 	return builder.String(), nil
// }
