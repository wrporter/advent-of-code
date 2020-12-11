package arrays

func Min(array []int) int {
	var min = array[0]
	for _, value := range array {
		if min > value {
			min = value
		}
	}
	return min
}

func Max(array []int) int {
	var max = array[0]
	for _, value := range array {
		if value > max {
			max = value
		}
	}
	return max
}
