package main

import (
	"fmt"
	"os"
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

func isSafe(line []string) bool {
	fmt.Println("Checking: ", line)
	dir := 0
	for j := 0; j < len(line)-1; j++ {
		curr, _ := strconv.Atoi(line[j])
		next, _ := strconv.Atoi(line[j+1])

		diff := next - curr
		absDiff := abs(diff)

		if absDiff < 1 || absDiff > 3 {
			return false
		}

		if dir == 0 {
			dir = diff
			continue
		}

		// asc
		if diff > 0 && dir > 0 {
			continue
		}

		// desc
		if diff < 0 && dir < 0 {
			continue
		}

		return false
	}

	return true
}

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fileString := strings.Split(string(content), "\n")

	safeCount := 0
	for i := 0; i < len(fileString); i++ {
		line := strings.Split(fileString[i], " ")

		safe := isSafe(line)
		for j := 0; j < len(line) && !safe; j++ {
			copied := make([]string, len(line))
			copy(copied, line)
			copied = append(copied[:j], copied[j+1:]...)

			if isSafe(copied) {
				safe = true
				break
			}
		}

		if safe {
			fmt.Println("safe")
			safeCount++
		} else {
			fmt.Println("unsafe")
		}
	}

	fmt.Println(safeCount)
}

/*

1 2 3 4 5

2 1 3 4 5
^
1 2 3 4 3
        ^
1 2 1 4 5
    ^
*/
