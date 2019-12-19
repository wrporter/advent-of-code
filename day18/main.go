package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day18/internal/vault"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"log"
	"time"
)

func main() {
	lines, _ := file.ReadFile("./day18/input.txt")
	//lines := []string{
	//	"#########",
	//	"#b.A.@.a#",
	//	"#########",
	//}
	defer timeTrack(time.Now(), "find keys")
	v := vault.New(lines)
	fmt.Println(v.MinSteps())
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
