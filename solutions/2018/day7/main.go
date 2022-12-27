package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/runes"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"sort"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2018, 7
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^Step ([A-Z]) must be finished before step ([A-Z]) can begin\.$`)

func part1(input []string) interface{} {
	dependencies := parseInput(input)
	remaining := make(map[rune]bool)
	for step := range dependencies {
		remaining[step] = true
	}

	done := make(map[rune]bool)
	order := make([]rune, len(remaining))
	for len(remaining) > 0 {
		next := getNextSteps(remaining, dependencies, done)[0]
		order = append(order, next)
		done[next] = true
		delete(remaining, next)
	}

	return string(order)
}

func part2(input []string) interface{} {
	numWorkers := 5
	baseTime := 60
	dependencies := parseInput(input)
	remaining := make(map[rune]bool)
	for step := range dependencies {
		remaining[step] = true
	}

	done := make(map[rune]bool)
	workers := newWorkers(numWorkers)
	totalTime := -1
	for len(remaining) > 0 || isWorkInProgress(workers) {
		// make sure we always look at working workers first, so we don't skip over others when a step is completed
		sort.SliceStable(workers, func(i, j int) bool {
			return workers[i].step != 0
		})
		for _, w := range workers {
			w.time--

			if w.time <= 0 {
				if w.step != 0 {
					//fmt.Printf("DONE >> %v [%d]\n", w, totalTime)
					done[w.step] = true
					w.step = 0
				}

				nextSteps := getNextSteps(remaining, dependencies, done)
				if len(nextSteps) > 0 {
					step := nextSteps[0]
					w.time = timeToComplete(baseTime, step)
					w.step = step
					delete(remaining, step)
					//fmt.Printf("START >> %v [%d]\n", w, totalTime)
				}
			}
		}
		totalTime++
	}

	return totalTime
}

func isWorkInProgress(workers []*worker) bool {
	for _, w := range workers {
		if w.time > 0 {
			return true
		}
	}
	return false
}

func timeToComplete(baseTime int, step rune) int {
	return baseTime + int(step-'A') + 1
}

func newWorkers(numWorkers int) []*worker {
	var workers []*worker
	for i := 0; i < numWorkers; i++ {
		workers = append(workers, &worker{id: i})
	}
	return workers
}

type worker struct {
	id   int
	step rune
	time int
}

func (w *worker) String() string {
	return fmt.Sprintf("(%d) %c - %d", w.id, w.step, w.time)
}

func getNextSteps(remaining map[rune]bool, dependencies map[rune]map[rune]bool, done map[rune]bool) []rune {
	var nextSteps []rune
	for step := range remaining {
		if available(dependencies, done, step) {
			nextSteps = append(nextSteps, step)
		}
	}
	runes.Sort(nextSteps)
	return nextSteps
}

func available(dependencies map[rune]map[rune]bool, done map[rune]bool, step rune) bool {
	for dependency := range dependencies[step] {
		if !done[dependency] {
			return false
		}
	}
	return true
}

func parseInput(input []string) map[rune]map[rune]bool {
	dependencies := make(map[rune]map[rune]bool)
	for _, line := range input {
		match := regex.FindStringSubmatch(line)
		dependency := rune(match[1][0])
		step := rune(match[2][0])

		if dependencies[step] == nil {
			dependencies[step] = map[rune]bool{dependency: true}
		} else {
			dependencies[step][dependency] = true
		}

		if dependencies[dependency] == nil {
			dependencies[dependency] = make(map[rune]bool)
		}
	}
	return dependencies
}
