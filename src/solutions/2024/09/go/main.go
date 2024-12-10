package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"slices"
	"strconv"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	diskMap := convert.ToIntsV2(strings.Split(input, ""))
	var disk []int
	fileId := 0

	for i, length := range diskMap {
		if i%2 == 0 { // length of file
			disk = append(disk, repeat(fileId, length)...)
			fileId++
		} else { // length of free space
			disk = append(disk, repeat(-1, length)...)
		}
	}

	//fmt.Printf("Before: %s\n", toString(disk))

	lastFileId := fileId
	j := len(disk) - 1
	for i := diskMap[0]; i < j; {
		if disk[i] == -1 && disk[j] != -1 {
			disk[i] = disk[j]
			lastFileId = disk[j]
			disk[j] = -1
			i++
			j--
		} else {
			if disk[i] != -1 {
				fileId := disk[i]
				lengthOfFile := diskMap[fileId*2]
				i += lengthOfFile
			}
			if disk[j] == -1 {
				lengthOfFreeSpace := diskMap[lastFileId*2-1]
				j -= lengthOfFreeSpace
			}
		}
	}

	//fmt.Printf("After:  %s\n", toString(disk))
	//fmt.Printf("Cut:    %s\n", toString(disk[:j+1]))

	checksum := 0
	for i, fileId := range disk[:j+1] {
		checksum += i * fileId
	}

	return checksum
}

type memory struct {
	index  int
	fileId int
	isFile bool
	length int
}

func part2(input string, _ ...interface{}) interface{} {
	diskMap := convert.ToIntsV2(strings.Split(input, ""))
	disk := make([]*memory, len(diskMap))
	fileId := 0
	index := 0
	var files []*memory

	for i, length := range diskMap {
		if i%2 == 0 { // length of file
			disk[i] = &memory{
				index:  index,
				fileId: fileId,
				isFile: true,
				length: length,
			}
			files = append(files, disk[i])
			fileId++
		} else { // length of free space
			disk[i] = &memory{
				index:  index,
				isFile: false,
				length: length,
			}
		}
		index += length
	}

	//fmt.Println(diskToString(disk))

	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		index := 0

		fileIndex := file.index
		fileDiskIndex := -1
		for diskIndex, chunk := range disk {
			if chunk.fileId == file.fileId {
				fileDiskIndex = diskIndex
				break
			}
		}

		for j := 0; j < len(disk) && j < fileDiskIndex; {
			chunk := disk[j]
			if !chunk.isFile && (chunk.length >= file.length) {
				// update the file memory index
				file.index = index

				if chunk.length > file.length {
					// update the free space size
					chunk.length -= file.length
					chunk.index += file.length
					// insert the file at the free space
					disk = slices.Insert(disk, j, file)
					// remove the file from its original disk position
					disk[fileDiskIndex+1] = &memory{
						index:  fileIndex,
						isFile: false,
						length: file.length,
					}
				} else {
					// remove the free space
					disk[j] = file
					// remove the file from its original disk position
					disk[fileDiskIndex] = &memory{
						index:  fileIndex,
						isFile: false,
						length: file.length,
					}
				}

				//fmt.Println(diskToString(disk))
				break
			} else {
				// scan for next available free space
				j++
			}
			index += chunk.length
		}
	}

	//fmt.Println(diskToString(disk))

	checksum := 0
	for _, file := range files {
		for i := file.index; i < file.index+file.length; i++ {
			checksum += i * file.fileId
		}
	}

	return checksum
}

func diskToString(disk []*memory) string {
	str := ""
	for _, chunk := range disk {
		for i := 0; i < chunk.length; i++ {
			if chunk.isFile {
				str += strconv.Itoa(chunk.fileId)
			} else {
				str += "."
			}
		}
	}
	return str
}

func repeat(value int, count int) []int {
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = value
	}
	return result
}

func toString(disk []int) string {
	result := make([]rune, len(disk))
	for i, value := range disk {
		if value == -1 {
			result[i] = '.'
		} else {
			result[i] = rune(strconv.Itoa(value)[0])
		}
	}
	return string(result)
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2024, Day: 9, Part1: part1, Part2: part2}
}
