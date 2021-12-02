package main

func calculatePosition(list []string, words []string, nums []int) (mult int) {
	horizPosition := 0
	depth := 0

	for i := range list {
		currentWord := words[i]
		currentNum := nums[i]

		if currentWord == "forward" {
			horizPosition += currentNum
		}

		if currentWord == "up" {
			depth -= currentNum
		}

		if currentWord == "down" {
			depth += currentNum
		}
	}

	mult = horizPosition * depth

	return mult
}
