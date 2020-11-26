package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
)

func part1(secret string) string {
	var password []rune

	for value := 1; len(password) < 8; value++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", secret, value)))
		hexValue := hex.EncodeToString(hash[:])
		if hexValue[:5] == "00000" {
			password = append(password, rune(hexValue[5]))
		}
	}

	return string(password)
}

func part2(secret string) string {
	password := make([]rune, 8)

	for value := 1; !allAreSet(password); value++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", secret, value)))
		hexValue := hex.EncodeToString(hash[:])
		index := conversion.RuneToInt(hexValue[5])

		if hexValue[:5] == "00000" && index < 8 && password[index] == 0 {
			value := rune(hexValue[6])
			password[index] = value
		}
	}

	return string(password)
}

func allAreSet(array []rune) bool {
	for _, value := range array {
		if value == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(part1("cxdnnyjw"))
	fmt.Println(part2("cxdnnyjw"))
}
