package main

import (
	"strconv"
	"strings"
)

// subConsumption takes a []string input from the input3.txt file, gets the gamma and epsilon values
// based on the most common bit for each index position, converts them to int64,
// then returns the multiplied value from those two numbers.
func subConsumption(input []string) (consumption int64, gammaBinArray []string, epsilonBinArray []string) {
	// Gets the gamma & epsilon values as strings for all 12 bit positions
	for i := 0; i <= 11; i++ {
		gammaStr, epsilonStr := compareCommonBits(input, i)
		gammaBinArray = append(gammaBinArray, gammaStr)
		epsilonBinArray = append(epsilonBinArray, epsilonStr)
	}

	gamma := getDecimal(gammaBinArray)

	epsilon := getDecimal(epsilonBinArray)

	// log.Infof("Gamma decimal: %v, Epsilon decimal: %v", gamma, epsilon)

	consumption = (gamma * epsilon)

	return consumption, gammaBinArray, epsilonBinArray
}

// compareCommonBits takes []string and int index arguments, counts the 0s and 1s for a the input index position,
// and returns gamma & epsilon strings based on the count.
func compareCommonBits(input []string, index int) (mostCommon string, leastCommon string) {
	var bitsArray []string
	var ones, zeros int

	for i := range input {
		currentLine := input[i]

		bitStr := strings.Split(string(currentLine), "")[index]

		bitsArray = append(bitsArray, bitStr)
	}

	for i := range bitsArray {
		currentBit := bitsArray[i]

		if currentBit == "0" {
			zeros++
		} else if currentBit == "1" {
			ones++
		}
	}

	if zeros > ones {
		mostCommon = "0"
		leastCommon = "1"
	} else if ones > zeros || ones == zeros {
		mostCommon = "1"
		leastCommon = "0"
	}

	return mostCommon, leastCommon
}

// Takes a []string argument, joins the elements to a single string, then converts the binary to decimal.
func getDecimal(input []string) int64 {
	var str string

	for i := range input {
		str += input[i]
	}

	decNum, err := strconv.ParseInt(str, 2, 64)

	if err != nil {
		log.WithError(err).Error("getDecimal(): Unable to convert []string to int and get decimal value.")
		return 0
	} else {
		return decNum
	}
}
