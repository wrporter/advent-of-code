package orbit

import (
	"strings"
)

type OrbitFactory struct{}
type Orbit struct {
	Parent string
	Child  string
}
type Node struct {
	Children  []string
	Parent    string
	NumOrbits int
}
type OrbitMap struct {
	Tree      map[string]*Node
	NumOrbits int
}

const CenterOfMass = "COM"
const Santa = "SAN"
const You = "YOU"

func New() *OrbitFactory {
	return &OrbitFactory{}
}

func (o *OrbitFactory) Generate(orbits []string) *OrbitMap {
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

	return &OrbitMap{tree, numOrbits}
}

func (o *OrbitFactory) Distance(tree map[string]*Node, planet1 string, planet2 string) int {
	ancestor := findNearestAncestor(tree, planet1, planet2)
	parentOrbits := tree[ancestor].NumOrbits
	return tree[planet1].NumOrbits - parentOrbits + tree[planet2].NumOrbits - parentOrbits - 2
}

func findNearestAncestor(tree map[string]*Node, planet1 string, planet2 string) string {
	var commonAncestor string
	ancestor1 := tree[planet1].Parent
	ancestor2 := tree[planet2].Parent
	ancestors1 := map[string]bool{ancestor1: true}
	ancestors2 := map[string]bool{ancestor2: true}

	for commonAncestor == "" {
		if ancestors1[ancestor2] {
			commonAncestor = ancestor2
			break
		}
		if ancestors2[ancestor1] {
			commonAncestor = ancestor1
			break
		}
		if ancestor1 != CenterOfMass {
			ancestor1 = tree[ancestor1].Parent
		}
		if ancestor2 != CenterOfMass {
			ancestor2 = tree[ancestor2].Parent
		}
		ancestors1[ancestor1] = true
		ancestors2[ancestor2] = true
	}

	return commonAncestor
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
