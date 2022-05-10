package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	values := make([]string, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Sort(sort.IntSlice(keys))
	for idx, key := range keys {
		values[idx] = input[key]
	}
	return values
}
