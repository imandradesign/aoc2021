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
	inputLen := len(input)

	for i := range input {
		// All number lists to compare must contain at least 3 values
		if i < (inputLen - 3) {
			currentValue := input[i]
			oneNext := input[i+1]
			twoNext := input[i+2]
			threeNext := input[i+3]
			var currentList, nextList []int
			var currentSum, nextSum int

			currentList = append(currentList, currentValue, oneNext, twoNext)
			nextList = append(nextList, oneNext, twoNext, threeNext)

			currentSum = listSum(currentList)
			nextSum = listSum(nextList)

			if nextSum > currentSum {
				depthIncreaseCount++
			}
		}
	}

	fmt.Println(depthIncreaseCount)
}
