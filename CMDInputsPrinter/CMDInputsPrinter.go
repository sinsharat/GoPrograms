package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func readCompares(r *bufio.Reader) {
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("err occured while reading string : ", err)
		}

		r, _ := utf8.DecodeRuneInString(line)

		fmt.Printf("input line : \"%v\", rune : \"%v\", %v\n", line, r, len(line))
		if len(line) == 1 || r == 13 {
			fmt.Println("found match for new line.", len(line))
			break
		}

		if line == ">END" {
			break
		}

		// remove trialling \n
		line = line[:len(line)-1]
		/*cmp, err := parseCompare(line)
		if err != nil {
			ExitWithError(ExitInvalidInput, err)
		}
		cmps = append(cmps, *cmp)*/
	}
}

func main() {
	fmt.Println("Provide inputs when done type in >END\n")
	readCompares(bufio.NewReader(os.Stdin))
}
