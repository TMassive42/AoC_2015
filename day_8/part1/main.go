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
	}

	defer file.Close()
	totalCodeChars := 0
	totalMemoryChars := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileText := scanner.Text()
		totalCodeChars += len(fileText)
		memoryChars := getDifference(fileText)
		totalMemoryChars += memoryChars
	}
	fmt.Println(totalCodeChars - totalMemoryChars)
}

func getDifference(str string) int {
	str = str[1 : len(str)-1]

	hexchars := "0123456789abcdef"
	count := 0

	for i := 0; i < len(str); i++ {
		if str[i] == '\\' {
			if i+3 < len(str) && str[i+1] == 'x' && strings.Contains(hexchars, string(str[i+2])) && strings.Contains(hexchars, string(str[i+3])) {
			i += 3
			} else if i+1 < len(str) && str[i] == '\\' && str[i+1] == '\\' || str[i+1] == '"' {
			i++
			}
		}
		count++
	}
	return count
}