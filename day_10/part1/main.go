package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Function to generate the next number in the look-and-say sequence using strings.Builder
func lookAndSay(sequence string) string {
	var builder strings.Builder
	count := 1

	for i := 0; i < len(sequence); i++ {
		if i < len(sequence)-1 && (sequence[i] == sequence[i+1]) {
			count++
		} else {
			builder.WriteString(strconv.Itoa(count))
			builder.WriteByte(sequence[i])
			count = 1
		}
	}
	return builder.String()
}

func main() {
	input := "1321131112"
	iterations := 40

	for i := 0; i < iterations; i++ {
		input = lookAndSay(input)
	}

	fmt.Println(len(input))
}
