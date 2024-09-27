package slices

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)

	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)

	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(numbers[1:])
		}
	}

	return sums
}
