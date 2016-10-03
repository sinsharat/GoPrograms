package main

import (
	"bufio"
	"fmt"
	"os"
)

func countDupLines(file *os.File, counter map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		if input.Text() != ">NIL" {
			counter[input.Text()]++
			continue
		}
		break
	}
}

func findDuplicateLines() {
	counter := make(map[string]int)
	fileNameList := os.Args[1:]
	if len(fileNameList) > 0 {
		for _, fileName := range fileNameList {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Errorf("Failed to open file : %v", file)
				continue
			}
			countDupLines(file, counter)
		}
	} else {
		fmt.Println("Enter the lines, when finished type in '>NIL'")
		countDupLines(os.Stdin, counter)
	}

	for key, value := range counter {
		fmt.Printf("Key : %v, Duplicate count : %v\n", key, value)
	}
}

func main() {
	findDuplicateLines()
}
