package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter file name")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	defer file.Close()

	totalAreaSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dimentions := string(line)
		parts := strings.Split(dimentions,"x")
		if len(parts) != 3 {
			fmt.Println("Invalid number of parts")
			return
		}
		length,_ := strconv.Atoi(parts[0])
		width,_ := strconv.Atoi(parts[1])
		height,_ := strconv.Atoi(parts[2])

		surfaceArea := (2*length*width) + (2*width*height) + (2*height*length)
		result := []int{
			length*width,
			width*height,
			height*length,
		}
		min := result[0]
		for _, char := range result {
			if char < min {
				min = char
			}
		}
		extraArea := min
		totalArea := surfaceArea + extraArea
		totalAreaSum += totalArea
	}
	fmt.Println(totalAreaSum)
}