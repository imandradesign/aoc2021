package main

import (
	"fmt"
)

func main() {
	input, err := readNumInput("input1.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	depthIncreaseCount := 0

	for i, _ := range input {
		currentValue := input[i]
		// We want to skip the first input since there's nothing to compare it to
		if i > 0 {
			if currentValue > input[i-1] {
				depthIncreaseCount++
			}
		}
	}

	fmt.Println(depthIncreaseCount)
}
