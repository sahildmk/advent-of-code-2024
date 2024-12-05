package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read the file
	content, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fileString := strings.Split(string(content), "\n")

	var left []int

	rightCounts := make(map[int]int)

	for i := 0; i < len(fileString); i++ {
		splitLine := strings.Split(fileString[i], "   ")

		leftInt, _ := strconv.Atoi(splitLine[0])
		rightInt, _ := strconv.Atoi(splitLine[1])

		left = append(left, leftInt)

		total, exists := rightCounts[rightInt]

		if exists {
			rightCounts[rightInt] = total + 1
		} else {
			rightCounts[rightInt] = 1
		}
	}

	total := 0

	for i := 0; i < len(left); i++ {
		count, exists := rightCounts[left[i]]

		if exists {
			total += left[i] * count
		}
	}

	fmt.Println(total)
}
