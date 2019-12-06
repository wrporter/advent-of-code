package orbit

import (
	"fmt"
	"strings"
)

type OrbitMap struct{}
type Orbit struct {
	Parent string
	Child  string
}
type Node struct {
	Children  []string
	Parent    string
	NumOrbits int
}

const CenterOfMass = "COM"

func New() *OrbitMap {
	return &OrbitMap{}
}

func (o *OrbitMap) Count(orbits []string) int {
	tree := make(map[string]*Node)
	numOrbits := 0

	for _, orbitString := range orbits {
		orbit := parseOrbit(orbitString)

		if _, ok := tree[orbit.Parent]; !ok {
			tree[orbit.Parent] = &Node{
				[]string{orbit.Child},
				"",
				0,
			}
		} else {
			tree[orbit.Parent].Children = append(tree[orbit.Parent].Children, orbit.Child)
		}

		if _, ok := tree[orbit.Child]; !ok {
			tree[orbit.Child] = &Node{
				nil,
				orbit.Parent,
				tree[orbit.Parent].NumOrbits + 1,
			}
			numOrbits += tree[orbit.Child].NumOrbits
		} else {
			tree[orbit.Child].Parent = orbit.Parent
			numOrbits += addOrbit(tree, tree[orbit.Child], tree[orbit.Parent].NumOrbits+1)
		}
	}

	//count := print(tree, CenterOfMass)
	//fmt.Println(count)

	return numOrbits
}

func print(tree map[string]*Node, node string) int {
	count := 0
	count += tree[node].NumOrbits
	fmt.Printf("%s -> %v\n", node, tree[node])
	for _, child := range tree[node].Children {
		count += print(tree, child)
	}
	return count
}

func addOrbit(tree map[string]*Node, planet *Node, numOrbits int) int {
	numAdditionalOrbits := numOrbits
	planet.NumOrbits += numOrbits
	for _, child := range planet.Children {
		numAdditionalOrbits += addOrbit(tree, tree[child], numOrbits)
	}
	return numAdditionalOrbits
}

func parseOrbit(orbit string) *Orbit {
	planets := strings.Split(orbit, ")")
	return &Orbit{planets[0], planets[1]}
}
