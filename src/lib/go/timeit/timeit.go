package timeit

import (
	"aoc/src/lib/go/out/color"
	"fmt"
	"time"
)

func Track(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("🕒 %s took %s", name, elapsed)
}

func Report(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("🕒 %s%s%s\n", color.Blue, elapsed, color.Reset)
}
