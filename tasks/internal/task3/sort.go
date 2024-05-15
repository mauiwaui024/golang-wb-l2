package task3

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Sort(filePath string, column int, numeric, reverse, unique bool) error {
	// Открытие файла для чтения
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Считывание строк из файла
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// Функция для сравнения строк в соответствии с флагами
	comparator := func(i, j int) bool {
		line1 := strings.Fields(lines[i])
		line2 := strings.Fields(lines[j])

		// Выбор колонки для сортировки
		if column > 0 && column <= len(line1) && column <= len(line2) {
			line1 = []string{line1[column-1]}
			line2 = []string{line2[column-1]}
		}
		// Преобразование в числа, если нужно
		if numeric {
			num1, err1 := strconv.Atoi(line1[0])
			num2, err2 := strconv.Atoi(line2[0])
			if err1 == nil && err2 == nil {
				line1[0] = strconv.Itoa(num1)
				line2[0] = strconv.Itoa(num2)
			}
		}

		// Сравнение строк
		comparison := strings.Compare(line1[0], line2[0])
		// Учет обратного порядка
		if reverse {
			return comparison > 0
		}
		return comparison < 0
	}
	// Сортировка строк
	if unique {
		// Удаление дубликатов
		lines = removeDuplicates(lines)
	}
	// Применение кастомного компаратора
	sort.SliceStable(lines, comparator)
	// Вывод отсортированных строк
	for _, line := range lines {
		fmt.Println(line)
	}
	return nil
}

// Функция для удаления дубликатов из слайса строк
func removeDuplicates(lines []string) []string {
	uniqueLines := make(map[string]bool)
	var unique []string
	for _, line := range lines {
		if !uniqueLines[line] {
			unique = append(unique, line)
			uniqueLines[line] = true
		}
	}
	return unique
}
