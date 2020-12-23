package main

import (
	"container/list"
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 23
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	cups := parse(input)
	cups = play(cups, 100)
	return labelsAfter1(cups)
}

func labelsAfter1(cups *list.List) string {
	var sb strings.Builder
	initial := get(cups, 1)

	for e := initial.Next(); e != initial; e = e.Next() {
		if e == nil {
			e = cups.Front()
		}
		_, _ = fmt.Fprintf(&sb, "%d", e.Value)
	}

	return sb.String()
}

func play(cups *list.List, moves int) *list.List {
	current := cups.Front()

	for move := 1; move <= moves; move++ {
		//fmt.Printf("-- move %d --\n", move)
		//fmt.Printf("cups: %s\n", render(cups, current))

		pickup := list.New()
		for pick := 0; pick < 3; pick++ {
			remove := current.Next()
			if remove == nil {
				remove = cups.Front()
			}
			pickup.PushBack(cups.Remove(remove))
		}
		//fmt.Printf("pickup: %s\n", renderList(pickup))

		destination := current
		destinationCup := getDestinationCup(cups, pickup, current)
		for destination == nil || destination.Value != destinationCup {
			if destination == nil {
				destination = cups.Front()
			} else {
				destination = destination.Next()
			}
		}
		//fmt.Printf("destination: %d\n\n", destination.Value)

		for cup := pickup.Back(); cup != nil; cup = cup.Prev() {
			cups.InsertAfter(cup.Value, destination)
		}

		current = current.Next()
		if current == nil {
			current = cups.Front()
		}
	}

	//fmt.Printf("-- final --\n")
	//fmt.Printf("cups: %s\n", render(cups, current))

	return cups
}

func getDestinationCup(cups *list.List, pickup *list.List, current *list.Element) int {
	high := highest(cups)
	low := lowest(cups)
	destination := current.Value.(int) - 1

	for contains(pickup, destination) || !contains(cups, destination) {
		destination--
		if destination < low {
			return high
		}
	}

	return destination
}

func contains(l *list.List, value int) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return true
		}
	}
	return false
}

func get(l *list.List, value int) *list.Element {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return e
		}
	}
	return nil
}

func highest(cups *list.List) int {
	max := 0
	for e := cups.Front(); e != nil; e = e.Next() {
		max = ints.Max(max, e.Value.(int))
	}
	return max
}

func lowest(cups *list.List) int {
	min := ints.MaxInt
	for e := cups.Front(); e != nil; e = e.Next() {
		min = ints.Min(min, e.Value.(int))
	}
	return min
}

func renderList(l *list.List) interface{} {
	var sb strings.Builder
	delimiter := ' '

	for e := l.Front(); e != nil; e = e.Next() {
		sb.WriteString(fmt.Sprintf("%d", e.Value))
		if e != l.Back() {
			sb.WriteRune(delimiter)
		}
	}

	return sb.String()
}

func render(cups *list.List, current *list.Element) string {
	var sb strings.Builder
	delimiter := ' '

	for cup := cups.Front(); cup != nil; cup = cup.Next() {
		if cup == current {
			sb.WriteString(fmt.Sprintf("(%d)", cup.Value))
		} else {
			sb.WriteString(fmt.Sprintf("%d", cup.Value))
		}
		if cup != cups.Back() {
			sb.WriteRune(delimiter)
		}
	}

	return sb.String()
}

func part2(input []string) interface{} {
	return 0
}

func parse(input []string) *list.List {
	labels := strings.Split(input[0], "")
	values, _ := conversion.ToInts(labels)
	cups := list.New()
	for _, value := range values {
		cups.PushBack(value)
	}
	return cups
}
