package arrays

func Sum(numbers []int) int {
	var sum int
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var results []int
	for _, numbers := range numbersToSum {
		results = append(results, Sum(numbers))
	}
	return results
}

func SumAllTails(numbersToSum ...[]int) []int {
	var results []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			results = append(results, 0)
			continue
		}
		trailNumbers := numbers[1:]
		results = append(results, Sum(trailNumbers))
	}
	return results
}
