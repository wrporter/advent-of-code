package main

import (
	"fmt"
	"strconv"
)

func main() {
	numPasswords := Solve(137683, 596253)
	fmt.Println(numPasswords)
}

func Solve(start int, end int) int {
	numPasswords := 0
	for candidate := start; candidate < end; candidate++ {
		if validPassword(start, end, candidate) {
			numPasswords++
		}
	}
	return numPasswords
}

func validPassword(start, end, candidate int) bool {
	candidateString := strconv.Itoa(candidate)
	return len(candidateString) == 6 &&
		candidate >= start && candidate <= end &&
		hasOnlyTwoConsecutiveDigits(candidateString) && // switch with hasTwoConsecutiveDigits for Part 1
		digitsNeverDecrease(candidateString)
}

func digitsNeverDecrease(candidate string) bool {
	lastDigit := runeToInt(candidate[0])
	for i := 1; i < len(candidate); i++ {
		curDigit := runeToInt(candidate[i])
		if curDigit < lastDigit {
			return false
		}
		lastDigit = curDigit
	}
	return true
}

func hasTwoConsecutiveDigits(candidate string) bool {
	lastDigit := candidate[0]
	for i := 1; i < len(candidate); i++ {
		curDigit := candidate[i]
		if curDigit == lastDigit {
			return true
		}
		lastDigit = curDigit
	}
	return false
}

func hasOnlyTwoConsecutiveDigits(candidate string) bool {
	lastDigit := candidate[0]
	numConsecutiveDigits := 1

	for i := 1; i < len(candidate); i++ {
		curDigit := candidate[i]

		if lastDigit != curDigit {
			if numConsecutiveDigits == 2 {
				return true
			}

			numConsecutiveDigits = 1
		}

		if curDigit == lastDigit {
			numConsecutiveDigits++
		}

		lastDigit = curDigit
	}

	if numConsecutiveDigits == 2 {
		return true
	}

	return false
}

func runeToInt(rune uint8) int {
	return int(rune - '0')
}
