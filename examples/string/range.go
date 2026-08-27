package stringexample

func RuneByteIndexes(value string) []int {
	indexes := make([]int, 0, len(value))
	for index := range value {
		indexes = append(indexes, index)
	}
	return indexes
}
