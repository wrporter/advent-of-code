package arrays

import "golang.org/x/exp/constraints"

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
