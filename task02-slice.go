package homework

func reverse(input []int64) (result []int64) {
	revSlice := make([]int64, len(input))
	for idx, value := range input {
		revSlice[len(input)-idx-1] = value
	}
	return revSlice
}
