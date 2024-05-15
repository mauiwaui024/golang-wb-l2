package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func grep(filename string, pattern string, options map[string]interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// Проверяем, нужно ли интерпретировать паттерн как фиксированную строку
	var useFixedPattern bool = options["F"].(bool)
	if !useFixedPattern {
		pattern = regexp.QuoteMeta(pattern)
	}
	scanner := bufio.NewScanner(file)
	var lineCount int = 0
	var lineNum int = 0
	var invertedCount int = 0 //если -c and -v
	var found bool = false
	var prevLines []string
	var nextLines []string
	var lowerLine string
	// будем читать построчно
	if options["C"].(int) > 0 {
		cValue := options["C"].(int)
		options["B"] = cValue
		options["A"] = cValue
	}
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		//нашли совпадение
		if options["i"].(bool) {
			pattern = strings.ToLower(pattern)
			lowerLine = strings.ToLower(line)
		}

		if strings.Contains(line, pattern) || strings.Contains(lowerLine, pattern) {
			lineCount++
			found = true
			if options["B"].(int) > 0 {
				for _, prevLine := range prevLines {
					fmt.Println(prevLine)
				}
			}
			//печатаем строчку с совпадением
			if !options["c"].(bool) && !options["v"].(bool) {
				if options["n"].(bool) {
					fmt.Printf("%d:%s\n", lineNum, line)
				} else {
					fmt.Printf("%s\n", line)
				}
			}
		}

		if !strings.Contains(line, pattern) {
			//если с и v то мы увеличиваем countline и ниче не печатаем
			//а если у нас v просто то печатаем
			if options["c"].(bool) && options["v"].(bool) {
				invertedCount++
			}
			if options["v"].(bool) && !options["c"].(bool) {
				fmt.Printf("%s\n", line)
			}
		}
		//если длина слайса равна флагу "B", то удаляем первый элемент слайса
		if len(prevLines) == options["B"].(int) && len(prevLines) != 0 {
			prevLines = prevLines[1:]
		}
		//сохранили строчку в слайс предыдущих
		prevLines = append(prevLines, line)
		//сохранили строчку в слайс следующих
		if found && !strings.Contains(line, pattern) {
			if len(nextLines) != options["A"].(int) && options["A"].(int) > 0 {
				nextLines = append(nextLines, line)
			}
		}
		//печатаем строчки
		if len(nextLines) == options["A"].(int) && options["A"].(int) > 0 && !options["c"].(bool) {
			found = false
			for _, nextLine := range nextLines {
				fmt.Println(nextLine)
			}
			nextLines = nextLines[:0]
		}
	}
	if options["c"].(bool) && !options["v"].(bool) {
		fmt.Println(lineCount)
	}
	if options["c"].(bool) && options["v"].(bool) {
		fmt.Println(invertedCount)
	}
	return nil
}

func main() {
	// Парсинг флагов
	before := flag.Int("B", 0, "Print N lines before matching")
	after := flag.Int("A", 0, "Print N lines after matching")
	context := flag.Int("C", 0, "Print N lines of output context")
	count := flag.Bool("c", false, "Print only a count of matched lines")
	ignoreCase := flag.Bool("i", false, "Ignore case distinctions")
	invert := flag.Bool("v", false, "Selected lines are those not matching any of the specified patterns")
	fixed := flag.Bool("F", false, "Interpret pattern as a fixed string")
	lineNum := flag.Bool("n", false, "Each output line is preceded by its relative line number in the file")
	flag.Parse()

	// Парсинг паттерна и имени файла из аргументов командной строки
	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	options := map[string]interface{}{
		"A": *after,
		"B": *before,
		"C": *context,
		"c": *count,
		"i": *ignoreCase,
		"v": *invert,
		"F": *fixed,
		"n": *lineNum,
	}

	if err := grep(filename, pattern, options); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
