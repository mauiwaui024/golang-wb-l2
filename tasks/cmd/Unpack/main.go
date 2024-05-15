package main

import (
	"fmt"
	"tasks/internal/task2"
)

func main() {
	testString := "a45"
	// testString := "aaaabccddddde"
	strin1, err := task2.Unpacker(testString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strin1)
}
