package bytes

import (
	"unicode"
)

func IsLetter(b byte) bool {
	return unicode.IsLetter(rune(b))
}

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

func CopyAdd(bytes []byte, b byte) []byte {
	newArray := make([]byte, len(bytes)+1)
	copy(newArray, bytes)
	newArray[len(bytes)] = b
	return newArray
}

func Copy(bytes []byte) []byte {
	newArray := make([]byte, len(bytes))
	copy(newArray, bytes)
	return newArray
}
