package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// 183,380,722

func main() {
	file, err := os.Open("./puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sum = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches, err := returnMatches(line)
		if err != nil {
			log.Fatalln(err)
		}

		for _, match := range matches {
			left, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatalln(err)
			}

			right, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatalln(err)
			}

			sum += left * right
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v\n", err)
	}

	fmt.Printf("Sum: %v\n", sum)
}

func returnMatches(line string) ([][]string, error) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return [][]string{}, errors.New("no matches found")
	}

	return matches, nil
}
