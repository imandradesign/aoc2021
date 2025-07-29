package main

import (
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New().WithFields(logrus.Fields{
		"name": "cmd",
	})
)

func main() {
	// Input from Day 1
	inputOne, err := readInputOne("inputs/input1.txt")
	if err != nil {
		log.WithError(err).Error("Error with Day 1 input.")
	}

	// Depth increase count from Day 1
	depthCount := depthCounter(inputOne)
	log.Infof("Day 1 - Depth count: %v", depthCount)

	// Input from Day 2
	inputFullTwo, inputWordsTwo, inputNumsTwo, err := readInputTwo("inputs/input2.txt")
	if err != nil {
		log.WithError(err).Error("Error with Day 2 input")
	}

	// Multiplied result of final horizontal position and final depth
	mult := calculatePosition(inputFullTwo, inputWordsTwo, inputNumsTwo)
	log.Infof("Day 2 - Multiplied position by depth: %v", mult)

	// Input from Day 3
	inputThree, err := readInputThree("inputs/input3.txt")
	if err != nil {
		log.WithError(err).Error("Error with Day 3 input")
	}

	// Multiplied result of epsilon & gamma rates to calculate sub power consumption
	cons, gammaBinArray, epsilonBinArray := subConsumption(inputThree)
	log.Infof("Day 3 - Power consumption: %v", cons)
	log.Infof("gammaBinArray: %v epsilonBinArray: %v", gammaBinArray, epsilonBinArray)

	log.Info("Desired co2 binary: 011011110010 Desired oxygen binary: 100100001100")

	// Multiplied the "oxygen generator rating" by the "CO2 scrubber rating" to get life support rating.
	lifeSupport := lifeSupportRating(inputThree)

	log.Infof("Day 3 - Life support rating: %v", lifeSupport)
}
