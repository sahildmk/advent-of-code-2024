package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	// "strconv"
)

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("error", err)
		return
	}

	// Regex pattern to match "mul(x,y)" with any numbers
	pattern := `mul\((-?\d+),(-?\d+)\)|do\(\)|don't\(\)`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Split the string
	parts := re.FindAllStringSubmatch(string(content), -1)

	sum := 0
	do := true
	for _, part := range parts {
		start := part[0]

		if strings.HasPrefix(start, "mul") && do {
			first, _ := strconv.Atoi(part[1])
			second, _ := strconv.Atoi(part[2])

			mul := first * second
			sum += mul
		} else if strings.HasPrefix(start, "do()") {
			do = true
		} else if strings.HasPrefix(start, "don't") {
			do = false
		}
	}

	fmt.Println(sum)
}
