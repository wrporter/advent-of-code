package convert

import "strconv"

func RuneToInt(rune uint8) int {
	return int(rune - '0')
}

func StringToInt(value string) int {
	valueInt64, _ := strconv.ParseInt(value, 10, 64)
	return int(valueInt64)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
