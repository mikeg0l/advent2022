package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
)

func main() {
	var calories []int
	var tempSum int
	file, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			calories = append(calories, tempSum)
			tempSum = 0
			continue
		}

		val, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		tempSum += val
	}

	fmt.Println(slices.Max(calories))

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] < calories[j]
	})

	lastEl := calories[len(calories)-1]
	secondLastEl := calories[len(calories)-2]
	thirdLastEl := calories[len(calories)-3]

	fmt.Printf("%v", lastEl+secondLastEl+thirdLastEl)
}
