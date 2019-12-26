package probability

func ComboSpots(values []string, startSize, endSize int, emit func([]string)) {
	for i := startSize; i <= endSize; i++ {
		Combo(values, i, emit)
	}
}

func Combo(values []string, size int, emit func([]string)) {
	s := make([]string, size)
	last := size - 1
	var rc func(int, int)

	rc = func(start, next int) {
		for end, value := range values {
			s[start] = value

			if start == last {
				emit(s)
			} else {
				rc(start+1, end+1)
			}
		}
		return
	}

	rc(0, 0)
}

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
