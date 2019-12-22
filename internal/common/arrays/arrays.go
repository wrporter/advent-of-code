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
