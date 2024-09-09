package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter file name")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file")
	}

	defer file.Close()
	totalOldCount := 0
	totalNewCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileText := scanner.Text()
		totalOldCount += len(fileText)
		newCount := newString(fileText)
		totalNewCount += newCount
	}
	fmt.Println(totalNewCount - totalOldCount)
}

func newString(str string) int {
	newStr := "\""

	for i := 0; i < len(str); i++ {
		if str[i] == '\\' || str[i] == '"' {
			newStr += "\\" + string(str[i])
		} else {
			newStr += string(str[i])
		}
	}
	newStr += "\""
	return len(newStr)
}