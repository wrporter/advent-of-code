package solution

import (
	"encoding/json"
	"sort"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	pairs := parseInput(input)
	count := 0

	for i, pair := range pairs {
		if compare(pair[0], pair[1]) < 0 {
			count += i + 1
		}
	}

	return count
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	pairs := parseInput(input)
	packets := make([][]interface{}, (len(pairs)*2)+2)
	dividerPackets := [][]interface{}{{[]interface{}{float64(2)}}, {[]interface{}{float64(6)}}}
	packets[0] = dividerPackets[0]
	packets[1] = dividerPackets[1]

	for i := 0; i < len(pairs); i++ {
		packets[(i*2)+2] = pairs[i][0]
		packets[(i*2)+3] = pairs[i][1]
	}

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) < 0
	})

	decoderKey := 1
	for i, packet := range packets {
		if compare(packet, dividerPackets[0]) == 0 || compare(packet, dividerPackets[1]) == 0 {
			decoderKey *= i + 1
		}
	}

	return decoderKey
}

func compare(a, b interface{}) int {
	aint, aintok := a.(float64)
	bint, bintok := b.(float64)

	if aintok && bintok {
		return int(aint - bint)
	} else if aintok {
		return compare([]interface{}{a}, b)
	} else if bintok {
		return compare(a, []interface{}{b})
	}

	alist := a.([]interface{})
	blist := b.([]interface{})

	for i := 0; i < len(alist) && i < len(blist); i++ {
		result := compare(alist[i], blist[i])
		if result != 0 {
			return result
		}
	}

	return len(alist) - len(blist)
}

func parseInput(input string) [][][]interface{} {
	pairStrings := strings.Split(input, "\n\n")
	pairs := make([][][]interface{}, len(pairStrings))

	for i, pairString := range pairStrings {
		parts := strings.Split(pairString, "\n")
		pairs[i] = make([][]interface{}, 2)
		_ = json.Unmarshal([]byte(parts[0]), &pairs[i][0])
		_ = json.Unmarshal([]byte(parts[1]), &pairs[i][1])
	}

	return pairs
}
