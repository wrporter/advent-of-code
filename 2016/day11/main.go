package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/mystrings"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 11
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`a ([a-z]+)(?:-compatible)? (generator|microchip)`)

func part1(input []string) interface{} {
	floors := parseFloors(input)
	return getMinNumMoves(floors)
}

func part2(input []string) interface{} {
	floors := parseFloors(input)
	floors[0] = append(floors[0], Item{Element: "elerium", Type: Generator})
	floors[0] = append(floors[0], Item{Element: "elerium", Type: Microchip})
	floors[0] = append(floors[0], Item{Element: "dilithium", Type: Generator})
	floors[0] = append(floors[0], Item{Element: "dilithium", Type: Microchip})
	return getMinNumMoves(floors)
}

func getMinNumMoves(floors [][]Item) int {
	state := State{
		NumMoves: 0,
		Floors:   floors,
		Elevator: 0,
	}
	queue := []State{state}
	seen := map[string]bool{state.Hash(): true}

	for len(queue) > 0 {
		state, queue = queue[0], queue[1:]

		if state.allItemsAreOnTheFourthFloor() {
			return state.NumMoves
		}

		for _, nextState := range state.NextStates() {
			if !seen[nextState.Hash()] {
				queue = append(queue, nextState)
				seen[nextState.Hash()] = true
			}
		}
	}

	return -1
}

func (s State) allItemsAreOnTheFourthFloor() bool {
	numItems := 0
	for _, floor := range s.Floors {
		numItems += len(floor)
	}
	return len(s.Floors[3]) == numItems
}

var (
	Generator ItemType = "generator"
	Microchip ItemType = "microchip"
)

type (
	State struct {
		NumMoves int
		Floors   [][]Item
		Elevator int
		Parent   *State
	}
	ItemType string
	Item     struct {
		Element string
		Type    ItemType
	}
)

func (s State) String() string {
	floors := make([]string, len(s.Floors))

	for i, floor := range s.Floors {
		level := "-"
		if s.Elevator == i {
			level = "E"
		}

		items := make([]string, len(floor))
		sort.SliceStable(floor, func(i, j int) bool {
			return strings.Compare(floor[i].Element, floor[j].Element) == -1
		})

		for j, item := range floor {
			items[j] = item.String()
		}

		floors[i] = fmt.Sprintf("F%d: %s %s", i+1, level, strings.Join(items, " "))
	}

	return strings.Join(mystrings.ReverseList(floors), "\n")
}

func (i Item) String() string {
	elementLetter := unicode.ToUpper(rune(i.Element[0]))
	generatorLetter := unicode.ToUpper(rune(i.Type[0]))
	return string(elementLetter) + string(generatorLetter)
}

// IsValid returns whether the state is valid. A state is invalid if a
// microchip is on a floor with a generator that is not its own.
func (s State) IsValid() bool {
	for _, floor := range s.Floors {
		for _, chip := range floor {
			if chip.Type != Microchip {
				continue
			}

			ownGeneratorExists := false
			otherGeneratorExists := false
			for _, generator := range floor {
				if generator.Type != Generator {
					continue
				}

				if chip.Element == generator.Element {
					ownGeneratorExists = true
				}
				if chip.Element != generator.Element {
					otherGeneratorExists = true
				}
			}
			if otherGeneratorExists && !ownGeneratorExists {
				return false
			}
		}
	}

	return true
}

func (s State) NextStates() []State {
	var nextStates []State

	var combinations [][]Item
	Combinations(s.Floors[s.Elevator], 1, 2, func(combo []Item) {
		combinations = append(combinations, deepCopyFloor(combo))
	})

	deltas := []int{-1, 1}
	for _, delta := range deltas {
		if s.Elevator+delta < 0 || s.Elevator+delta >= len(s.Floors) {
			continue
		}

		for _, itemsToMove := range combinations {
			newFloors := s.deepCopyFloors()
			for _, item := range itemsToMove {
				newFloors[s.Elevator] = remove(newFloors[s.Elevator], item)
				newFloors[s.Elevator+delta] = append(newFloors[s.Elevator+delta], item)
			}
			next := State{
				NumMoves: s.NumMoves + 1,
				Floors:   newFloors,
				Elevator: s.Elevator + delta,
				Parent:   &s,
			}
			if next.IsValid() {
				nextStates = append(nextStates, next)
			}
		}
	}

	return nextStates
}

func (s State) PrintAll() {
	for node := &s; node != nil; node = node.Parent {
		fmt.Println(node, "\n")
	}
}

func remove(items []Item, item Item) []Item {
	index := -1
	for i, suspect := range items {
		if suspect == item {
			index = i
			break
		}
	}
	return append(items[:index], items[index+1:]...)
}

func (s State) deepCopyFloors() [][]Item {
	floors := make([][]Item, len(s.Floors))
	for i, floor := range s.Floors {
		floors[i] = deepCopyFloor(floor)
	}
	return floors
}

func deepCopyFloor(floor []Item) []Item {
	newFloor := make([]Item, len(floor))
	for j, item := range floor {
		newFloor[j] = Item{Element: item.Element, Type: item.Type}
	}
	return newFloor
}

func (s State) Hash() string {
	floors := make([]string, len(s.Floors))

	for i, floor := range s.Floors {
		items := make([]string, len(floor))
		sort.SliceStable(floor, func(i, j int) bool {
			return strings.Compare(floor[i].Element, floor[j].Element) == -1
		})

		for j, item := range floor {
			items[j] = item.String()
		}

		floors[i] = fmt.Sprintf("%d: %s", i+1, strings.Join(items, " "))
	}

	return fmt.Sprintf("State<%d, %s>", s.Elevator, strings.Join(mystrings.ReverseList(floors), ","))
}

func parseFloors(input []string) [][]Item {
	floors := make([][]Item, len(input))
	for floor, line := range input {
		if !strings.Contains(line, "nothing relevant") {
			matches := regex.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				item := Item{
					Element: match[1],
					Type:    ItemType(match[2]),
				}
				floors[floor] = append(floors[floor], item)
			}
		}
	}
	return floors
}

func Combinations(values []Item, startSize int, endSize int, emit func([]Item)) {
	var comboSize func([]Item, int, int)

	comboSize = func(current []Item, index int, size int) {
		if len(current) == size {
			emit(current)
			return
		}

		for i := index; i < len(values); i++ {
			current = append(current, values[i])
			comboSize(current, i+1, size)
			current = current[:len(current)-1]
		}
	}

	for size := startSize; size <= endSize; size++ {
		comboSize(nil, 0, size)
	}
}
