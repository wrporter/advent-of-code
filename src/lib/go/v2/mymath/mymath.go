package mymath

import (
	"aoc/src/lib/go/v2/contain"
	"fmt"
	"math"
	"math/big"
)

type (
	Int interface {
		int | int8 | int16 | int32 | int64
	}
	UInt interface {
		uint | uint8 | uint16 | uint32 | uint64
	}
	Float interface {
		float32 | float64
	}
	Number interface {
		Int | UInt | Float
	}
)

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func MinOf[T Number]() (r T) {
	switch x := any(&r).(type) {
	case *int:
		*x = math.MinInt
	case *int8:
		*x = math.MinInt8
	case *int16:
		*x = math.MinInt16
	case *int32:
		*x = math.MinInt32
	case *int64:
		*x = math.MinInt64
	case *uint:
		*x = 0
	case *uint8:
		*x = 0
	case *uint16:
		*x = 0
	case *uint32:
		*x = 0
	case *uint64:
		*x = 0
	case *float32:
		*x = -math.MaxFloat32
	case *float64:
		*x = -math.MaxFloat64
	default:
		panic("unreachable")
	}
	return
}

func MaxOf[T Number]() (r T) {
	switch x := any(&r).(type) {
	case *int:
		*x = math.MaxInt
	case *int8:
		*x = math.MaxInt8
	case *int16:
		*x = math.MaxInt16
	case *int32:
		*x = math.MaxInt32
	case *int64:
		*x = math.MaxInt64
	case *uint:
		*x = math.MaxUint
	case *uint8:
		*x = math.MaxUint8
	case *uint16:
		*x = math.MaxUint16
	case *uint32:
		*x = math.MaxUint32
	case *uint64:
		*x = math.MaxUint64
	case *float32:
		*x = math.MaxFloat32
	case *float64:
		*x = math.MaxFloat64
	default:
		panic("unreachable")
	}
	return
}

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

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Max[T Number](values ...T) T {
	max := MinOf[T]()
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func Min[T Number](values ...T) T {
	min := MaxOf[T]()
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func Round[T Float](num T) int {
	return int(float64(num) + math.Copysign(0.5, float64(num)))
}

func Sqrt(num int) int {
	return int(math.Sqrt(float64(num)))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}

// GCD returns the greatest common divisor (GCD) via the Euclidean algorithm.
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM returns the Least Common Multiple (LCM) via GCD.
func LCM(values ...int) int {
	if len(values) == 0 {
		return 0
	}
	if len(values) == 1 {
		return values[0]
	}

	a, b := values[0], values[1]
	result := a * b / GCD(a, b)

	rest := values[2:]
	for i := 0; i < len(rest); i++ {
		result = LCM(result, rest[i])
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

func Sum[T Number](array []T) T {
	sum := T(0)
	for _, value := range array {
		sum += value
	}
	return sum
}

func SumGrid[T Number](values [][]T) T {
	sum := T(0)
	for _, row := range values {
		sum += Sum(row)
	}
	return sum
}

func Product[T Number](array []T) T {
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
				lastOnes = contain.Prepend(lastOnes, num/candidate)
			}
		}
	}

	divisors = append(divisors, lastOnes...)
	return divisors
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

func HashCode[T Number](a []T) T {
	if len(a) == 0 {
		return 0
	}

	result := T(1)
	for _, element := range a {
		result = T(31)*result + element
	}

	return result
}
