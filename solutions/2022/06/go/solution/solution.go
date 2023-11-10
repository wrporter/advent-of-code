package solution

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	return getMarker(input, 4)
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	return getMarker(input, 14)
}

func getMarker(buffer string, size int) interface{} {
	for i := size; i < len(buffer); i++ {
		marker := buffer[i-size : i]

		unique := make(map[rune]bool)
		for _, c := range marker {
			unique[c] = true
		}

		if len(unique) == size {
			return i
		}
	}

	return "TBD"
}
