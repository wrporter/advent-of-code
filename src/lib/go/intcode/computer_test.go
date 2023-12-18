package intcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestComputer_Run(t *testing.T) {
	tests := []struct {
		description string
		program     []int
		input       []int
		wantOutput  []int
	}{
		{
			"position mode: 5 is not equal to 8",
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{5},
			[]int{0},
		},
		{
			"position mode: 8 is equal to 8",
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{8},
			[]int{1},
		},
		{
			"position mode: 9 is not equal to 8",
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{9},
			[]int{0},
		},
		{
			"position mode: 5 is less than 8",
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{5},
			[]int{1},
		},
		{
			"position mode: 8 is not less than 8",
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{8},
			[]int{0},
		},
		{
			"position mode: 9 is not less than 8",
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{9},
			[]int{0},
		},
		{
			"immediate mode: 5 is not equal to 8",
			[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			[]int{5},
			[]int{0},
		},
		{
			"immediate mode: 8 is equal to 8",
			[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			[]int{8},
			[]int{1},
		},
		{
			"immediate mode: 9 is not equal to 8",
			[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			[]int{9},
			[]int{0},
		},
		{
			"immediate mode: 5 is less than 8",
			[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			[]int{5},
			[]int{1},
		},
		{
			"immediate mode: 8 is not less than 8",
			[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			[]int{8},
			[]int{0},
		},
		{
			"immediate mode: 9 is not less than 8",
			[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			[]int{9},
			[]int{0},
		},
		{
			"position mode: input is 0",
			[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			[]int{0},
			[]int{0},
		},
		{
			"position mode: input is non-zero",
			[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			[]int{5},
			[]int{1},
		},
		{
			"immedate mode: input is 0",
			[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			[]int{0},
			[]int{0},
		},
		{
			"immedate mode: input is non-zero",
			[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			[]int{5},
			[]int{1},
		},
		{
			"input is less than 8",
			[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			[]int{5},
			[]int{999},
		},
		{
			"input equals 8",
			[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			[]int{8},
			[]int{1000},
		},
		{
			"input is greater than 8",
			[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			[]int{9},
			[]int{1001},
		},
		{
			"TEST diagnostic program: air conditioner unit",
			[]int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 72, 20, 224, 1001, 224, -1440, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1002, 147, 33, 224, 101, -3036, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 32, 90, 225, 101, 65, 87, 224, 101, -85, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 223, 224, 223, 1102, 33, 92, 225, 1102, 20, 52, 225, 1101, 76, 89, 225, 1, 117, 122, 224, 101, -78, 224, 224, 4, 224, 102, 8, 223, 223, 101, 1, 224, 224, 1, 223, 224, 223, 1102, 54, 22, 225, 1102, 5, 24, 225, 102, 50, 84, 224, 101, -4600, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 3, 224, 224, 1, 223, 224, 223, 1102, 92, 64, 225, 1101, 42, 83, 224, 101, -125, 224, 224, 4, 224, 102, 8, 223, 223, 101, 5, 224, 224, 1, 224, 223, 223, 2, 58, 195, 224, 1001, 224, -6840, 224, 4, 224, 102, 8, 223, 223, 101, 1, 224, 224, 1, 223, 224, 223, 1101, 76, 48, 225, 1001, 92, 65, 224, 1001, 224, -154, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 329, 101, 1, 223, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 344, 1001, 223, 1, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 359, 1001, 223, 1, 223, 8, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 374, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 389, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 404, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 419, 101, 1, 223, 223, 1008, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 434, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 449, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 464, 1001, 223, 1, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 479, 101, 1, 223, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 494, 1001, 223, 1, 223, 7, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 509, 101, 1, 223, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 539, 1001, 223, 1, 223, 108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 554, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 569, 101, 1, 223, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 584, 1001, 223, 1, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 599, 1001, 223, 1, 223, 1007, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 614, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 629, 101, 1, 223, 223, 1108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 644, 1001, 223, 1, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 659, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226},
			[]int{1},
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 11933517},
		},
		{
			"TEST diagnostic program: thermal radiator controller",
			[]int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 72, 20, 224, 1001, 224, -1440, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1002, 147, 33, 224, 101, -3036, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 32, 90, 225, 101, 65, 87, 224, 101, -85, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 223, 224, 223, 1102, 33, 92, 225, 1102, 20, 52, 225, 1101, 76, 89, 225, 1, 117, 122, 224, 101, -78, 224, 224, 4, 224, 102, 8, 223, 223, 101, 1, 224, 224, 1, 223, 224, 223, 1102, 54, 22, 225, 1102, 5, 24, 225, 102, 50, 84, 224, 101, -4600, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 3, 224, 224, 1, 223, 224, 223, 1102, 92, 64, 225, 1101, 42, 83, 224, 101, -125, 224, 224, 4, 224, 102, 8, 223, 223, 101, 5, 224, 224, 1, 224, 223, 223, 2, 58, 195, 224, 1001, 224, -6840, 224, 4, 224, 102, 8, 223, 223, 101, 1, 224, 224, 1, 223, 224, 223, 1101, 76, 48, 225, 1001, 92, 65, 224, 1001, 224, -154, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 329, 101, 1, 223, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 344, 1001, 223, 1, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 359, 1001, 223, 1, 223, 8, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 374, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 389, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 404, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 419, 101, 1, 223, 223, 1008, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 434, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 449, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 464, 1001, 223, 1, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 479, 101, 1, 223, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 494, 1001, 223, 1, 223, 7, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 509, 101, 1, 223, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 539, 1001, 223, 1, 223, 108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 554, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 569, 101, 1, 223, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 584, 1001, 223, 1, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 599, 1001, 223, 1, 223, 1007, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 614, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 629, 101, 1, 223, 223, 1108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 644, 1001, 223, 1, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 659, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226},
			[]int{5},
			[]int{10428568},
		},
		{
			"relative mode: copy program",
			[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			[]int{},
			[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			"relative mode: output 16-digit number",
			[]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			[]int{},
			[]int{1219070632396864},
		},
		{
			"relative mode: output large number in middle",
			[]int{104, 1125899906842624, 99},
			[]int{},
			[]int{1125899906842624},
		},
		{
			"BOOST: test all opcodes",
			[]int{1102, 34463338, 34463338, 63, 1007, 63, 34463338, 63, 1005, 63, 53, 1101, 3, 0, 1000, 109, 988, 209, 12, 9, 1000, 209, 6, 209, 3, 203, 0, 1008, 1000, 1, 63, 1005, 63, 65, 1008, 1000, 2, 63, 1005, 63, 904, 1008, 1000, 0, 63, 1005, 63, 58, 4, 25, 104, 0, 99, 4, 0, 104, 0, 99, 4, 17, 104, 0, 99, 0, 0, 1101, 37, 0, 1005, 1101, 30, 0, 1013, 1102, 1, 33, 1019, 1102, 1, 25, 1003, 1102, 1, 28, 1018, 1101, 26, 0, 1006, 1102, 1, 866, 1029, 1101, 760, 0, 1023, 1102, 39, 1, 1012, 1102, 23, 1, 1009, 1101, 281, 0, 1026, 1102, 1, 20, 1011, 1102, 1, 34, 1008, 1101, 0, 36, 1017, 1101, 38, 0, 1000, 1102, 0, 1, 1020, 1102, 278, 1, 1027, 1101, 21, 0, 1010, 1102, 875, 1, 1028, 1101, 0, 212, 1025, 1102, 1, 1, 1021, 1102, 1, 24, 1014, 1102, 763, 1, 1022, 1101, 0, 31, 1007, 1102, 1, 221, 1024, 1101, 0, 32, 1002, 1102, 1, 29, 1004, 1102, 1, 35, 1016, 1102, 22, 1, 1015, 1101, 0, 27, 1001, 109, 9, 1207, -6, 26, 63, 1005, 63, 199, 4, 187, 1105, 1, 203, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 19, 2105, 1, -4, 4, 209, 1001, 64, 1, 64, 1106, 0, 221, 1002, 64, 2, 64, 109, -33, 1207, 5, 37, 63, 1005, 63, 241, 1001, 64, 1, 64, 1106, 0, 243, 4, 227, 1002, 64, 2, 64, 109, 16, 2102, 1, -2, 63, 1008, 63, 23, 63, 1005, 63, 269, 4, 249, 1001, 64, 1, 64, 1106, 0, 269, 1002, 64, 2, 64, 109, 16, 2106, 0, 0, 1106, 0, 287, 4, 275, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -11, 21101, 40, 0, 0, 1008, 1016, 38, 63, 1005, 63, 311, 1001, 64, 1, 64, 1105, 1, 313, 4, 293, 1002, 64, 2, 64, 109, 4, 21107, 41, 40, -9, 1005, 1011, 329, 1105, 1, 335, 4, 319, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -14, 21108, 42, 42, 5, 1005, 1011, 353, 4, 341, 1106, 0, 357, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 2, 2107, 33, 0, 63, 1005, 63, 379, 4, 363, 1001, 64, 1, 64, 1105, 1, 379, 1002, 64, 2, 64, 109, -7, 1201, 2, 0, 63, 1008, 63, 25, 63, 1005, 63, 401, 4, 385, 1105, 1, 405, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 11, 1201, -8, 0, 63, 1008, 63, 28, 63, 1005, 63, 429, 1001, 64, 1, 64, 1106, 0, 431, 4, 411, 1002, 64, 2, 64, 109, -7, 2108, 26, 1, 63, 1005, 63, 449, 4, 437, 1105, 1, 453, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 9, 1206, 7, 465, 1105, 1, 471, 4, 459, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 4, 21102, 43, 1, -3, 1008, 1015, 42, 63, 1005, 63, 491, 1106, 0, 497, 4, 477, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 7, 21108, 44, 43, -7, 1005, 1018, 517, 1001, 64, 1, 64, 1105, 1, 519, 4, 503, 1002, 64, 2, 64, 109, -28, 2101, 0, 7, 63, 1008, 63, 29, 63, 1005, 63, 545, 4, 525, 1001, 64, 1, 64, 1105, 1, 545, 1002, 64, 2, 64, 109, 11, 2107, 28, -7, 63, 1005, 63, 561, 1105, 1, 567, 4, 551, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -4, 2101, 0, -1, 63, 1008, 63, 26, 63, 1005, 63, 587, 1105, 1, 593, 4, 573, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 9, 1206, 7, 607, 4, 599, 1105, 1, 611, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -10, 1208, 1, 27, 63, 1005, 63, 627, 1106, 0, 633, 4, 617, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 26, 1205, -9, 649, 1001, 64, 1, 64, 1106, 0, 651, 4, 639, 1002, 64, 2, 64, 109, -20, 1208, 0, 23, 63, 1005, 63, 669, 4, 657, 1105, 1, 673, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -7, 2102, 1, 1, 63, 1008, 63, 28, 63, 1005, 63, 693, 1105, 1, 699, 4, 679, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 18, 21102, 45, 1, -6, 1008, 1014, 45, 63, 1005, 63, 725, 4, 705, 1001, 64, 1, 64, 1106, 0, 725, 1002, 64, 2, 64, 109, -23, 1202, 6, 1, 63, 1008, 63, 25, 63, 1005, 63, 751, 4, 731, 1001, 64, 1, 64, 1106, 0, 751, 1002, 64, 2, 64, 109, 20, 2105, 1, 6, 1106, 0, 769, 4, 757, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -22, 2108, 39, 10, 63, 1005, 63, 789, 1001, 64, 1, 64, 1106, 0, 791, 4, 775, 1002, 64, 2, 64, 109, 3, 1202, 6, 1, 63, 1008, 63, 32, 63, 1005, 63, 815, 1001, 64, 1, 64, 1105, 1, 817, 4, 797, 1002, 64, 2, 64, 109, 23, 21107, 46, 47, -9, 1005, 1012, 835, 4, 823, 1106, 0, 839, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 1, 1205, -1, 853, 4, 845, 1105, 1, 857, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -2, 2106, 0, 8, 4, 863, 1001, 64, 1, 64, 1105, 1, 875, 1002, 64, 2, 64, 109, -8, 21101, 47, 0, -2, 1008, 1010, 47, 63, 1005, 63, 897, 4, 881, 1106, 0, 901, 1001, 64, 1, 64, 4, 64, 99, 21102, 27, 1, 1, 21101, 0, 915, 0, 1105, 1, 922, 21201, 1, 27810, 1, 204, 1, 99, 109, 3, 1207, -2, 3, 63, 1005, 63, 964, 21201, -2, -1, 1, 21102, 1, 942, 0, 1106, 0, 922, 22101, 0, 1, -1, 21201, -2, -3, 1, 21101, 957, 0, 0, 1106, 0, 922, 22201, 1, -1, -2, 1106, 0, 968, 22101, 0, -2, -2, 109, -3, 2106, 0, 0},
			[]int{1},
			[]int{2775723069},
		},
		{
			"BOOST: sensor boost mode - distress signal coordinates",
			[]int{1102, 34463338, 34463338, 63, 1007, 63, 34463338, 63, 1005, 63, 53, 1101, 3, 0, 1000, 109, 988, 209, 12, 9, 1000, 209, 6, 209, 3, 203, 0, 1008, 1000, 1, 63, 1005, 63, 65, 1008, 1000, 2, 63, 1005, 63, 904, 1008, 1000, 0, 63, 1005, 63, 58, 4, 25, 104, 0, 99, 4, 0, 104, 0, 99, 4, 17, 104, 0, 99, 0, 0, 1101, 37, 0, 1005, 1101, 30, 0, 1013, 1102, 1, 33, 1019, 1102, 1, 25, 1003, 1102, 1, 28, 1018, 1101, 26, 0, 1006, 1102, 1, 866, 1029, 1101, 760, 0, 1023, 1102, 39, 1, 1012, 1102, 23, 1, 1009, 1101, 281, 0, 1026, 1102, 1, 20, 1011, 1102, 1, 34, 1008, 1101, 0, 36, 1017, 1101, 38, 0, 1000, 1102, 0, 1, 1020, 1102, 278, 1, 1027, 1101, 21, 0, 1010, 1102, 875, 1, 1028, 1101, 0, 212, 1025, 1102, 1, 1, 1021, 1102, 1, 24, 1014, 1102, 763, 1, 1022, 1101, 0, 31, 1007, 1102, 1, 221, 1024, 1101, 0, 32, 1002, 1102, 1, 29, 1004, 1102, 1, 35, 1016, 1102, 22, 1, 1015, 1101, 0, 27, 1001, 109, 9, 1207, -6, 26, 63, 1005, 63, 199, 4, 187, 1105, 1, 203, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 19, 2105, 1, -4, 4, 209, 1001, 64, 1, 64, 1106, 0, 221, 1002, 64, 2, 64, 109, -33, 1207, 5, 37, 63, 1005, 63, 241, 1001, 64, 1, 64, 1106, 0, 243, 4, 227, 1002, 64, 2, 64, 109, 16, 2102, 1, -2, 63, 1008, 63, 23, 63, 1005, 63, 269, 4, 249, 1001, 64, 1, 64, 1106, 0, 269, 1002, 64, 2, 64, 109, 16, 2106, 0, 0, 1106, 0, 287, 4, 275, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -11, 21101, 40, 0, 0, 1008, 1016, 38, 63, 1005, 63, 311, 1001, 64, 1, 64, 1105, 1, 313, 4, 293, 1002, 64, 2, 64, 109, 4, 21107, 41, 40, -9, 1005, 1011, 329, 1105, 1, 335, 4, 319, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -14, 21108, 42, 42, 5, 1005, 1011, 353, 4, 341, 1106, 0, 357, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 2, 2107, 33, 0, 63, 1005, 63, 379, 4, 363, 1001, 64, 1, 64, 1105, 1, 379, 1002, 64, 2, 64, 109, -7, 1201, 2, 0, 63, 1008, 63, 25, 63, 1005, 63, 401, 4, 385, 1105, 1, 405, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 11, 1201, -8, 0, 63, 1008, 63, 28, 63, 1005, 63, 429, 1001, 64, 1, 64, 1106, 0, 431, 4, 411, 1002, 64, 2, 64, 109, -7, 2108, 26, 1, 63, 1005, 63, 449, 4, 437, 1105, 1, 453, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 9, 1206, 7, 465, 1105, 1, 471, 4, 459, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 4, 21102, 43, 1, -3, 1008, 1015, 42, 63, 1005, 63, 491, 1106, 0, 497, 4, 477, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 7, 21108, 44, 43, -7, 1005, 1018, 517, 1001, 64, 1, 64, 1105, 1, 519, 4, 503, 1002, 64, 2, 64, 109, -28, 2101, 0, 7, 63, 1008, 63, 29, 63, 1005, 63, 545, 4, 525, 1001, 64, 1, 64, 1105, 1, 545, 1002, 64, 2, 64, 109, 11, 2107, 28, -7, 63, 1005, 63, 561, 1105, 1, 567, 4, 551, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -4, 2101, 0, -1, 63, 1008, 63, 26, 63, 1005, 63, 587, 1105, 1, 593, 4, 573, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 9, 1206, 7, 607, 4, 599, 1105, 1, 611, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -10, 1208, 1, 27, 63, 1005, 63, 627, 1106, 0, 633, 4, 617, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 26, 1205, -9, 649, 1001, 64, 1, 64, 1106, 0, 651, 4, 639, 1002, 64, 2, 64, 109, -20, 1208, 0, 23, 63, 1005, 63, 669, 4, 657, 1105, 1, 673, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -7, 2102, 1, 1, 63, 1008, 63, 28, 63, 1005, 63, 693, 1105, 1, 699, 4, 679, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 18, 21102, 45, 1, -6, 1008, 1014, 45, 63, 1005, 63, 725, 4, 705, 1001, 64, 1, 64, 1106, 0, 725, 1002, 64, 2, 64, 109, -23, 1202, 6, 1, 63, 1008, 63, 25, 63, 1005, 63, 751, 4, 731, 1001, 64, 1, 64, 1106, 0, 751, 1002, 64, 2, 64, 109, 20, 2105, 1, 6, 1106, 0, 769, 4, 757, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -22, 2108, 39, 10, 63, 1005, 63, 789, 1001, 64, 1, 64, 1106, 0, 791, 4, 775, 1002, 64, 2, 64, 109, 3, 1202, 6, 1, 63, 1008, 63, 32, 63, 1005, 63, 815, 1001, 64, 1, 64, 1105, 1, 817, 4, 797, 1002, 64, 2, 64, 109, 23, 21107, 46, 47, -9, 1005, 1012, 835, 4, 823, 1106, 0, 839, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 1, 1205, -1, 853, 4, 845, 1105, 1, 857, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -2, 2106, 0, 8, 4, 863, 1001, 64, 1, 64, 1105, 1, 875, 1002, 64, 2, 64, 109, -8, 21101, 47, 0, -2, 1008, 1010, 47, 63, 1005, 63, 897, 4, 881, 1106, 0, 901, 1001, 64, 1, 64, 4, 64, 99, 21102, 27, 1, 1, 21101, 0, 915, 0, 1105, 1, 922, 21201, 1, 27810, 1, 204, 1, 99, 109, 3, 1207, -2, 3, 63, 1005, 63, 964, 21201, -2, -1, 1, 21102, 1, 942, 0, 1106, 0, 922, 22101, 0, 1, -1, 21201, -2, -3, 1, 21101, 957, 0, 0, 1106, 0, 922, 22201, 1, -1, -2, 1106, 0, 968, 22101, 0, -2, -2, 109, -3, 2106, 0, 0},
			[]int{2},
			[]int{49115},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			c := New()
			p := NewProgram(tt.program)
			p.Output = make(chan int, 100)
			for _, value := range tt.input {
				p.Input <- value
			}

			c.executeProgram(p)
			var actualOutput []int
			for value := range p.Output {
				actualOutput = append(actualOutput, value)
			}
			if !reflect.DeepEqual(actualOutput, tt.wantOutput) {
				t.Errorf("Run() = %v, want %v", actualOutput, tt.wantOutput)
			}
		})
	}
}