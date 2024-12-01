package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Sample: 11
// Answer: 2057374

func main() {
	var listLeft []int
	var listRight []int
	var err error
	var l, r int

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

		l, err = strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		listLeft = append(listLeft, l)

		r, err = strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		listRight = append(listRight, r)
	}

	slices.Sort(listLeft)
	slices.Sort(listRight)

	var totalDistance int
	for i := range listLeft {
		totalDistance += int(math.Abs(float64(listLeft[i] - listRight[i])))
	}

	fmt.Printf("Total distance: %v\n", totalDistance)
}
