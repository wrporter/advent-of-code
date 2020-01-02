package monitor

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMonitor_FindBestAsteroid(t *testing.T) {
	tests := []struct {
		lines []string
		want  Station
	}{
		{
			[]string{
				".#..#",
				".....",
				"#####",
				"....#",
				"...##",
			},
			Station{
				NumVisibleAsteroids: 8,
				Point:               Point{3, 4},
			},
		},
		{
			[]string{
				"......#.#.",
				"#..#.#....",
				"..#######.",
				".#.#.###..",
				".#..#.....",
				"..#....#.#",
				"#..#....#.",
				".##.#..###",
				"##...#..#.",
				".#....####",
			},
			Station{
				NumVisibleAsteroids: 33,
				Point:               Point{5, 8},
			},
		},
		{
			[]string{
				"#.#...#.#.",
				".###....#.",
				".#....#...",
				"##.#.#.#.#",
				"....#.#.#.",
				".##..###.#",
				"..#...##..",
				"..##....##",
				"......#...",
				".####.###.",
			},
			Station{
				NumVisibleAsteroids: 35,
				Point:               Point{1, 2},
			},
		},
		{
			[]string{
				".#..#..###",
				"####.###.#",
				"....###.#.",
				"..###.##.#",
				"##.##.#.#.",
				"....###..#",
				"..#.#..#.#",
				"#..#.#.###",
				".##...##.#",
				".....#.#..",
			},
			Station{
				NumVisibleAsteroids: 41,
				Point:               Point{6, 3},
			},
		},
		{
			[]string{
				".#..##.###...#######",
				"##.############..##.",
				".#.######.########.#",
				".###.#######.####.#.",
				"#####.##.#.##.###.##",
				"..#####..#.#########",
				"####################",
				"#.####....###.#.#.##",
				"##.#################",
				"#####.##.###..####..",
				"..######..##.#######",
				"####.##.####...##..#",
				".#####..#.######.###",
				"##...#.##########...",
				"#.##########.#######",
				".####.#.###.###.#.##",
				"....##.##.###..#####",
				".#.#.###########.###",
				"#.#.#.#####.####.###",
				"###.##.####.##.#..##",
			},
			Station{
				NumVisibleAsteroids: 210,
				Point:               Point{11, 13},
			},
		},
		{
			[]string{
				"#.....#...#.........###.#........#..",
				"....#......###..#.#.###....#......##",
				"......#..###.......#.#.#.#..#.......",
				"......#......#.#....#.##....##.#.#.#",
				"...###.#.#.......#..#...............",
				"....##...#..#....##....#...#.#......",
				"..##...#.###.....##....#.#..##.##...",
				"..##....#.#......#.#...#.#...#.#....",
				".#.##..##......##..#...#.....##...##",
				".......##.....#.....##..#..#..#.....",
				"..#..#...#......#..##...#.#...#...##",
				"......##.##.#.#.###....#.#..#......#",
				"#..#.#...#.....#...#...####.#..#...#",
				"...##...##.#..#.....####.#....##....",
				".#....###.#...#....#..#......#......",
				".##.#.#...#....##......#.....##...##",
				".....#....###...#.....#....#........",
				"...#...#....##..#.#......#.#.#......",
				".#..###............#.#..#...####.##.",
				".#.###..#.....#......#..###....##..#",
				"#......#.#.#.#.#.#...#.#.#....##....",
				".#.....#.....#...##.#......#.#...#..",
				"...##..###.........##.........#.....",
				"..#.#..#.#...#.....#.....#...###.#..",
				".#..........#.......#....#..........",
				"...##..#..#...#..#...#......####....",
				".#..#...##.##..##..###......#.......",
				".##.....#.......#..#...#..#.......#.",
				"#.#.#..#..##..#..............#....##",
				"..#....##......##.....#...#...##....",
				".##..##..#.#..#.................####",
				"##.......#..#.#..##..#...#..........",
				"#..##...#.##.#.#.........#..#..#....",
				".....#...#...#.#......#....#........",
				"....#......###.#..#......##.....#..#",
				"#..#...##.........#.....##.....#....",
			},
			Station{
				NumVisibleAsteroids: 303,
				Point:               Point{26, 29},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			m := &Monitor{}
			if got, _ := m.FindBestAsteroid(SplitLines(tt.lines)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBestAsteroid() = %v, want %v", got, tt.want)
			}
		})
	}
}

var test1 = []string{
	".#..##.###...#######",
	"##.############..##.",
	".#.######.########.#",
	".###.#######.####.#.",
	"#####.##.#.##.###.##",
	"..#####..#.#########",
	"####################",
	"#.####....###.#.#.##",
	"##.#################",
	"#####.##.###..####..",
	"..######..##.#######",
	"####.##.####...##..#",
	".#####..#.######.###",
	"##...#.##########...",
	"#.##########.#######",
	".####.#.###.###.#.##",
	"....##.##.###..#####",
	".#.#.###########.###",
	"#.#.#.#####.####.###",
	"###.##.####.##.#..##",
}

func TestMonitor_ZapAsteroids(t *testing.T) {
	type args struct {
		lines  []string
		target int
	}
	tests := []struct {
		args  args
		want  int
		want1 *Point
	}{
		{
			args{test1, 1},
			1112,
			&Point{11, 12},
		},
		{
			args{test1, 2},
			1201,
			&Point{12, 1},
		},
		{
			args{test1, 3},
			1202,
			&Point{12, 2},
		},
		{
			args{test1, 10},
			1208,
			&Point{12, 8},
		},
		{
			args{test1, 20},
			1600,
			&Point{16, 0},
		},
		{
			args{test1, 50},
			1609,
			&Point{16, 9},
		},
		{
			args{test1, 100},
			1016,
			&Point{10, 16},
		},
		{
			args{test1, 199},
			906,
			&Point{9, 6},
		},
		{
			args{test1, 200},
			802,
			&Point{8, 2},
		},
		{
			args{test1, 201},
			1009,
			&Point{10, 9},
		},
		{
			args{test1, 299},
			1101,
			&Point{11, 1},
		},
		{
			args{
				[]string{
					"#.....#...#.........###.#........#..",
					"....#......###..#.#.###....#......##",
					"......#..###.......#.#.#.#..#.......",
					"......#......#.#....#.##....##.#.#.#",
					"...###.#.#.......#..#...............",
					"....##...#..#....##....#...#.#......",
					"..##...#.###.....##....#.#..##.##...",
					"..##....#.#......#.#...#.#...#.#....",
					".#.##..##......##..#...#.....##...##",
					".......##.....#.....##..#..#..#.....",
					"..#..#...#......#..##...#.#...#...##",
					"......##.##.#.#.###....#.#..#......#",
					"#..#.#...#.....#...#...####.#..#...#",
					"...##...##.#..#.....####.#....##....",
					".#....###.#...#....#..#......#......",
					".##.#.#...#....##......#.....##...##",
					".....#....###...#.....#....#........",
					"...#...#....##..#.#......#.#.#......",
					".#..###............#.#..#...####.##.",
					".#.###..#.....#......#..###....##..#",
					"#......#.#.#.#.#.#...#.#.#....##....",
					".#.....#.....#...##.#......#.#...#..",
					"...##..###.........##.........#.....",
					"..#.#..#.#...#.....#.....#...###.#..",
					".#..........#.......#....#..........",
					"...##..#..#...#..#...#......####....",
					".#..#...##.##..##..###......#.......",
					".##.....#.......#..#...#..#.......#.",
					"#.#.#..#..##..#..............#....##",
					"..#....##......##.....#...#...##....",
					".##..##..#.#..#.................####",
					"##.......#..#.#..##..#...#..........",
					"#..##...#.##.#.#.........#..#..#....",
					".....#...#...#.#......#....#........",
					"....#......###.#..#......##.....#..#",
					"#..#...##.........#.....##.....#....",
				}, 200},
			408,
			&Point{4, 8},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			m := &Monitor{}
			field := SplitLines(tt.args.lines)
			got, _ := m.ZapAsteroids(field, tt.args.target)
			if got != tt.want {
				t.Errorf("ZapAsteroids() got = %v, want %v", got, tt.want)
			}
		})
	}
}