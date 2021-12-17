package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_part2(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{
			input: []string{"C200B40A82"},
			want:  3,
		},
		{
			input: []string{"04005AC33890"},
			want:  54,
		},
		{
			input: []string{"880086C3E88112"},
			want:  7,
		},
		{
			input: []string{"CE00C43D881120"},
			want:  9,
		},
		{
			input: []string{"D8005AC2A8F0"},
			want:  1,
		},
		{
			input: []string{"F600BC2D8F"},
			want:  0,
		},
		{
			input: []string{"9C005AC2F8F0"},
			want:  0,
		},
		{
			input: []string{"9C0141080250320F1802104A08"},
			want:  1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := part2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parsePacket(t *testing.T) {
	tests := []struct {
		hex  string
		want packet
	}{
		{
			hex: "D2FE28",
			want: packet{
				Binary:       "110100101111111000101",
				Version:      6,
				TypeID:       4,
				LiteralValue: 2021,
			},
		},
		{
			hex: "38006F45291200",
			want: packet{
				Binary:   "0011100000000000011011110100010100101001000100100",
				Version:  1,
				TypeID:   6,
				LengthID: 0,
				SubPackets: []packet{{
					Binary:       "11010001010",
					Version:      6,
					TypeID:       4,
					LiteralValue: 10,
				}, {
					Binary:       "0101001000100100",
					Version:      2,
					TypeID:       4,
					LiteralValue: 20,
				}},
			},
		},
		{
			hex: "EE00D40C823060",
			want: packet{
				Binary:   "111011100000000011010100000011001000001000110000011",
				Version:  7,
				TypeID:   3,
				LengthID: 1,
				SubPackets: []packet{{
					Binary:       "01010000001",
					Version:      2,
					TypeID:       4,
					LiteralValue: 1,
				}, {
					Binary:       "10010000010",
					Version:      4,
					TypeID:       4,
					LiteralValue: 2,
				}, {
					Binary:       "00110000011",
					Version:      1,
					TypeID:       4,
					LiteralValue: 3,
				}},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := parsePacket(tt.hex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calc(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "2 + 3 = 5",
			data: "1100001000000000101101000001001010000011",
			//     VVVTTTILLLLLLLLLLLVVVTTT.^^^^VVVTTT.^^^^
			want: 5,
		},
		{
			name: "2 + . = 2",
			data: "11000010000000000111010000010",
			//     VVVTTTILLLLLLLLLLLVVVTTT.^^^^
			want: 2,
		},
		{
			name: "2 * 3 = 6",
			data: "1100011000000000101101000001001010000011",
			//     VVVTTTILLLLLLLLLLLVVVTTT.^^^^VVVTTT.^^^^
			want: 6,
		},
		{
			name: "2 * . = 2",
			data: "11000110000000000111010000010",
			//     VVVTTTILLLLLLLLLLLVVVTTT.^^^^
			want: 2,
		},
		{
			name: "min(2, 5, 3) = 5",
			data: "110010100000000011110100000100101000001101010000101",
			//     VVVTTTILLLLLLLLLLLVVVTTT.^^^^VVVTTT.^^^^VVVTTT.^^^^
			want: 2,
		},
		{
			name: "max(2, 5, 3) = 5",
			data: "110011100000000011110100000100101000001101010000101",
			//     VVVTTTILLLLLLLLLLLVVVTTT.^^^^VVVTTT.^^^^VVVTTT.^^^^
			want: 5,
		},
		{
			name: "2 > 3 = false",
			data: "1101011000000000101101000001001010000011",
			//     VVVTTTILLLLLLLLLLLVVVTTT.^^^^VVVTTT.^^^^
			want: 0,
		},
		{
			name: "3 > 2 = true",
			data: "1101011000000000101101000001101010000010",
			//     VVVTTTILLLLLLLLLLLVVVTTT.^^^^VVVTTT.^^^^
			want: 1,
		},
		{
			name: "2 < 3 = true",
			data: "1101101000000000101101000001001010000011",
			//     VVVTTTILLLLLLLLLLLVVVTTT.^^^^VVVTTT.^^^^
			want: 1,
		},
		{
			name: "3 < 2 = false",
			data: "1101101000000000101101000001101010000010",
			//     VVVTTTILLLLLLLLLLLVVVTTT.^^^^VVVTTT.^^^^
			want: 0,
		},
		{
			name: "2 == 2 = true",
			data: "1101111000000000101101000001001010000010",
			//     VVVTTTILLLLLLLLLLLVVVTTT.^^^^VVVTTT.^^^^
			want: 1,
		},
		{
			name: "3 == 2 = false",
			data: "1101111000000000101101000001101010000010",
			//     VVVTTTILLLLLLLLLLLVVVTTT.^^^^VVVTTT.^^^^
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calc(tt.data); got != tt.want {
				t.Errorf("calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
