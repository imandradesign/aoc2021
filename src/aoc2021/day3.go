package main

func subConsumption(input []string) (consumption int64) {
	var gamma, epsilon int64
	var gammaBinArray, epsilonBinArray []string

	// Gets the gamma & epsilon values as strings for all 12 bit positions
	for i := 0; i <= 11; i++ {
		gammaStr, epsilonStr := compareBits(input, i)
		gammaBinArray = append(gammaBinArray, gammaStr)
		epsilonBinArray = append(epsilonBinArray, epsilonStr)
	}

	// Takes the []string values, joins them to a single string, then converts them to decimal
	gamma, err := getDecimal(gammaBinArray)
	if err != nil {
		log.WithError(err).Error("Unable to join Gamma array into int.")
	}

	epsilon, err = getDecimal(epsilonBinArray)
	if err != nil {
		log.WithError(err).Error("Unable to join Epsilon array into int.")
	}

	mult := (gamma * epsilon)

	return mult
}
