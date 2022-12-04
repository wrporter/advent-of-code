package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/mystrings"
	"strings"
)

func incrementPassword(password string) string {
	remainder := byte(1)
	var next strings.Builder
	next.Grow(len(password) + 1)

	for i := len(password) - 1; i >= 0; i-- {
		if password[i]+remainder <= 'z' {
			next.WriteByte(password[i] + remainder)
			remainder = 0
		} else if i == 0 {
			next.WriteString("aa")
			// carry remainder
		} else {
			next.WriteByte('a')
			// carry remainder
		}
	}

	return mystrings.Reverse(next.String())
}

func isValid(password string) bool {
	return hasIncreasingStraight(password) &&
		!hasBadLetters(password) &&
		hasNonOverlappingPairs(password)
}

func hasIncreasingStraight(password string) bool {
	size := 3
	straightSize := 1

	for i := 0; i < len(password)-1; i++ {
		if password[i]+1 == password[i+1] {
			straightSize++
		} else {
			straightSize = 1
		}

		if straightSize == size {
			return true
		}
	}
	return false
}

func hasBadLetters(password string) bool {
	for i := 0; i < len(password); i++ {
		if password[i] == 'i' || password[i] == 'o' || password[i] == 'l' {
			return true
		}
	}
	return false
}

func hasNonOverlappingPairs(password string) bool {
	pairs := make(map[string]int)
	var prev rune

	for i, letter := range password {
		if i != 0 && letter == prev {
			pair := string(prev) + string(letter)
			if index, ok := pairs[pair]; ok {
				if index < i-2 {
					return true
				} else {
					continue
				}
			}
			pairs[pair] = i - 1
			if len(pairs) >= 2 {
				return true
			}
		}
		prev = letter
	}

	return false
}

func nextValid(password string) string {
	next := password
	for i := 0; ; i++ {
		next = incrementPassword(next)
		if isValid(next) {
			//fmt.Printf("Iterations: %d\n", i)
			return next
		}
	}
}

func main() {
	fmt.Println(nextValid("vzbxkghb"))
	fmt.Println(nextValid(nextValid("vzbxkghb")))
}
