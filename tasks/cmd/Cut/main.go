package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Определение флагов командной строки
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	// Разделитель по умолчанию
	defaultDelimiter := "\t"

	// Используем разделитель по умолчанию, если не указан другой
	if *delimiter != "" {
		// Извлекаем первый символ из значения флага -d
		defaultDelimiter = string((*delimiter)[0])
	}

	// Разбор запрошенных полей
	fieldIndices := make(map[int]bool)
	if *fields != "" {
		fieldsList := strings.Split(*fields, ",")
		for _, field := range fieldsList {
			fieldIndex := atoi(field)
			fieldIndices[fieldIndex-1] = true
		}
	}

	// Чтение ввода построчно
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Проверка, содержит ли строка разделитель
		if !strings.Contains(line, defaultDelimiter) {
			if *separated {
				continue // Пропускаем строки без разделителя, если указан флаг -s
			}
			// Иначе выводим строку целиком
			fmt.Println(line)
			continue
		}

		// Разделение строки на поля
		fields := strings.Split(line, defaultDelimiter)

		// Формирование результата
		var resultFields []string
		for index, field := range fields {
			// Если поле запрошено, добавляем его в результат
			if len(fieldIndices) == 0 || fieldIndices[index] {
				resultFields = append(resultFields, field)
			}
		}

		// Вывод результата
		fmt.Println(strings.Join(resultFields, defaultDelimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка при чтении стандартного ввода:", err)
		os.Exit(1)
	}
}

func atoi(s string) int {
	result := 0
	for _, c := range s {
		result = result*10 + int(c-'0')
	}
	return result
}
