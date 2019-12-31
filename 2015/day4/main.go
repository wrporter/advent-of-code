package main

import (
	"crypto/md5"
	"fmt"
)

func with5LeadingZeroes(secret string) int {
	for value := 1; ; value++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", secret, value)))
		if hash[0] == 0 && hash[1] == 0 && hash[2] < 16 {
			return value
		}
	}
}

func with6LeadingZeroes(secret string) int {
	for value := 1; ; value++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", secret, value)))
		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
			return value
		}
	}
}

func main() {
	fmt.Println(with5LeadingZeroes("yzbqklnj"))
	fmt.Println(with6LeadingZeroes("yzbqklnj"))
}
