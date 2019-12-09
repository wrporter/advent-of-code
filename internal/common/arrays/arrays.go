package arrays

func CopyInts(array []int) []int {
	cpy := make([]int, len(array))
	copy(cpy, array)
	return cpy
}

func Poll(array []int) (int, []int) {
	return array[0], array[1:]
}

func Pop(array []int) (int, []int) {
	size := len(array)
	return array[size-1], array[:size-1]
}
