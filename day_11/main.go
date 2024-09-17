package main

import (
	"fmt"
	"strings"
)

// Increment password
func incrementPassword(pwd string) string {
	newPwd := []rune(pwd)
	for i := len(newPwd) - 1; i >= 0; i-- {
		if newPwd[i] == 'z' {
			newPwd[i] = 'a'
		} else {
			newPwd[i]++
			break
		}
	}
	return string(newPwd)
}

// Check if the password contains any forbidden characters
func containsInvalidChars(pwd string) bool {
	return strings.ContainsAny(pwd, "iol")
}

// Check for an increasing straight of at least three letters
func hasStraight(pwd string) bool {
	for i := 0; i < len(pwd)-2; i++ {
		if pwd[i+1] == pwd[i]+1 && pwd[i+2] == pwd[i]+2 {
			return true
		}
	}
	return false
}

// Check for at least two different non-overlapping pairs of letters
func hasTwoPairs(pwd string) bool {
	pairs := make(map[rune]bool)
	i := 0
	for i < len(pwd)-1 {
		if pwd[i] == pwd[i+1] {
			pairs[rune(pwd[i])] = true
			i += 2
		} else {
			i++
		}
	}
	return len(pairs) >= 2
}

// Check if password meets all the criteria
func isValidPassword(pwd string) bool {
	return !containsInvalidChars(pwd) && hasStraight(pwd) && hasTwoPairs(pwd)
}

func findNextPassword(pwd string) string {
	for {
		pwd = incrementPassword(pwd)
		if isValidPassword(pwd) {
			return pwd
		}
	}
}

func main() {
	currentPassword := "cqjxxyzz"
	nextPassword := findNextPassword(currentPassword)
	fmt.Println(nextPassword)
}
