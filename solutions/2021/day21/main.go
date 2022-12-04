package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 21
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^Player (\d+) starting position: (\d+)$`)

func part1(input []string) interface{} {
	players := parseInput(input)

	d := &dice{}
	var loser *player

	for turn := 0; ; turn = (turn + 1) % len(players) {
		value := d.roll() + d.roll() + d.roll()
		p := players[turn]
		p.position = modOffset(p.position, value, 10, 1)
		p.score += p.position

		if p.score >= 1000 {
			loser = players[(turn+1)%len(players)]
			break
		}
	}

	return loser.score * d.value
}

func part2(input []string) interface{} {
	players := parseInput(input)
	dp := make(map[state]winCounter)
	wins := countWins(dp, state{
		p1: players[0].position,
		s1: 0,
		p2: players[1].position,
		s2: 0,
	})

	if wins.p1 > wins.p2 {
		return wins.p1
	}
	return wins.p2
}

func countWins(dp map[state]winCounter, s state) winCounter {
	if s.s1 >= 21 {
		return winCounter{p1: 1, p2: 0}
	} else if s.s2 >= 21 {
		return winCounter{p1: 0, p2: 1}
	} else if wins, ok := dp[s]; ok {
		return wins
	}

	wins := winCounter{p1: 0, p2: 0}
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {
				position := modOffset(s.p1, d1+d2+d3, 10, 1)
				score := s.s1 + position
				win := countWins(dp, state{
					p1: s.p2,
					s1: s.s2,
					p2: position,
					s2: score,
				})
				wins = winCounter{p1: wins.p1 + win.p2, p2: wins.p2 + win.p1}
			}
		}
	}

	dp[s] = wins

	return wins
}

func parseInput(input []string) []*player {
	players := make([]*player, len(input))
	for i, line := range input {
		match := regex.FindStringSubmatch(line)
		players[i] = &player{
			score:    0,
			position: convert.StringToInt(match[2]),
		}
	}
	return players
}

func modOffset(value, plus, modulo, offset int) int {
	return ((value + plus - offset) % modulo) + offset
}

type winCounter struct {
	p1 int64
	p2 int64
}

type state struct {
	p1 int
	s1 int
	p2 int
	s2 int
}

type player struct {
	score    int
	position int
}

type dice struct {
	value int
}

func (d *dice) roll() int {
	d.value++
	return d.value
}
