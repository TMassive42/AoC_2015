package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter file name")
		return
	}

	grid := make([][]bool, 1000)
	for i :=range grid {
		grid[i] = make([]bool, 1000)
	}


	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileText := scanner.Text()
		execGrid(fileText, grid)
	
	}
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func execGrid(str string, grid [][]bool) {
	parts := strings.Split(str, " ")
	
	var action string
	var  start, end string
	if parts[0] == "toggle" {
		action = "toggle"
		start = parts[1]
		end = parts[3]
	} else {
		action = parts[1]
		start = parts[2]
		end = parts[4]
	}

	startCoordinates := strings.Split(start, ",")
	endCoordinates := strings.Split(end, ",")
	startX, _ := strconv.Atoi(startCoordinates[0])
	endX, _ := strconv.Atoi(endCoordinates[0])
	startY, _ := strconv.Atoi(startCoordinates[1])
	endY, _ := strconv.Atoi(endCoordinates[1])

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			switch action {
			case "on":
				grid[i][j] = true
			case "off":
				grid[i][j] = false
			case "toggle":
				grid[i][j] = !grid[i][j]
			}
		}
	}
}