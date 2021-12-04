package main

import (
	"io/ioutil"
	"strconv"
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

// Takes []string and int index arguments, counts the 0s and 1s for a the input index position, and returns gamma & epsilon strings based on the count.
func compareBits(input []string, index int) (gamma string, epsilon string) {
	var bitsArray []int
	var ones, zeros int

	for i := range input {
		currentLine := input[i]

		bitStr := strings.Split(string(currentLine), "")[index]

		bitInt, err := strconv.Atoi(bitStr)
		if err != nil {
			log.WithError(err)
		}

		bitsArray = append(bitsArray, bitInt)
	}

	for i := range bitsArray {
		currentBit := bitsArray[i]

		if currentBit == 0 {
			zeros++
		} else if currentBit == 1 {
			ones++
		}
	}

	if zeros > ones {
		gamma = "0"
		epsilon = "1"
	} else if ones > zeros {
		gamma = "1"
		epsilon = "0"
	}

	return gamma, epsilon
}

// Takes a []string argument, joins the elements to a single string, then converts the binary to decimal.
func getDecimal(input []string) (int64, error) {
	var str string

	for i := range input {
		str += input[i]
	}

	decNum, err := strconv.ParseInt(str, 2, 64)

	if err != nil {
		log.WithError(err).Error("Error joining []int to single integer using joinInts().")
		return 0, err
	} else {
		return decNum, nil
	}
}
