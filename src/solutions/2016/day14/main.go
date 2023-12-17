package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2016, 14
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

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

func part2(input []string) interface{} {
	salt := input[0]
	hashFunc := getStretchedHash
	return getIndexFor64thKey(hashFunc, salt)
}

func getIndexFor64thKey(hash func(salt string, index int) string, salt string) interface{} {
	var history []potentialKey
	targetNumKeys := 64
	numKeys := 0
	var key potentialKey

	for index := 0; numKeys < targetNumKeys; index++ {
		stream := hash(salt, index)
		next := potentialKey{
			stream: stream,
			triple: getFirstTriple(stream),
			quints: getQuints(stream),
			index:  index,
		}
		history = append(history, next)

		if index < 1000 {
			continue
		}

		// look back and grab potential key
		key = history[index-1000]
		if key.triple == "" {
			continue
		}
		targetQuint := strings.Repeat(string(key.triple[0]), 5)

		// evaluate potential key against next 1000 hashes
		for i := index - 1000 + 1; i <= index; i++ {
			quint := history[i]
			if quint.quints[targetQuint] {
				numKeys++
				break
			}
		}
	}

	return key.index
}

func getHash(salt string, index int) string {
	hash := md5.Sum([]byte(fmt.Sprintf("%s%d", salt, index)))
	stream := hex.EncodeToString(hash[:])
	return stream
}

func getStretchedHash(salt string, index int) string {
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

func getFirstTriple(stream string) string {
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
	stream string
	triple string
	quints map[string]bool
	index  int
}
