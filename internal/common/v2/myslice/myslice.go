package myslice

import "golang.org/x/exp/constraints"

// Rotate90Copy rotates the grid by 90 degrees. For example, the following grid
//
// ```
// 1 2 3
// 4 5 6
// 7 8 9
// ```
//
// Would result in the following:
//
// ```
// 7 4 1
// 8 5 2
// 9 6 3
// ```
func Rotate90Copy[T any](grid [][]T) [][]T {
	next := Copy2D(grid)
	Rotate90(next)
	return next
}

// Rotate90 does the same thing as Rotate90Copy, but rotates the grid in place rather than copying
// it to a new one.
func Rotate90[T any](grid [][]T) {
	size := len(grid)

	for i := 0; i < size/2; i++ {
		for j := i; j < size-i-1; j++ {
			Swap(&grid[i][j], &grid[size-j-1][i])
			Swap(&grid[size-j-1][i], &grid[size-i-1][size-j-1])
			Swap(&grid[size-i-1][size-j-1], &grid[j][size-i-1])
		}
	}
}

func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func Copy2D[T any](grid [][]T) [][]T {
	cpy := make([][]T, len(grid))
	for i := range grid {
		cpy[i] = Copy(grid[i])
	}
	return cpy
}

func Copy[T any](array []T) []T {
	cpy := make([]T, len(array))
	copy(cpy, array)
	return cpy
}

func GetCol[T any](grid [][]T, col int) []T {
	column := make([]T, len(grid))
	for row := 0; row < len(grid); row++ {
		column[row] = grid[row][col]
	}
	return column
}

func Min[T constraints.Ordered](array []T) T {
	result := array[0]
	for _, value := range array {
		result = min(result, value)
	}
	return result
}

func Max[T constraints.Ordered](array []T) T {
	result := array[0]
	for _, value := range array {
		result = max(result, value)
	}
	return result
}

// Remove removes the element at the given index and returns a new slice.
func Remove[T any](s []T, index int) []T {
	result := make([]T, 0)
	result = append(result, s[:index]...)
	return append(result, s[index+1:]...)
}
