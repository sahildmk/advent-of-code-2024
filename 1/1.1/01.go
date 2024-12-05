package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Generic absolute value function
func abs[T int | int8 | int16 | int32 | int64 | float32 | float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Read the file
	content, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fileString := strings.Split(string(content), "\n")

	var left []int
	var right []int

	for i := 0; i < len(fileString); i++ {
		splitLine := strings.Split(fileString[i], "   ")

		leftInt, _ := strconv.Atoi(splitLine[0])
		rightInt, _ := strconv.Atoi(splitLine[1])

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	sort.Ints(left)
	sort.Ints(right)

	total := 0

	for i := 0; i < len(left); i++ {
		total += abs(left[i] - right[i])
	}

	fmt.Println(total)
}
