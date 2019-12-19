package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day18/internal/vault"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
)

func main() {
	lines, _ := file.ReadFile("./day18/input.txt")
	//lines := []string{
	//	"#################",
	//	"#i.G..c...e..H.p#",
	//	"########.########",
	//	"#j.A..b...f..D.o#",
	//	"########@########",
	//	"#k.E..a...g..B.n#",
	//	"########.########",
	//	"#l.F..d...h..C.m#",
	//	"#################",
	//}

	v := vault.New(lines)
	fmt.Println(v.MinSteps())
}
