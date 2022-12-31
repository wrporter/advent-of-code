package contain

func Poll[T any](array []T) (T, []T) {
	return array[0], array[1:]
}

func Pop[T any](array []T) (T, []T) {
	size := len(array)
	return array[size-1], array[:size-1]
}
