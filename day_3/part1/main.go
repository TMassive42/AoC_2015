package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter file name")
		return
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	fileText := string(file)

	// Initialize a grid with a sufficiently large size
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	startX, startY := 500, 500

	result := visitedHouses(fileText, grid, startX, startY)
	fmt.Println(result)
}

func visitedHouses(fileText string, grid [][]int, x, y int) int {
	count := 1
	for _, char := range fileText {
		switch char {
		case '^':
			y--
		case '>':
			x++
		case 'v':
			y++
		case '<':
			x--
		}

		if grid[x][y] == 0 {
			count++
		}
		grid[x][y]++
	}
	return count
}
