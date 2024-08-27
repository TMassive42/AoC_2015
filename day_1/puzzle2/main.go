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
	count := 0
	for i := 0; i < len(fileText); i++ {
		if fileText[i] == '(' {
			count++
		} else if fileText[i] == ')' {
			count--
		}
		if count == -1 {
			fmt.Println(i+1)
			return
		}
	}

}
