package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	points := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	results := map[string]map[string]int{
		"X": {
			"A": 3,
			"B": 0,
			"C": 6,
		},
		"Y": {
			"A": 6,
			"B": 3,
			"C": 0,
		},
		"Z": {
			"A": 0,
			"B": 6,
			"C": 3,
		},
	}

	pointsPerResult := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	pointsToMatch := map[string]map[string]int{
		"A": {
			"X": 3,
			"Y": 1,
			"Z": 2,
		},
		"B": {
			"X": 1,
			"Y": 2,
			"Z": 3,
		},
		"C": {
			"X": 2,
			"Y": 3,
			"Z": 1,
		},
	}

	file, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	total, total2 := 0, 0

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		player := line[1]
		rival := line[0]

		total += points[player]
		for k, e := range results[player] {
			if k == rival {
				total += e
			}
		}

		total2 += pointsPerResult[player]
		total2 += pointsToMatch[rival][player]
	}

	fmt.Println(total)
	fmt.Println(total2)
}
