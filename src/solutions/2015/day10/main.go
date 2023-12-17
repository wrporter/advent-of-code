package main

import (
	"fmt"
	"strconv"
	"strings"
)

func lookAndSay(sequence string) string {
	var next strings.Builder
	count := 0

	for i, cur := range sequence {
		count++
		if (i+1 < len(sequence) && cur != rune(sequence[i+1])) || i == len(sequence)-1 {
			next.WriteString(strconv.Itoa(count))
			next.WriteRune(cur)
			count = 0
		}
	}

	return next.String()
}

func lookAndSayRepeat(sequence string, times int) string {
	next := sequence
	for i := 0; i < times; i++ {
		next = lookAndSay(next)
	}
	return next
}

func main() {
	fmt.Println(len(lookAndSayRepeat("1321131112", 40)))
	fmt.Println(len(lookAndSayRepeat("1321131112", 50)))
}
