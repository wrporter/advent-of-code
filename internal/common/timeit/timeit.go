package timeit

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/out/color"
	"time"
)

func Track(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("🕒 %s took %s", name, elapsed)
}

func Report(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("🕒 %s%s%s", color.Blue, elapsed, color.Reset)
}
