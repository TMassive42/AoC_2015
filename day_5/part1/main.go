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

func IsNice(str string) bool {
	if strings.Contains(str, "ab") || strings.Contains(str, "cd") || strings.Contains(str, "pq") || strings.Contains(str, "xy") {
		return false
	}

	count := 0
	hasDouble := false
	for i := 0; i < len(str); i++ {
		if IsVowel(rune(str[i])) {
			count++
		}

		if i < len(str)-1 && str[i] == str[i+1] {
			hasDouble = true
		}
	}
	return count >= 3 && hasDouble
}

func IsVowel (c rune) bool {
	return ContainRune("aeiou", c)
}

func ContainRune(str string, r rune) bool {
	for _, char := range str {
		if char == r {
			return true
		} 
	}
	return false
}