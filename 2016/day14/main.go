package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 14
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/sample-input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	salt := input[0]
	hashFunc := getHash
	return getIndexFor64thKey(hashFunc, salt)
}

// 35364 is too high
// 20219 is too high
// 20218 is too high
// 20217 not correct
func part2(input []string) interface{} {
	salt := input[0]
	hashFunc := getHash2017
	return getIndexFor64thKey(hashFunc, salt)
}

func getIndexFor64thKey(hash func(salt string, index int) string, salt string) interface{} {
	var history []potentialKey
	targetNumKeys := 64
	index := 0
	numKeys := 0
	counted := make(map[string]bool)

	for ; ; index++ {
		stream := hash(salt, index)
		quintKey := potentialKey{
			stream:       stream,
			firstTriplet: getFirstTriplet(stream),
			quints:       getQuints(stream),
		}
		history = append(history, quintKey)

		start := index - 1001
		if start < 0 {
			start = 0
		}
		for i := start; i < index; i++ {
			tripleKey := history[i]
			if tripleKey.firstTriplet != "" &&
				!counted[tripleKey.stream] &&
				quintKey.quints[strings.Repeat(string(tripleKey.firstTriplet[0]), 5)] {
				counted[tripleKey.stream] = true
				numKeys++
			}

			if numKeys == targetNumKeys {
				return i
			}
		}
	}
}

func getHash(salt string, index int) string {
	hash := md5.Sum([]byte(fmt.Sprintf("%s%d", salt, index)))
	stream := hex.EncodeToString(hash[:])
	return stream
}

func getHash2017(salt string, index int) string {
	hash := md5.Sum([]byte(fmt.Sprintf("%s%d", salt, index)))
	stream := hex.EncodeToString(hash[:])
	for i := 0; i < 2016; i++ {
		hash = md5.Sum([]byte(stream))
		stream = hex.EncodeToString(hash[:])
	}
	stream = hex.EncodeToString(hash[:])
	return stream
}

func getQuints(stream string) map[string]bool {
	quints := make(map[string]bool)
	count := 1
	var prev int32
	for _, char := range stream {
		if char == prev {
			count++
		} else {
			count = 1
		}
		if count == 5 {
			quints[strings.Repeat(string(char), 5)] = true
		}
		prev = char
	}
	return quints
}

func getFirstTriplet(stream string) string {
	count := 1
	var prev int32
	for _, char := range stream {
		if char == prev {
			count++
		} else {
			count = 1
		}
		if count == 3 {
			return strings.Repeat(string(char), 3)
		}
		prev = char
	}
	return ""
}

type potentialKey struct {
	stream       string
	firstTriplet string
	quints       map[string]bool
}
