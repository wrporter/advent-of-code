package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	image2 "github.com/wrporter/advent-of-code/solutions/2019/day8/internal/image"
)

func main() {
	dataLines, _ := file.ReadFile("./2019/day8/input.txt")
	i := image2.New(dataLines[0], 25, 6)
	//i := image.New("123456789012", 3, 2)
	//i := image.New("0222112222120000", 2, 2)
	fmt.Println(i.Validate())
	fmt.Println(image2.Draw(i.Decode()))
}
