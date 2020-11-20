package ints

import "math"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Pow(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Sign(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func Sqrt(num int) int {
	return int(math.Sqrt(float64(num)))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func WrapMod(value, modulus int) int {
	var res = value % modulus
	if (res < 0 && modulus > 0) || (res > 0 && modulus < 0) {
		return res + modulus
	}
	return res
}

func Copy(array []int) []int {
	cpy := make([]int, len(array))
	copy(cpy, array)
	return cpy
}

func Copy2D(grid [][]int) [][]int {
	cpy := make([][]int, len(grid))
	for i := range grid {
		cpy[i] = Copy(grid[i])
	}
	return cpy
}

func Prepend(array []int, value int) []int {
	array = append(array, 0)
	copy(array[1:], array)
	array[0] = value
	return array
}

func Poll(array []int) (int, []int) {
	return array[0], array[1:]
}

func Pop(array []int) (int, []int) {
	size := len(array)
	return array[size-1], array[:size-1]
}

func Sum(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func GetDivisors(num int) []int {
	var divisors []int
	var lastOnes []int

	end := Sqrt(num)
	for candidate := 1; candidate <= end; candidate++ {
		if num%candidate == 0 {
			divisors = append(divisors, candidate)
			if num/candidate != candidate {
				lastOnes = Prepend(lastOnes, num/candidate)
			}
		}
	}

	divisors = append(divisors, lastOnes...)
	return divisors
}

func TakeLast(values []int, count int) []int {
	start := len(values) - count
	if start < 0 {
		start = 0
	}
	return values[start:]
}
