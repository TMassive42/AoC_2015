package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"log"
	"strconv"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parseDistance(line)
		
	}

	cities := []string{"AlphaCentauri", "Snowdin", "Tambi", "Faerun", "Norrath", "Straylight", "Tristram", "Arbre"}

	// Find and print the shortest route
	route, distance := findLongestRoute(cities)
	fmt.Printf("Longest route: %v\n", route)
	fmt.Printf("Longest distance: %d\n", distance)
}

// Global distance map
var distances = make(map[string]map[string]int)

// Function to parse each line and update the distances map
func parseDistance(line string) {
	parts := strings.Split(line, " = ")
	if len(parts) != 2 {
		log.Fatalf("invalid line format: %s", line)
	}

	cities := strings.Split(parts[0], " to ")
	if len(cities) != 2 {
		log.Fatalf("invalid cities format: %s", parts[0])
	}

	distance, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("invalid distance format: %s", parts[1])
	}

	// Add distance to the map (bidirectional)
	addDistance(cities[0], cities[1], distance)
	addDistance(cities[1], cities[0], distance)
}

// Helper function to add distances to the map
func addDistance(city1, city2 string, distance int) {
	if distances[city1] == nil {
		distances[city1] = make(map[string]int)
	}
	distances[city1][city2] = distance
}

// Helper function to generate all permutations of the cities
func permutations(cities []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(cities []string, n int) {
		if n == 1 {
			tmp := make([]string, len(cities))
			copy(tmp, cities)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(cities, n-1)
				if n%2 == 1 {
					cities[0], cities[n-1] = cities[n-1], cities[0]
				} else {
					cities[i], cities[n-1] = cities[n-1], cities[i]
				}
			}
		}
	}

	helper(cities, len(cities))
	return res
}

// Calculate the total distance for a given route
func calculateDistance(route []string) int {
	totalDistance := 0
	for i := 0; i < len(route)-1; i++ {
		totalDistance += distances[route[i]][route[i+1]]
	}
	return totalDistance
}

func findLongestRoute(cities []string) (longestRoute []string, longestDistance int) {
	perms := permutations(cities)
	longestDistance = math.MinInt
	for _, route := range perms {
		dist := calculateDistance(route)
		if dist > longestDistance {
			longestDistance = dist
			longestRoute = route
		}
	}
	return longestRoute, longestDistance
}
