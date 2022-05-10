package homework

func average(input [15]float32) (result float32) {
	var total float32 = 0
	for _, value := range input {
		total += value
	}
	return total / float32(len(input))
}
