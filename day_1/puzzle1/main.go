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
	fileName := os.Args[1]
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	fileText := string(file)
	count1 := 0
	for _, char := range fileText {
		if char == '(' {
			count1++
		}
	}
	count2 := 0
	for _, char := range fileText {
		if char == ')' {
			count2++
		}
	}
	floorNo := count1 - count2
	fmt.Println(floorNo)
}
