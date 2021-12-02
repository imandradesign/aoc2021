package main

func calculatePosition(list []string, words []string, nums []int) (mult int) {
	var horizPosition, depth, aim int

	for i := range list {
		currentWord := words[i]
		currentNum := nums[i]

		if currentWord == "forward" {
			horizPosition += currentNum
			depth += (aim * currentNum)
		}

		if currentWord == "up" {
			depth -= currentNum
			aim -= currentNum
		}

		if currentWord == "down" {
			depth += currentNum
			aim += currentNum
		}
	}

	mult = horizPosition * depth

	log.Infof("Final horizontal position: %v", horizPosition)
	log.Infof("Final depth: %v", depth)
	log.Infof("Final aim: %v", aim)

	return mult
}
