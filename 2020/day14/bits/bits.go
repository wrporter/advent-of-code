package main

import (
	"fmt"
	"strconv"
)

func main() {
	value := 7
	converted := strconv.FormatInt(int64(value), 2)
	fmt.Println(converted)

	back, _ := strconv.ParseInt(converted, 2, 64)
	fmt.Println(back)
}
