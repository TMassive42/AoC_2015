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
	grid := make(map[[2]int]int)

	santaX, santaY := 0, 0
	roboX, roboY := 0, 0

	// Mark the starting position as visited
	grid[[2]int{santaX, santaY}] = 1

	for i, char := range fileText {
		if i%2 == 0 {
			switch char {
			case '^':
				santaY++
			case '>':
				santaX++
			case 'v':
				santaY--
			case '<':
				santaX--
			}
			grid[[2]int{santaX, santaY}]++
		} else {
			switch char {
			case '^':
				roboY++
			case '>':
				roboX++
			case 'v':
				roboY--
			case '<':
				roboX--
			}
			grid[[2]int{roboX, roboY}]++
		}
	}

	// Count the number of unique houses visited
	uniqueHouses := len(grid)
	fmt.Println(uniqueHouses)
}
