package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
}
