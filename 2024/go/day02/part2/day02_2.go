package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// 285

func main() {
	file, err := os.Open("./puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	safeReport := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		intFields := stringSliceToInts(fields)

		variances := generateVariances(intFields, -1)

		if isValid(variances) {
			safeReport++
		} else {
			if canDropOneLevel(intFields) {
				safeReport++
			}
		}
	}

	fmt.Printf("Safe reports: %v\n", safeReport)
}

func canDropOneLevel(fields []int) bool {
	for i := range fields {
		variances := generateVariances(fields, i)
		if isValid(variances) {
			return true
		}
	}

	return false
}

func generateVariances(fields []int, skipIndex int) []int {
	var newFields []int
	var variances []int

	for x := 0; x < len(fields); x++ {
		if skipIndex == x {
			continue
		}

		newFields = append(newFields, fields[x])
	}

	for x := 0; x < len(newFields)-1; x++ {
		diff := newFields[x] - newFields[x+1]
		variances = append(variances, diff)
	}

	return variances
}

func isValid(variances []int) bool {
	for i, l := range variances {
		if l == 0 || l > 3 || l < -3 {
			return false
		}

		if i > 0 && !sameSign(variances[i-1], variances[i]) {
			return false
		}
	}

	return true
}

func sameSign(a, b int) bool {
	return (a ^ b) >= 0
}

func stringSliceToInts(fields []string) []int {
	var ints []int

	for _, s := range fields {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, i)
	}

	return ints
}
