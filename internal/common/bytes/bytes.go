package bytes

import "unicode"

func ToUpper(b byte) byte {
	return byte(unicode.ToUpper(rune(b)))
}

func ToLower(b byte) byte {
	return byte(unicode.ToLower(rune(b)))
}

func IsUpper(b byte) bool {
	return unicode.IsUpper(rune(b))
}

func IsLower(b byte) bool {
	return unicode.IsLower(rune(b))
}
