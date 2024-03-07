package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	var sum, sum2 int
	var lines []string

	file, err := os.Open("input")

	if err != nil {
		log.Fatal("Error while opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for _, line := range lines {
		firstPart := line[:len(line)/2]
		secondPart := line[len(line)/2:]

		common, _ := findCommonInTwoStrings(firstPart, secondPart)

		if unicode.IsUpper(common) {
			sum += int(common) - 38
		} else {
			sum += int(common) - 96
		}
	}

	for i := 0; i < len(lines); i += 3 {
		first := lines[i]
		second := lines[i+1]
		third := lines[i+2]

		common, _ := findCommonInThreeStrings(first, second, third)

		if unicode.IsUpper(common) {
			sum2 += int(common) - 38
		} else {
			sum2 += int(common) - 96
		}
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

func findCommonInTwoStrings(first string, second string) (rune, bool) {
	for _, l1 := range first {
		for _, l2 := range second {
			if l1 == l2 {
				return l1, true
			}
		}
	}
	return 0, false
}

func findCommonInThreeStrings(first string, second string, third string) (rune, bool) {
	for _, l1 := range first {
		for _, l2 := range second {
			if l1 == l2 {
				for _, l3 := range third {
					if l3 == l2 {
						return l1, true
					}
				}
			}
		}
	}

	return 0, false
}
