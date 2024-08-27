package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter file name")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file %s", err)
		return
	}

	defer file.Close()

	totalFeetSum := 0
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

		vals := []int{
			length,
			width,
			height,
		}

		sort.Ints(vals)
		ribbonVal := 2*(vals[0]) + 2*(vals[1])
		bowVal := length * width *height
		totalFeet := ribbonVal + bowVal
		totalFeetSum += totalFeet
	}
	fmt.Println(totalFeetSum)
}