package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Pulls string data from a text file and returns it as an int array
func readInputTwo(fileName string) (fullList []string, words []string, nums []int, err error) {
	// Reads the input2.txt file
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, nil, nil, err
	}

	// Splits the strings from the file into separate items to be read based on new line
	lines := strings.Split(string(file), "\n")

	// Creates an int and string array the same length as the number of lines in the input2.txt file
	nums = make([]int, 0, len(lines))
	words = make([]string, 0, len(lines))

	// Loops through the separate strings from the input file, converts them to ints, and adds them to the nums array
	for i, l := range lines {
		if len(l) == 0 {
			continue
		}

		currentLine := lines[i]

		numList := strings.Split(string(currentLine), " ")[1]
		word := strings.Split(string(currentLine), " ")[0]
		num, err := strconv.Atoi(numList)
		if err != nil {
			log.WithError(err)
		}

		nums = append(nums, num)
		words = append(words, word)
	}

	return lines, words, nums, nil
}
