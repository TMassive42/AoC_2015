package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var wires = make(map[string]uint16)
var instructions = make(map[string]string)

// Evaluate the signal for a wire or return a direct numeric value
func Evaluate(wire string) uint16 {
    if signal, err := strconv.Atoi(wire); err == nil {
        return uint16(signal)
    }

    if signal, exists := wires[wire]; exists {
        return signal
    }

    // Get the instruction for this wire
    parts := strings.Fields(instructions[wire])
    var result uint16

    switch len(parts) {
    case 1:
        result = getSignal(parts[0])
    case 2:
        result = ^getSignal(parts[1])
    case 3:
        left, right := getSignal(parts[0]), getSignal(parts[2])
        switch parts[1] {
        case "AND":
            result = left & right
        case "OR":
            result = left | right
        case "LSHIFT":
            shift := getSignal(parts[2])
            result = left << shift
        case "RSHIFT":
            shift := getSignal(parts[2])
            result = left >> shift
        }
    }

    wires[wire] = result
    return result
}

func getSignal(input string) uint16 {
    if signal, err := strconv.Atoi(input); err == nil {
        return uint16(signal)
    }
    return Evaluate(input)
}

func ParseInstructions(lines []string) {
    for _, line := range lines {
        parts := strings.Split(line, " -> ")
        instructions[parts[1]] = parts[0]
    }
}

func main() {
    file, err := os.Open("data.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var instructionLines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        instructionLines = append(instructionLines, scanner.Text())
    }

    ParseInstructions(instructionLines)

    // Calculate the signal for wire "a"
    signal := Evaluate("a")
    fmt.Printf("The signal provided to wire 'a' is: %d\n", signal)
}
