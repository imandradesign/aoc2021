package main

const (
	bitIndexMax = 11
)

// removeFromSlice removes a specific item from the slice if it doesn't match the criteria.
func removeFromSlice(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// returnCommonBit takes the current []string slice of possible values from the full puzzle input,
// loops through the values for a particular int index, and returns the most common and least common bit
func returnCommonBit(possibleResults []string, index int) (mostCommonBit, leastCommonBit int) {
	var zeroes, ones int
	for i := range possibleResults {
		currentLine := possibleResults[i]
		currentBit := currentLine[index]

		if currentBit == 0 {
			zeroes++
		}

		if currentBit == 1 {
			ones++
		}
	}

	if zeroes == ones || ones > zeroes {
		mostCommonBit = 1
		leastCommonBit = 0
	}

	if zeroes > ones {
		mostCommonBit = 0
		leastCommonBit = 1
	}

	log.Infof("mostCommonBit: %v, leastCommonBit: %v", mostCommonBit, leastCommonBit)

	return mostCommonBit, leastCommonBit
}

func narrowInputListsByIndex(remainingResultsMost []string, remainingResultsLeast []string, index int) ([]string, []string) {
	most, _ := returnCommonBit(remainingResultsMost, index)
	_, least := returnCommonBit(remainingResultsLeast, index)

	// Loop removes items from the remaining results if they don't match the MOST common bit
	for i := range remainingResultsMost {
		currentLine := remainingResultsMost[i]
		currentBit := int(currentLine[index])

		if currentBit == least {
			remainingResultsMost = removeFromSlice(remainingResultsMost, i)
		}
	}

	// Loop removes items from the remaining results if they don't match the LEAST common bit
	for i := range remainingResultsLeast {
		currentLine := remainingResultsLeast[i]
		currentBit := int(currentLine[index])

		if currentBit == most {
			remainingResultsLeast = removeFromSlice(remainingResultsLeast, i)
		}
	}

	return remainingResultsMost, remainingResultsLeast
}

// oxygenrating uses MOST common
// co2 rating uses LEAST common
func lifeSupportRating(input []string) int64 {
	var oxygenRatingBin, co2RatingBin []string
	possibleResultsMost, possibleResultsLeast := make([]string, len(input)), make([]string, len(input))
	mostCopied := copy(possibleResultsMost, input)
	leastCopied := copy(possibleResultsLeast, input)

	log.Infof("mostCopied: %v, leastCopied: %v", mostCopied, leastCopied)

	for i := 0; i <= bitIndexMax; i++ {
		mostLen := len(possibleResultsMost)
		leastLen := len(possibleResultsLeast)

		if (mostLen < 1) || (leastLen < 1) {
			log.Error("Somehow one or both of the possible results lists had all values removed! How odd.")
			break
		}

		if mostLen == 1 {
			oxygenRatingBin = possibleResultsMost
			log.Infof("mostLen is 1 | oxygenRatingBin: %v", oxygenRatingBin)
		}

		if mostLen > 1 {
			possibleResultsMost, _ = narrowInputListsByIndex(possibleResultsMost, possibleResultsLeast, i)
		}

		if leastLen == 1 {
			co2RatingBin = possibleResultsLeast
			log.Infof("leastLen is 1 | co2RatingBin: %v", co2RatingBin)
		}

		if leastLen > 1 {
			_, possibleResultsLeast = narrowInputListsByIndex(possibleResultsMost, possibleResultsLeast, i)
		}

		if (mostLen == 1) && (leastLen == 1) {
			break
		}
	}

	co2Rate := getDecimal(co2RatingBin)
	oxygenRate := getDecimal(oxygenRatingBin)
	rate := (co2Rate * oxygenRate)

	log.Infof("lifeSupportRating() func | possibleResultsMost: %v, possibleResultsLeast: %v, oxygenRatingBin: %v, co2RatingBin: %v, co2Rate: %v, oxygenRate: %v", possibleResultsMost, possibleResultsLeast, oxygenRatingBin, co2RatingBin, co2Rate, oxygenRate)

	return rate
}
