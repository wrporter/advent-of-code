package scaffold

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day13/public/computer"
	"strings"
	"time"
)

type Scaffold struct {
	code []int
}

type VacuumRobot struct {
	input     chan<- int
	output    <-chan int
	point     Point
	direction Direction
}

type TurnDirection int

const (
	TurnRight TurnDirection = 1
	TurnLeft  TurnDirection = -1
)

func (t TurnDirection) getCommand() string {
	if t == TurnRight {
		return "R"
	}
	return "L"
}

type Direction int

const (
	Up    Direction = 1
	Right Direction = 2
	Down  Direction = 3
	Left  Direction = 4
)

type Point struct {
	X int
	Y int
}

func New(code []int) *Scaffold {
	return &Scaffold{code}
}

func (s *Scaffold) Scan() (grid [][]rune, robot *VacuumRobot) {
	cpu := computer.New()
	program := computer.NewProgram(s.code)
	output := program.Output
	cpu.Run(program)

	return readGrid(output)
}

func (s *Scaffold) DisplayScaffold(grid [][]rune) {
	displayGrid(grid)
}

func (s *Scaffold) SumAlignmentIntersections(grid [][]rune) int {
	intersections := getIntersections(grid)
	sum := 0
	for _, intersection := range intersections {
		sum += getAlignmentParameter(intersection)
	}
	return sum
}

func (s *Scaffold) MoveRobot(grid [][]rune, robot *VacuumRobot) int {
	robot.wakeUp(s.code)
	//commands := buildCommands(grid, robot)
	commands := "A,B,A,C,B,C,A,B,A,C\n" +
		"R,6,L,10,R,8,R,8\n" +
		"R,12,L,8,L,10\n" +
		"R,12,L,10,R,6,L,10\n" +
		"n\n"

	go func() {
		for _, command := range commands {
			robot.input <- int(command)
		}
	}()

	var dustCollected int
	for c := range robot.output {
		dustCollected = c
	}
	return dustCollected
}

func showVideoFeed(robot *VacuumRobot) {
	for c := range robot.output {
		fmt.Print(string(c))
	}
}

var directions = []Direction{Up, Right, Down, Left}
var directionModifiers = []Point{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func (p Point) Add(direction Direction) Point {
	x := p.X + directionModifiers[direction-1].X
	y := p.Y + directionModifiers[direction-1].Y
	return Point{x, y}
}

func (d Direction) GetTurnDirection(dest Direction) TurnDirection {
	if directions[(int(d)+int(TurnRight)-1)%4] == dest {
		return TurnRight
	} else {
		return TurnLeft
	}
}

type node struct {
	point     Point
	direction Direction
}

func buildCommands(grid [][]rune, robot *VacuumRobot) {
	queue := []node{{robot.point, robot.direction}}
	visited := make(map[Point]bool)
	var commands []string
	var next node

	for len(queue) > 0 {
		next, queue = queue[0], queue[1:]

		for _, direction := range directions {
			point := next.point.Add(direction)
			if !visited[point] && isInBounds(grid, point) && grid[point.Y][point.X] == '#' {
				steps := 1
				for nextIsOpen(grid, point, direction) {
					point = point.Add(direction)
					visited[point] = true
					steps++
				}

				visited[next.point] = true
				commands = append(commands, next.direction.GetTurnDirection(direction).getCommand())
				commands = append(commands, fmt.Sprint(steps))
				queue = append(queue, node{point, direction})
				break
			}
		}
	}

	fmt.Println(strings.Join(commands, ","))
}

func nextIsOpen(grid [][]rune, point Point, direction Direction) bool {
	p := point.Add(direction)
	return isInBounds(grid, p) && grid[p.Y][p.X] == '#'
}

func isInBounds(grid [][]rune, p Point) bool {
	return p.Y >= 0 &&
		p.Y <= len(grid) &&
		p.X >= 0 &&
		p.X < len(grid[p.Y])
}

func (r *VacuumRobot) wakeUp(code []int) {
	cpu := computer.New()
	program := computer.NewProgram(code)
	program.Memory[0] = 2
	r.input = program.Input
	r.output = program.Output
	cpu.Run(program)
}

func getAlignmentParameter(intersection Point) int {
	return intersection.X * intersection.Y
}

func getIntersections(grid [][]rune) []Point {
	var intersections []Point
	for y, row := range grid {
		for x := range row {
			if isIntersection(grid, y, x) {
				intersections = append(intersections, Point{x, y})
			}
		}
	}
	return intersections
}

func isScaffold(grid [][]rune, y, x int) bool {
	return y >= 0 &&
		y < len(grid) &&
		x >= 0 &&
		x < len(grid[y]) &&
		grid[y][x] == '#'
}

func isIntersection(grid [][]rune, y, x int) bool {
	return isScaffold(grid, y, x) &&
		isScaffold(grid, y-1, x) &&
		isScaffold(grid, y+1, x) &&
		isScaffold(grid, y, x-1) &&
		isScaffold(grid, y, x+1)
}

func displayGrid(grid [][]rune) {
	out := &strings.Builder{}
	//out.WriteString("=====================================================\n")
	out.WriteString("\033c")
	out.WriteString(renderGrid(grid))
	fmt.Print(out.String())
	time.Sleep(time.Millisecond * 20)
}

func renderGrid(grid [][]rune) string {
	result := ""
	for _, row := range grid {
		for _, spot := range row {
			result += string(spot)
		}
		result += "\n"
	}
	return result
}

func readGrid(output <-chan int) ([][]rune, *VacuumRobot) {
	var robot *VacuumRobot
	grid := make([][]rune, 0)
	row := make([]rune, 0)

	for ascii := range output {
		if ascii == '\n' {
			grid = append(grid, row)
			row = make([]rune, 0)
		} else {
			row = append(row, rune(ascii))

			if isRobot(rune(ascii)) {
				robot = &VacuumRobot{
					input:     nil,
					output:    nil,
					point:     Point{len(row) - 1, len(grid)},
					direction: parseDirection(rune(ascii)),
				}
			}
		}
	}

	return grid, robot
}

func isRobot(r rune) bool {
	return r == '^' || r == '<' || r == '>' || r == 'v'
}

func parseDirection(r rune) Direction {
	switch r {
	case '^':
		return Up
	case '>':
		return Right
	case 'v':
		return Down
	case '<':
		return Left
	default:
		return Up
	}
}
