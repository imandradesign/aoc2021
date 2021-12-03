package main

func calculatePosition(list []string, words []string, nums []int) (mult int) {
	var horizPosition, depth, aim int

	for i := range list {
		currentWord := words[i]
		currentNum := nums[i]

		if currentWord == "forward" {
			horizPosition += currentNum
			// log.Infof("Depth: %v. Aim: %v. Current num: %v.", depth, aim, currentNum)
			depth += (aim * currentNum)
			// log.Infof("Depth after mult: %v", depth)
		}

		if currentWord == "up" {
			aim -= currentNum
		}

		if currentWord == "down" {
			aim += currentNum
		}
	}

	mult = (horizPosition * depth)

	log.Infof("Final horizontal position: %v", horizPosition)
	log.Infof("Final depth: %v", depth)
	log.Infof("Final aim: %v", aim)

	return mult
}
