package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"regexp"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	workflows, partRatings := parse(input)
	sum := 0

	for _, partRating := range partRatings {
		current := "in"

		for current != "A" && current != "R" {
			for _, rule := range workflows[current].Workflow {
				rating, ok := partRating[rule.Rating]
				if rule.Rating == "" {
					current = rule.Result
					break
				} else if ok &&
					(rule.Operator == ">" && rating > rule.Value) ||
					(rule.Operator == "<" && rating < rule.Value) {
					current = rule.Result
					break
				}
			}
		}

		if current == "A" {
			sum += partRating["total"]
		}
	}

	return sum
}

func part2(input string, _ ...interface{}) interface{} {
	workflows, _ := parse(input)
	var count func(current string, ranges map[string]*Range) int

	count = func(current string, ranges map[string]*Range) int {
		if current == "R" {
			return 0
		}

		if current == "A" {
			product := 1
			for _, r := range ranges {
				product *= r.end - r.start + 1
			}
			return product
		}

		result := 0
		for _, rule := range workflows[current].Workflow {
			next := make(map[string]*Range)
			for k, v := range ranges {
				next[k] = &Range{v.start, v.end}
			}

			if rule.Rating == "" {
				result += count(rule.Result, ranges)
			} else if rule.Operator == "<" {
				next[rule.Rating].end = rule.Value - 1
				ranges[rule.Rating].start = rule.Value
				result += count(rule.Result, next)
			} else if rule.Operator == ">" {
				next[rule.Rating].start = rule.Value + 1
				ranges[rule.Rating].end = rule.Value
				result += count(rule.Result, next)
			}
		}
		return result
	}

	return count("in", map[string]*Range{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	})
}

func parse(input string) (map[string]Part, []map[string]int) {
	chunks := strings.Split(input, "\n\n")
	ratingsStr := strings.Split(chunks[1], "\n")

	ratings := make([]map[string]int, len(ratingsStr))
	for i, ratingStr := range ratingsStr {
		valuesStr := ratingStr[1 : len(ratingStr)-1]
		values := make(map[string]int)
		total := 0
		for _, valueStr := range strings.Split(valuesStr, ",") {
			value := convert.StringToInt(valueStr[2:])
			values[string(valueStr[0])] = value
			total += value
		}
		values["total"] = total
		ratings[i] = values
	}

	workflows := make(map[string]Part)
	for _, partStr := range strings.Split(chunks[0], "\n") {
		match := partRegex.FindStringSubmatch(partStr)
		workflowStr := match[2]
		rulesStr := strings.Split(workflowStr, ",")
		part := Part{Name: match[1], Workflow: make([]Rule, len(rulesStr))}

		for i, ruleStr := range rulesStr {
			match = ruleRegex.FindStringSubmatch(ruleStr)
			if match[2] != "" {
				part.Workflow[i] = Rule{
					Rating:   match[2],
					Operator: match[3],
					Value:    convert.StringToInt(match[4]),
					Result:   match[5],
				}
			} else {
				part.Workflow[i] = Rule{Result: match[1]}
			}
		}

		workflows[part.Name] = part
	}
	return workflows, ratings
}

var partRegex = regexp.MustCompile(`([a-z]+)\{(.+)}`)
var ruleRegex = regexp.MustCompile(`(([xmas])([<>])(\d+):([a-z]+|A|R)|[a-z]+|A|R)`)

type Part struct {
	Name     string
	Workflow []Rule
}

type Rule struct {
	Rating   string
	Operator string
	Value    int
	Result   string
}

type Range struct {
	start, end int
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 19, Part1: part1, Part2: part2}
}
