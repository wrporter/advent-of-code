package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"regexp"
	"sort"
	"strings"
	"time"
)

func main() {
	year, day := 2018, 4
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^\[(.+)] (falls asleep|wakes up|Guard #(\d+) begins shift)$`)

const timeLayout = "2006-01-02 15:04"

type (
	Action string
	Log    struct {
		Time    time.Time
		GuardID int
		Action  Action
	}
)

const (
	BeginsShift = "begins shift"
	FallsAsleep = "falls asleep"
	WakesUp     = "wakes up"
)

func part1(input []string) interface{} {
	logs := parselogs(input)
	guardMinutesAsleep, guardAsleepTime := getGuardAsleepTime(logs)

	var guardThatSleptTheMost int
	maxSleepTime := 0
	for guardID, sleepTime := range guardAsleepTime {
		if sleepTime > maxSleepTime {
			guardThatSleptTheMost = guardID
			maxSleepTime = sleepTime
		}
	}

	minuteAsleepTheMost := -1
	maxSleepMoments := 0
	for minute, momentsAsleep := range guardMinutesAsleep[guardThatSleptTheMost] {
		if momentsAsleep > maxSleepMoments {
			minuteAsleepTheMost = minute
			maxSleepMoments = momentsAsleep
		}
	}

	return guardThatSleptTheMost * minuteAsleepTheMost
}

func part2(input []string) interface{} {
	logs := parselogs(input)
	guardMinutesAsleep, _ := getGuardAsleepTime(logs)

	var guardAsleepTheMostOnSameMinute int
	var minuteMostFrequentlySlept int
	maxSleepTimeOnSameMinute := 0
	for guardID, minutesAsleep := range guardMinutesAsleep {
		for minute, momentsAsleep := range minutesAsleep {
			if momentsAsleep > maxSleepTimeOnSameMinute {
				guardAsleepTheMostOnSameMinute = guardID
				minuteMostFrequentlySlept = minute
				maxSleepTimeOnSameMinute = momentsAsleep
			}
		}
	}

	return guardAsleepTheMostOnSameMinute * minuteMostFrequentlySlept
}

func getGuardAsleepTime(logs []Log) (map[int]map[int]int, map[int]int) {
	guardMinutesAsleep := make(map[int]map[int]int)
	guardAsleepTime := make(map[int]int)
	var guardID int
	var asleep time.Time
	for _, log := range logs {
		switch log.Action {
		case BeginsShift:
			guardID = log.GuardID
			if _, ok := guardMinutesAsleep[guardID]; !ok {
				guardMinutesAsleep[guardID] = make(map[int]int)
			}
			if _, ok := guardAsleepTime[guardID]; !ok {
				guardAsleepTime[guardID] = 0
			}
		case FallsAsleep:
			asleep = log.Time
		case WakesUp:
			for minute := asleep.Minute(); minute < log.Time.Minute(); minute++ {
				guardMinutesAsleep[guardID][minute]++
			}
			guardAsleepTime[guardID] += int(log.Time.Sub(asleep).Minutes())
		}
	}
	return guardMinutesAsleep, guardAsleepTime
}

func parselogs(input []string) []Log {
	logs := make([]Log, len(input))
	for i, line := range input {
		match := regex.FindStringSubmatch(line)
		t, _ := time.Parse(timeLayout, match[1])
		action := match[2]
		if strings.Contains(action, "begins shift") {
			action = BeginsShift
		}
		guardID := 0
		if len(match) > 3 {
			guardID = conversion.StringToInt(match[3])
		}

		logs[i] = Log{
			Time:    t,
			GuardID: guardID,
			Action:  Action(action),
		}
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].Time.Before(logs[j].Time)
	})
	return logs
}
