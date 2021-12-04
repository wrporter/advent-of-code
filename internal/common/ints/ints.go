package ints

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

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

func Max(values ...int) int {
	max := 0
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
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
	if res < 0 {
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

func Product(array []int) int {
	product := array[0]
	for i := 1; i < len(array); i++ {
		product *= array[i]
	}
	return product
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

func ChineseRemainderTheorem(a, n []int) int {
	aBig := make([]*big.Int, len(a))
	nBig := make([]*big.Int, len(n))
	for i, value := range a {
		aBig[i] = big.NewInt(int64(value))
	}
	for i, value := range n {
		nBig[i] = big.NewInt(int64(value))
	}
	result, _ := ChineseRemainderTheoremBig(aBig, nBig)
	return int(result.Int64())
}

var one = big.NewInt(1)

func ChineseRemainderTheoremBig(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func HashCode(a []int) int {
	if len(a) == 0 {
		return 0
	}

	result := 1
	for _, element := range a {
		result = 31*result + element
	}

	return result
}

func Join(elems []int, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%d", elems[0])
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(fmt.Sprintf("%d", elems[i]))
	}

	var b strings.Builder
	b.Grow(n)
	_, _ = fmt.Fprintf(&b, "%d", elems[0])
	for _, elem := range elems[1:] {
		b.WriteString(sep)
		_, _ = fmt.Fprintf(&b, "%d", elem)
	}
	return b.String()
}

func Contains(values []int, value int) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func Reverse(values []int) []int {
	result := Copy(values)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func GetWindow(values [][]int, x, y, size int) [][]int {
	window := make([][]int, size)
	for row := 0; row < size; row++ {
		window[row] = make([]int, size)
		for col := 0; col < size; col++ {
			window[row][col] = values[row+y][col+x]
		}
	}
	return window
}

func SumGrid(values [][]int) int {
	sum := 0
	for _, row := range values {
		sum += Sum(row)
	}
	return sum
}
