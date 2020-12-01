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
		for current, value := range values {
			s[start] = value

			if start == last {
				emit(s)
			} else {
				rc(start+1, current+1)
			}
		}
	}

	rc(0, 0)
}

func ComboInts(values []int, size int, emit func([]int)) {
	s := make([]int, size)
	last := size - 1
	var rc func(int, int)

	rc = func(start, next int) {
		for current, value := range values {
			s[start] = value

			if start == last {
				emit(s)
			} else {
				rc(start+1, current+1)
			}
		}
	}

	rc(0, 0)
}

func PermuteSize(values []int, startSize int, endSize int, emit func([]int)) {
	var permuteSize func([]int, int, int)

	permuteSize = func(current []int, index int, size int) {
		if len(current) == size {
			emit(current)
			return
		}

		for i := index; i < len(values); i++ {
			current = append(current, values[i])
			permuteSize(current, i+1, size)
			current = current[:len(current)-1]
		}
	}

	for size := startSize; size <= endSize; size++ {
		permuteSize(nil, 0, size)
	}
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

func PermuteStrings(values []string, output func([]string)) {
	permuteStrings(values, output, 0)
}

func permuteStrings(values []string, output func([]string), i int) {
	if i > len(values) {
		output(values)
		return
	}

	permuteStrings(values, output, i+1)

	for j := i + 1; j < len(values); j++ {
		values[i], values[j] = values[j], values[i]
		permuteStrings(values, output, i+1)
		values[i], values[j] = values[j], values[i]
	}
}
