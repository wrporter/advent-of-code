package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"math"
	"strings"
)

type Node struct {
	Parent   *Node
	Children map[string]*Node
	Name     string
	Size     int
	IsDir    bool
}

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	fs := createFileSystem(input)

	sum := 0
	walk(fs, func(dir *Node) {
		if dir.Size <= 100_000 {
			sum += dir.Size
		}
	})

	return sum
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	fs := createFileSystem(input)
	diskSpace := 70_000_000
	spaceRequired := 30_000_000
	free := diskSpace - fs.Size
	needed := spaceRequired - free

	smallest := math.MaxInt
	walk(fs, func(dir *Node) {
		if dir.Size >= needed && dir.Size < smallest {
			smallest = dir.Size
		}
	})

	return smallest
}

func createFileSystem(input string) *Node {
	fs := &Node{Name: "/", Children: make(map[string]*Node), IsDir: true}
	cwd := fs

	for _, line := range strings.Split(input, "\n")[1:] {
		parts := strings.Split(line, " ")

		if strings.HasPrefix(line, "$ cd") {
			dir := parts[2]
			if dir == ".." {
				cwd = cwd.Parent
			} else {
				cwd = cwd.Children[dir]
			}
		} else if line != "$ ls" {
			name := parts[1]
			if parts[0] == "dir" {
				cwd.Children[name] = &Node{
					Parent:   cwd,
					Children: make(map[string]*Node),
					Name:     name,
					IsDir:    true,
				}
			} else {
				size := convert.StringToInt(parts[0])
				cwd.Children[name] = &Node{
					Parent: cwd,
					Name:   name,
					IsDir:  false,
					Size:   size,
				}

				for current := cwd; current != nil; current = current.Parent {
					current.Size += size
				}
			}
		}
	}
	return fs
}
func walk(node *Node, onDir func(node *Node)) {
	if node.IsDir {
		onDir(node)
		for _, child := range node.Children {
			walk(child, onDir)
		}
	}
}
