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
	inputOne, err := readInputOne("input1.txt")
	if err != nil {
		log.WithError(err).Error("Error with Day 1 input.")
	}

	// Depth increase count from Day 1
	depthCount := depthCounter(inputOne)
	log.Infof("Day 1 - Depth count: %v", depthCount)

	// Input from Day 2
	inputFullTwo, inputWordsTwo, inputNumsTwo, err := readInputTwo("input2.txt")
	if err != nil {
		log.WithError(err).Error("Error with Day 2 input")
	}

	// Multiplied result of final horizontal position and final depth
	mult := calculatePosition(inputFullTwo, inputWordsTwo, inputNumsTwo)
	log.Infof("Multiplied position by depth: %v", mult)
}
