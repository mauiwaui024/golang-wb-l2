package main

import (
	"flag"
	"fmt"
	"tasks/internal/task3"
)

func main() {
	var (
		column  = flag.Int("k", 0, "Specify column for sorting")
		numeric = flag.Bool("n", false, "Sort by numeric value")
		reverse = flag.Bool("r", false, "Sort in reverse order")
		unique  = flag.Bool("u", false, "Do not output duplicate lines")
	)
	flag.Parse()
	filePath := flag.Arg(0)

	fmt.Println("file path is ", filePath)

	// Check if file path is provided
	if filePath == "" {
		fmt.Println("File path is required")
		return
	}

	err := task3.Sort(filePath, *column, *numeric, *reverse, *unique)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
