package contain

import (
	"fmt"
	"strings"
)

func Copy[T any](array []T) []T {
	cpy := make([]T, len(array))
	copy(cpy, array)
	return cpy
}

func Copy2D[T any](grid [][]T) [][]T {
	cpy := make([][]T, len(grid))
	for i := range grid {
		cpy[i] = Copy(grid[i])
	}
	return cpy
}

func Prepend[T any](array []T, value T) []T {
	array = append(array, GetZero[T]())
	copy(array[1:], array)
	array[0] = value
	return array
}

func Poll[T any](array []T) (T, []T) {
	return array[0], array[1:]
}

func Pop[T any](array []T) (T, []T) {
	size := len(array)
	return array[size-1], array[:size-1]
}

func TakeLast[T any](values []T, count int) []T {
	start := len(values) - count
	if start < 0 {
		start = 0
	}
	return values[start:]
}

func Contains[T comparable](values []T, value T) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func Reverse[T any](values []T) []T {
	result := Copy(values)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func GetWindow[T any](values [][]T, x, y, size int) [][]T {
	window := make([][]T, size)
	for row := 0; row < size; row++ {
		window[row] = make([]T, size)
		for col := 0; col < size; col++ {
			window[row][col] = values[row+y][col+x]
		}
	}
	return window
}

func Join[T any](elems []T, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%v", elems[0])
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(fmt.Sprintf("%v", elems[i]))
	}

	var b strings.Builder
	b.Grow(n)
	_, _ = fmt.Fprintf(&b, "%v", elems[0])
	for _, elem := range elems[1:] {
		b.WriteString(sep)
		_, _ = fmt.Fprintf(&b, "%v", elem)
	}
	return b.String()
}
