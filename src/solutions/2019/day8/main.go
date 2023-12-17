package main

import (
	"aoc/src/lib/go/file"
	image2 "aoc/src/solutions/2019/day8/lib/image"
	"fmt"
)

func main() {
	dataLines, _ := file.ReadFile("./2019/day8/input.txt")
	i := image2.New(dataLines[0], 25, 6)
	//i := image.New("123456789012", 3, 2)
	//i := image.New("0222112222120000", 2, 2)
	fmt.Println(i.Validate())
	fmt.Println(image2.Draw(i.Decode()))
}
