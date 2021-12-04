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
	cons := subConsumption(inputThree)
	log.Infof("Day 3 - Power consumption: %v", cons)
}
