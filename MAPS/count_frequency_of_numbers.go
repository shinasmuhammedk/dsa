package maps

func CountFrequencyOfNumbers(arr []int) map[int]int {
	count := make(map[int]int)

	for _, v := range arr {
		count[v]++
	}

	return count
}
