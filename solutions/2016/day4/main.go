package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"regexp"
	"sort"
	"strings"
)

var regex = regexp.MustCompile(`([a-z\-]+)(\d+)\[([a-z]+)]`)

type RoomChecksum struct {
	Checksum     []rune
	LetterCounts map[rune]int
}

func NewRoomChecksum(encryptedName string) RoomChecksum {
	letterCounts := make(map[rune]int)

	for _, char := range encryptedName {
		if _, ok := letterCounts[char]; !ok {
			letterCounts[char] = 1
		} else {
			letterCounts[char]++
		}
	}

	uniqueCharacters := make([]rune, len(letterCounts))
	index := 0
	for key := range letterCounts {
		uniqueCharacters[index] = key
		index++
	}

	return RoomChecksum{uniqueCharacters, letterCounts}
}

func (s RoomChecksum) Less(i, j int) bool {
	if s.LetterCounts[s.Checksum[i]] > s.LetterCounts[s.Checksum[j]] {
		return true
	} else if s.LetterCounts[s.Checksum[i]] == s.LetterCounts[s.Checksum[j]] {
		return s.Checksum[i] < s.Checksum[j]
	}
	return false
}

func (s RoomChecksum) Swap(i, j int) {
	s.Checksum[i], s.Checksum[j] = s.Checksum[j], s.Checksum[i]
}

func (s RoomChecksum) Len() int {
	return len(s.Checksum)
}

func (s RoomChecksum) GetChecksum() string {
	return string(s.Checksum[:5])
}

func part1(input []string) int {
	result := 0

	for _, room := range input {
		match := regex.FindStringSubmatch(room)
		encryptedName := strings.ReplaceAll(match[1], "-", "")
		sectorID := convert.StringToInt(match[2])
		checksum := match[3]

		roomChecksum := NewRoomChecksum(encryptedName)
		sort.Sort(roomChecksum)

		if checksum == roomChecksum.GetChecksum() {
			result += sectorID
		}
	}

	return result
}

type ShiftCipher struct {
	EncryptedValue string
	ShiftAmount    int
}

func NewShiftCipher(encryptedValue string, shiftAmount int) ShiftCipher {
	return ShiftCipher{encryptedValue, shiftAmount}
}

func (s ShiftCipher) Decrypt() string {
	decryptedValue := make([]rune, len(s.EncryptedValue))

	for i, char := range s.EncryptedValue {
		if char < 'a' || char > 'z' {
			decryptedValue[i] = ' '
		}
		decryptedValue[i] = caesar(char, rune(s.ShiftAmount))
	}

	return string(decryptedValue)
}

func caesar(r rune, shift rune) rune {
	if r < 'a' || r > 'z' {
		return r
	}
	return 'a' + ((r + shift - 'a') % ('z' - 'a' + 1))
}

func part2(input []string, roomName string) int {
	for _, room := range input {
		match := regex.FindStringSubmatch(room)
		encryptedName := match[1]
		sectorID := convert.StringToInt(match[2])
		checksum := match[3]

		roomChecksum := NewRoomChecksum(strings.ReplaceAll(encryptedName, "-", ""))
		sort.Sort(roomChecksum)

		if checksum == roomChecksum.GetChecksum() {
			if NewShiftCipher(encryptedName, sectorID).Decrypt() == roomName {
				return sectorID
			}
		}
	}

	return -1
}

func main() {
	input, _ := file.ReadFile("./2016/day4/input.txt")
	answer1 := part1(input)
	answer2 := part2(input, "northpole-object-storage-")
	fmt.Println(answer1)
	fmt.Println(answer2)
}
