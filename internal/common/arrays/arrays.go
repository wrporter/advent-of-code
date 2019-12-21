package arrays

func Poll(array []int) (int, []int) {
	return array[0], array[1:]
}

func Pop(array []int) (int, []int) {
	size := len(array)
	return array[size-1], array[:size-1]
}

func Min(array []int) int {
	var min = array[0]
	for _, value := range array {
		if min > value {
			min = value
		}
	}
	return min
}
