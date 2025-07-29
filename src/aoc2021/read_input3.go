package main

import (
	"io/ioutil"
	"strings"
)

// Pulls string data from the input3.txt file.
func readInputThree(fileName string) (fullBinaryList []string, err error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Splits the strings from the file into separate items to be read based on new line
	binaryList := strings.Split(string(file), "\n")

	// Loops through the separate strings from the input file to create separate arrays of the right types
	for i, l := range binaryList {
		if len(l) == 0 {
			log.Info("The 'binaryList' list was empty.")
			break
		}

		currentLine := binaryList[i]
		fullBinaryList = append(fullBinaryList, currentLine)
	}

	return fullBinaryList, nil
}
