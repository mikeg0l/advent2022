package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lines []string
	sum, sum2 := 0, 0

	file, err := os.Open("input")

	if err != nil {
		log.Fatal("Error occured while opening the file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		pairs := strings.Split(line, ",")
		firstRange := strings.Split(pairs[0], "-")
		secondRange := strings.Split(pairs[1], "-")

		firstRangeLeft, _ := strconv.Atoi(firstRange[0])
		firstRangeRight, _ := strconv.Atoi(firstRange[1])
		secondRangeLeft, _ := strconv.Atoi(secondRange[0])
		secondRangeRight, _ := strconv.Atoi(secondRange[1])

		if (firstRangeLeft <= secondRangeLeft && firstRangeRight >= secondRangeRight) ||
			(firstRangeLeft >= secondRangeLeft && firstRangeRight <= secondRangeRight) {
			sum++
		}

		if firstRangeRight >= secondRangeLeft || firstRangeLeft <= secondRangeRight {
			sum2++
		}
	}

	fmt.Println(sum)
}
