package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Sample: 31
// Answer: 23,177,084

func main() {
	var listLeft []int
	var listRight []int

	file, err := os.Open("./puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			log.Fatalf("Invalid number of fields. Want 2 got %v", len(fields))
		}

		l, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		listLeft = append(listLeft, l)

		r, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		listRight = append(listRight, r)
	}

	var occurances = make(map[int]int)
	for _, rv := range listRight {
		occurances[rv]++
	}

	var similarityScore int
	for _, lv := range listLeft {
		similarityScore += occurances[lv] * lv
	}

	fmt.Printf("Similarity score: %v\n", similarityScore)
}
