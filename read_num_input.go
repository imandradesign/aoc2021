package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readNumInput(fileName string) (nums []int, err error) {
	// Reads the input1.txt file
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Splits the strings from the file into separate items to be read based on new line
	lines := strings.Split(string(file), "\n")

	// Creates an int array the same length as the number of lines in the input1.txt file
	nums = make([]int, 0, len(lines))

	// Loops through the separate strings from the input file, converts them to ints, and adds them to the nums array
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}
