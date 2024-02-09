package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	var sum int

	file, err := os.Open("input")

	if err != nil {
		log.Fatal("Error while opening file")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		firstPart := line[:len(line)/2]
		secondPart := line[len(line)/2:]

		common, _ := findCommonInTwo(firstPart, secondPart)

		if unicode.IsUpper(common) {
			sum += int(common) - 38
		} else {
			sum += int(common) - 96
		}
	}

	fmt.Println(sum)
}

func findCommonInTwo(first string, second string) (rune, bool) {
	for _, l1 := range first {
		for _, l2 := range second {
			if l1 == l2 {
				return l1, true
			}
		}
	}
	return 0, false
}
