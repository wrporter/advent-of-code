package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"regexp"
)

var regex = regexp.MustCompile(`([a-zA-Z]+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)

type Reindeer struct {
	name  string
	speed int
	time  int
	rest  int
}

type DeerState struct {
	name     string
	distance int
	flyTime  int
	restTime int
}

func getWinner(reindeerSpeeds []string, time int) DeerState {
	race, speeds := parse(reindeerSpeeds)
	leader := &DeerState{distance: 0}

	for timeStep := 0; timeStep < time; timeStep++ {
		for name, state := range race {
			if state.restTime > 0 {
				state.restTime--
				if state.restTime == 0 {
					state.flyTime = speeds[name].time
				}
			} else {
				state.distance += speeds[name].speed
				state.flyTime--
				if state.flyTime == 0 {
					state.restTime = speeds[name].rest
				}
			}
			if state.distance > leader.distance {
				leader = state
			}
		}
	}

	return *leader
}

func parse(reindeerSpeeds []string) (map[string]*DeerState, map[string]Reindeer) {
	race := make(map[string]*DeerState)
	speeds := make(map[string]Reindeer)

	for _, reindeerSpeed := range reindeerSpeeds {
		match := regex.FindStringSubmatch(reindeerSpeed)
		deer := Reindeer{
			name:  match[1],
			speed: conversion.StringToInt(match[2]),
			time:  conversion.StringToInt(match[3]),
			rest:  conversion.StringToInt(match[4]),
		}
		race[deer.name] = &DeerState{
			name:     deer.name,
			distance: 0,
			flyTime:  deer.time,
			restTime: 0,
		}
		speeds[deer.name] = deer
	}
	return race, speeds
}

func main() {
	lines, _ := file.ReadFile("./2015/day14/input.txt")
	//lines := []string{
	//	"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
	//	"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
	//}
	fmt.Println(getWinner(lines, 2503))
}
