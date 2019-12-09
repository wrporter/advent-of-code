package probability

func Permute(values []int, output func([]int)) {
	permute(values, output, 0)
}

func permute(values []int, output func([]int), i int) {
	if i > len(values) {
		output(values)
		return
	}

	permute(values, output, i+1)

	for j := i + 1; j < len(values); j++ {
		values[i], values[j] = values[j], values[i]
		permute(values, output, i+1)
		values[i], values[j] = values[j], values[i]
	}
}
