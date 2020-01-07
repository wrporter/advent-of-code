package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/mystrings"
	"strings"
)

func incrementPassword(password string) string {
	remainder := byte(1)
	pass := mystrings.Reverse(password)
	var next strings.Builder

	for i := 0; i < len(pass); i++ {
		if pass[i]+remainder <= 'z' {
			next.WriteByte(pass[i] + remainder)
			remainder = 0
		} else if i+1 == len(pass) {
			remainder = pass[i] + remainder - 'z'
			next.WriteByte('a')
			next.WriteByte('a')
		} else {
			remainder = pass[i] + remainder - 'z'
			next.WriteByte('a')
		}
	}

	return mystrings.Reverse(next.String())
}

func main() {
	password := "a"
	for i := 0; i < 1000; i++ {
		password = incrementPassword(password)
		fmt.Println(password)
	}
}
