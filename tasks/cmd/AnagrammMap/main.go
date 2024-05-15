package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	// fmt.Println("sorted string is: ", string(r))
	return string(r)
}

func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)
	for _, word := range words {
		// Приведение слова к нижнему регистру
		word = strings.ToLower(word)
		// Сортировка букв в слове для создания ключа
		sortedWord := sortString(word)
		// Добавление слова в соответствующее множество анаграмм
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	// Удаление множеств из одного элемента
	for key, value := range anagramMap {
		if len(value) == 1 {
			delete(anagramMap, key)
		} else {
			// Сортировка элементов внутри множества
			sort.Strings(value)
			anagramMap[key] = value
		}
	}

	return anagramMap
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagrams := findAnagrams(words)
	for key, value := range anagrams {
		fmt.Println(key, ":", value)
	}
}
