package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/v2/myslice"
	"slices"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	sequence := strings.Split(input, ",")
	sum := 0
	for _, step := range sequence {
		sum += hash(step)
	}
	return sum
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	sequence := strings.Split(input, ",")
	boxes := make(map[int][]Lens)

	for _, step := range sequence {
		if i := strings.IndexRune(step, '='); i >= 0 {
			lens := Lens{
				label:       step[:i],
				focalLength: convert.StringToInt(step[i+1:]),
			}
			box := hash(lens.label)

			slot := slices.IndexFunc(boxes[box], lens.equals)
			if slot >= 0 {
				boxes[box][slot] = lens
			} else {
				boxes[box] = append(boxes[box], lens)
			}
		} else {
			lens := Lens{label: step[:len(step)-1]}
			box := hash(lens.label)

			slot := slices.IndexFunc(boxes[box], lens.equals)
			if slot >= 0 {
				boxes[box] = myslice.Remove(boxes[box], slot)
			}
		}
	}

	total := 0
	for box, lenses := range boxes {
		for slot, lens := range lenses {
			total += (box + 1) * (slot + 1) * lens.focalLength
		}
	}
	return total
}

type Lens struct {
	label       string
	focalLength int
}

func (l Lens) equals(l2 Lens) bool {
	return l.label == l2.label
}

func hash(str string) int {
	var value int
	for _, c := range str {
		value += int(c)
		value *= 17
		value %= 256
	}
	return value
}
