package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Pulls string data from the input2.txt file and returns 3 arrays & error
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

	// Loops through the separate strings from the input file to create separate arrays of the right types
	for i, l := range lines {
		if len(l) == 0 {
			log.Info("The 'lines' list was empty.")
			break
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
