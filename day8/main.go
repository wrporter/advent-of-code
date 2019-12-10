package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day8/internal/image"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
)

func main() {
	dataLines, _ := file.ReadFile("./day8/input.txt")
	i := image.New(dataLines[0], 25, 6)
	//i := image.New("123456789012", 3, 2)
	//i := image.New("0222112222120000", 2, 2)
	fmt.Println(i.Validate())
	image.Print(i.Decode())
}
