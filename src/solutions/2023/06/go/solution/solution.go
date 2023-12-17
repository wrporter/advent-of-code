package solution

import (
	"aoc/src/lib/go/convert"
	"math"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	parts := strings.Split(input, "\n")
	times, _ := convert.ToInts(strings.Fields(parts[0])[1:])
	records, _ := convert.ToInts(strings.Fields(parts[1])[1:])
	product := 1

	for i := range times {
		product *= countWaysToWin_Slowest(times[i], records[i])
	}

	return product
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	parts := strings.Split(input, "\n")
	time := convert.StringToInt(strings.Join(strings.Fields(parts[0])[1:], ""))
	record := convert.StringToInt(strings.Join(strings.Fields(parts[1])[1:], ""))

	return countWaysToWin_Optimal(time, record)
}

func countWaysToWin_Slowest(raceDuration int, record int) int {
	waysToWin := 0

	for timeToPushButton := 1; timeToPushButton < raceDuration; timeToPushButton++ {
		remainingTime := raceDuration - timeToPushButton
		distance := timeToPushButton * remainingTime
		if distance > record {
			waysToWin++
		}
	}

	return waysToWin
}

func countWaysToWin_Faster(raceDuration int, record int) int {
	start := 0
	end := 0

	for timeToPushButton := 1; timeToPushButton < raceDuration; timeToPushButton++ {
		remainingTime := raceDuration - timeToPushButton
		distance := timeToPushButton * remainingTime
		if distance > record {
			start = timeToPushButton
			break
		}
	}

	for timeToPushButton := raceDuration - 1; timeToPushButton > 0; timeToPushButton-- {
		remainingTime := raceDuration - timeToPushButton
		distance := timeToPushButton * remainingTime
		if distance > record {
			end = timeToPushButton
			break
		}
	}

	return end - start + 1
}

func countWaysToWin_Optimal(time int, record int) int {
	// a = -1 due to the step time being constant and a < 0 indicates a
	// parabola open downward (upside down U)
	a := float64(-1)
	// b is the duration of the race and we shift the parabola to the right of
	// the y-axis
	b := float64(time)
	// c shifts the roots of the parabola to the edges of the record we need to
	// beat
	c := float64(-(record + 1))

	// Use the quadratic formula to find the x roots
	x1 := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	x2 := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)

	return int(math.Floor(x2) - math.Ceil(x1) + 1)
}
