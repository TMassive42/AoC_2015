package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter file name")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	defer file.Close()
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		file := scanner.Text()
		fileText := string(file)

		
		if IsNice(fileText) {
			count++
		}
	}
	fmt.Println(count)
	
}


func IsNice(s string) bool {
	// Check if the string contains a pair of any two letters that appears at least twice without overlapping
	hasPair := false
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		// Check if this pair appears again later in the string, starting from i+2 to avoid overlap
		if strings.Contains(s[i+2:], pair) {
			hasPair = true
		}
	}
	
	hasDouble := false
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			hasDouble = true
		}
	}
	return hasPair && hasDouble
}
