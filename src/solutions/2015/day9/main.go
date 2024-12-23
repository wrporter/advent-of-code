package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"fmt"
	"regexp"
)

type Destination struct {
	to       string
	distance int
}

var regex = regexp.MustCompile(`^([a-zA-Z]+) to ([a-zA-Z]+) = (\d+)$`)

func parse(distanceStrings []string) map[string][]Destination {
	tickets := make(map[string][]Destination)
	for _, routeString := range distanceStrings {
		match := regex.FindStringSubmatch(routeString)
		tickets[match[1]] = append(tickets[match[1]], Destination{
			to:       match[2],
			distance: convert.StringToInt(match[3]),
		})
		tickets[match[2]] = append(tickets[match[2]], Destination{
			to:       match[1],
			distance: convert.StringToInt(match[3]),
		})
	}
	return tickets
}

type Node struct {
	from     *Node
	location string
	distance int
	visited  map[string]bool
}

func getShortestRoute(distanceStrings []string) ([]string, int, []string, int) {
	tickets := parse(distanceStrings)
	shortest := ints.MaxInt
	longest := 0
	var shortestRoute []string
	var longestRoute []string

	var queue []*Node
	for start := range tickets {
		queue = append(queue, &Node{
			location: start,
			visited:  make(map[string]bool),
		})
	}

	var node *Node
	for len(queue) > 0 {
		node, queue = queue[0], queue[1:]
		node.visited[node.location] = true

		if allVisited(tickets, node.visited) {
			if node.distance < shortest {
				shortest = node.distance
				shortestRoute = buildRoute(node)
				continue
			}
			if node.distance > longest {
				longest = node.distance
				longestRoute = buildRoute(node)
				continue
			}
		}

		for _, destination := range tickets[node.location] {
			if !node.visited[destination.to] {
				queue = append(queue, &Node{
					from:     node,
					location: destination.to,
					distance: node.distance + destination.distance,
					visited:  cloneMap(node.visited),
				})
			}
		}
	}

	return shortestRoute, shortest, longestRoute, longest
}

func buildRoute(node *Node) []string {
	var route []string
	for node.from != nil {
		route = prepend(route, node.location)
		node = node.from
	}
	route = prepend(route, node.location)
	return route
}

func allVisited(tickets map[string][]Destination, visited map[string]bool) bool {
	for k := range tickets {
		if !visited[k] {
			return false
		}
	}
	return true
}

func main() {
	lines, _ := file.ReadFile("./2015/day9/input.txt")
	fmt.Println(getShortestRoute(lines))
	// [Tristram AlphaCentauri Norrath Straylight Faerun Snowdin Tambi Arbre] 141
	// [AlphaCentauri Arbre Tristram Snowdin Straylight Tambi Norrath Faerun] 736
}

func prepend(array []string, value string) []string {
	array = append(array, "")
	copy(array[1:], array)
	array[0] = value
	return array
}

func cloneMap(originalMap map[string]bool) map[string]bool {
	clonedMap := make(map[string]bool)
	for key, value := range originalMap {
		clonedMap[key] = value
	}
	return clonedMap
}
