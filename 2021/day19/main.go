package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 19
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/sample-input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	chunks := input

	var scanners [][]vector
	for _, chunk := range chunks {
		var beacons []vector
		for _, line := range chunk[1:] {
			var x, y, z int
			_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
			if err != nil {
				panic(err)
			}
			beacons = append(beacons, vector{x, y, z})
		}
		scanners = append(scanners, beacons)
	}

	aligned := scanners[0]
	remaining := scanners[1:]

	shifts := []vector{{0, 0, 0}}

	for len(remaining) > 0 {
		gotOne := false
		for i, next := range remaining {
			mat, shift, overlap := FindBestRotatedOverlap(aligned, next)
			if overlap >= 12 {
				shifted := util.Map(next, func(p vector) vector { return p.CrossProduct(mat).Sub(shift) })
				// XXX more efficient to de-dupe here
				aligned = DeDupe(append(aligned, shifted...))
				remaining = append(remaining[:i], remaining[i+1:]...)
				//fmt.Printf("Got one! %d/%d\n", len(aligned), len(remaining))
				shifts = append(shifts, shift)
				gotOne = true
				break
			}
		}

		if !gotOne {
			panic("Unable to align!")
		}
	}

	allvectors := map[string]vector{}
	for _, point := range aligned {
		// for _, point := range a {
		allvectors[point.String()] = point
		// }
	}

	return len(allvectors)
}

func parseInput(input []string) []*scanner {
	var scanners []*scanner
	var s *scanner

	for _, line := range input {
		if line == "" {
			continue
		}
		if line[:3] == "---" {
			s = &scanner{}
			scanners = append(scanners, s)
			continue
		}
		v, _ := convert.ToInts(strings.Split(line, ","))
		s.addBeacon(v[0], v[1], v[2])
	}
	return scanners
}

func part2(input []string) interface{} {
	return 0
}

type scanner struct {
	*vector
	beacons []*beacon
}

func (s *scanner) addBeacon(x, y, z int) {
	b := &beacon{
		id:     len(s.beacons),
		vector: &vector{x: x, y: y, z: z},
	}
	s.beacons = append(s.beacons, b)
}

type beacon struct {
	id int
	*vector
}

type vector struct {
	x, y, z int
}

func (v vector) Add(o vector) vector {
	return vector{v.x + o.x, v.y + o.y, v.z + o.z}
}
func (v vector) Sub(o vector) vector {
	return vector{v.x - o.x, v.y - o.y, v.z - o.z}
}
func (v vector) Scale(o vector) vector {
	return vector{v.x * o.x, v.y * o.y, v.z * o.z}
}
func (v vector) Rot90Z() vector {
	return vector{-v.y, v.x, v.z}
}
func (v vector) Rot90Y() vector {
	return vector{-v.z, v.y, v.x}
}
func (v vector) Rot90X() vector {
	return vector{v.x, -v.z, v.y}
}
func (v vector) String() string {
	return fmt.Sprintf("(%d,%d,%d)", v.x, v.y, v.z)
}
func (v vector) Manhattan(b vector) int {
	return ints.Abs(v.x-b.x) + ints.Abs(v.y-b.y) + ints.Abs(v.z-b.z)
}

type matrix [][]int

func (v vector) CrossProduct(m matrix) vector {
	return vector{
		v.x*m[0][0] + v.y*m[0][1] + v.z*m[0][2],
		v.x*m[1][0] + v.y*m[1][1] + v.z*m[1][2],
		v.x*m[2][0] + v.y*m[2][1] + v.z*m[2][2],
	}
}

var ID = matrix{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
var FLIPS = []matrix{
	{{-1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
	{{1, 0, 0}, {0, -1, 0}, {0, 0, 1}},
	{{1, 0, 0}, {0, 1, 0}, {0, 0, -1}},
}
var ROTS = []matrix{
	{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}}, // rotate 90 z
	{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}}, // rotate 90 y
	{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}}, // rotate 90 y
}

func (m matrix) CrossProduct(r matrix) matrix {
	res := matrix{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res[i][j] = m[i][0]*r[0][j] + m[i][1]*r[1][j] + m[i][2]*r[2][j]
		}
	}
	return res
}
func (m matrix) Clone() matrix {
	return ID.CrossProduct(m)
}
func (m matrix) String() string {
	return fmt.Sprintf("%d,%d,%d; %d,%d,%d; %d,%d,%d",
		m[0][0], m[0][1], m[0][2],
		m[1][0], m[1][1], m[1][2],
		m[2][0], m[2][1], m[2][2],
	)
}

func FindAllOrientations() []matrix {
	orients := map[string]matrix{}
	for nx := 0; nx < 4; nx++ {
		for ny := 0; ny < 4; ny++ {
			for nz := 0; nz < 4; nz++ {
				for fx := 0; fx <= 1; fx++ {
					for fy := 0; fy <= 1; fy++ {
						for fz := 0; fz <= 1; fz++ {
							m := ID
							for n := 0; n < nx; n++ {
								m = m.CrossProduct(ROTS[0])
							}
							for n := 0; n < ny; n++ {
								m = m.CrossProduct(ROTS[1])
							}
							for n := 0; n < nz; n++ {
								m = m.CrossProduct(ROTS[2])
							}
							/*
								if fx > 0 {
									m = m.Mult(FLIPS[0])
								}
								if fy > 0 {
									m = m.Mult(FLIPS[1])
								}
								if fz > 0 {
									m = m.Mult(FLIPS[2])
								}
							*/
							orients[m.String()] = m
						}
					}
				}
			}
		}
	}

	return getMatrixValues(orients)
}

var ROTATIONS []matrix

func init() {
	ROTATIONS = FindAllOrientations()
}
func NumOverlapping(aMap map[vector]bool, bs []vector) int {
	n := 0
	for _, b := range bs {
		if _, ok := aMap[b]; ok {
			n++
		}
	}
	return n
}

// Returns shift of bs relative to as, number of overlapping beacons
func FindBestOverlap(as, bs []vector) (vector, int) {
	aMap := make(map[vector]bool)
	for _, a := range as {
		aMap[a] = true
	}

	// Consider that each pair might be the same
	bestOverlap := 0
	var bestShift vector
	for i, a := range as {
		for j, b := range bs {
			if j > i {
				continue
			}

			// Assume a and b are the same
			shift := b.Sub(a)
			newBs := util.Map(bs, func(p vector) vector { return p.Sub(shift) })
			overlap := NumOverlapping(aMap, newBs)
			if overlap > bestOverlap {
				// Ties are potentially problematic
				// And the greedy strategy might not work.
				bestOverlap = overlap
				bestShift = shift
				if bestOverlap >= 12 {
					return bestShift, bestOverlap
				}
			}
		}
	}
	return bestShift, bestOverlap
}

func DeDupe(points []vector) []vector {
	m := map[string]vector{}
	for _, p := range points {
		s := p.String()
		_, ok := m[s]
		if !ok {
			m[s] = p
		}
	}
	return getVectorValues(m)
}

func FindBestRotatedOverlap(as, bs []vector) (matrix, vector, int) {
	bestOverlap := 0
	var bestShift vector
	bestMat := ID

	for _, m := range ROTATIONS {
		rotBs := util.Map(bs, func(p vector) vector { return p.CrossProduct(m) })
		shift, overlap := FindBestOverlap(as, rotBs)
		if overlap > bestOverlap {
			bestOverlap = overlap
			bestShift = shift
			bestMat = m
			if bestOverlap >= 12 {
				break
			}
		}
	}

	return bestMat, bestShift, bestOverlap
}

func getVectorValues(m map[string]vector) []vector {
	var values []vector
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

func getMatrixValues(m map[string]matrix) []matrix {
	var values []matrix
	for _, value := range m {
		values = append(values, value)
	}
	return values
}
