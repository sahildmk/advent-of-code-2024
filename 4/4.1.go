package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("error", err)
		return
	}

	grid := strings.Split(string(content), "\n")

	total := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			// fmt.Printf("%c", grid[row][col])
			char := grid[row][col]

			if char == 'X' {
				// Start searching
				fmt.Println("found X at", "(", row, ",", col, ")")
				total += searchWord2(grid, row, col)
			}
		}
	}
	fmt.Println(total)
}

func searchWord2(grid []string, row int, col int) int {
	XMAS := "XMAS"
	found := 0
	// right
	if col+3 < len(grid[row]) {
		str := grid[row][col : col+4]
		if str == XMAS {
			found += 1
		}
		fmt.Println("check right", str)
	}

	// left
	if col-3 >= 0 {
		str := reverseString(grid[row][col-3 : col+1])
		if str == XMAS {
			found += 1
		}
		fmt.Println("check left", str)
	}

	// bottom
	if row+3 < len(grid) {
		slice := []byte("")
		for i := 0; i < 4; i++ {
			slice = append(slice, grid[row+i][col])
		}
		str := string(slice)
		if str == XMAS {
			found += 1
		}
		fmt.Println("check bottom", str)
	}

	// top
	if row-3 >= 0 {
		slice := []byte("")
		for i := 0; i < 4; i++ {
			slice = append(slice, grid[row-i][col])
		}
		str := string(slice)
		if str == XMAS {
			found += 1
		}
		fmt.Println("check top", str)
	}

	// top left
	if row-3 >= 0 && col-3 >= 0 {
		slice := []byte("")
		for i := 0; i < 4; i++ {
			slice = append(slice, grid[row-i][col-i])
		}
		str := string(slice)
		if str == XMAS {
			found += 1
		}
		fmt.Println("check top left", str)
	}

	// top right
	if row-3 >= 0 && col+3 < len(grid[row]) {
		slice := []byte("")
		for i := 0; i < 4; i++ {
			slice = append(slice, grid[row-i][col+i])
		}
		str := string(slice)
		if str == XMAS {
			found += 1
		}
		fmt.Println("check top right", str)
	}

	// bottom left
	if row+3 < len(grid) && col-3 >= 0 {
		slice := []byte("")
		for i := 0; i < 4; i++ {
			slice = append(slice, grid[row+i][col-i])
		}
		str := string(slice)
		if str == XMAS {
			found += 1
		}
		fmt.Println("check bottom left")
	}

	// bottom right
	if row+3 < len(grid) && col+3 < len(grid[row]) {
		slice := []byte("")
		for i := 0; i < 4; i++ {
			slice = append(slice, grid[row+i][col+i])
		}
		str := string(slice)
		if str == XMAS {
			found += 1
		}
		fmt.Println("check bottom right")
	}

	fmt.Println(found)

	return found
}
