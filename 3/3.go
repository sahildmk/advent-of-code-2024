package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main1() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("error", err)
		return
	}

	// Regex pattern to match "mul(x,y)" with any numbers
	pattern := `mul\((-?\d+),(-?\d+)\)`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Split the string
	parts := re.FindAllStringSubmatch(string(content), -1)

	sum := 0
	for _, part := range parts {
		first, _ := strconv.Atoi(part[1])
		second, _ := strconv.Atoi(part[2])

		mul := first * second

		fmt.Println(part[1], "x", part[2], "=", mul)
		sum += mul
	}

	fmt.Println(sum)
}
