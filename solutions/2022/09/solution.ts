import { AbstractSolution } from '~/solution';
import { gridToString } from '~/';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 9;
    filename = 'input.txt';

    part1(input: string): string | number {
        return this.solve(input, 2);
    }

    part2(input: string): string | number {
        return this.solve(input, 10);
    }

    solve(input: string, knots: number): string | number {
        const motions = input.split('\n');
        const rope = [...new Array(knots)].map(() => new Position());
        const tail = rope[rope.length - 1];
        const visited = { [tail.toString()]: '#' };

        motions.forEach((motion) => {
            const [directionStr, amountStr] = motion.split(' ');
            const direction = MAP_DIRECTION[directionStr];
            const amount = Number.parseInt(amountStr, 10);

            for (let step = 0; step < amount; step++) {
                moveRope(rope, direction);
                visited[tail.toString()] = '#';

                // Debug
                // const ropeMap = ropeToMap(rope);
                // const gridMap = GridMap.fromMap(ropeMap);
                // console.log(`== ${motion} - ${step + 1} ==\n`);
                // console.log(gridMap.toString());
            }
        });

        return Object.keys(visited).length;
    }
}

function ropeToMap(rope: Position[]) {
    const map = { '0,0': 's', [rope[0].toString()]: 'H' };
    for (let i = rope.length - 1; i > 0; i--) {
        map[rope[i].toString()] = `${i}`;
    }
    return map;
}

function moveRope(rope: Position[], direction: Direction) {
    rope[0].move(direction);
    for (let i = 1; i < rope.length && !rope[i - 1].equals(rope[i]); i++) {
        const prev = rope[i - 1];
        const cur = rope[i];

        if (prev.stepDistance(cur) > 1) {
            if (prev.row > cur.row && prev.col > cur.col) {
                cur.move(Direction.DownRight);
            } else if (prev.row > cur.row && prev.col < cur.col) {
                cur.move(Direction.DownLeft);
            } else if (prev.row < cur.row && prev.col > cur.col) {
                cur.move(Direction.UpRight);
            } else if (prev.row < cur.row && prev.col < cur.col) {
                cur.move(Direction.UpLeft);
            } else if (prev.row > cur.row) {
                cur.move(Direction.Down);
            } else if (prev.row < cur.row) {
                cur.move(Direction.Up);
            } else if (prev.col > cur.col) {
                cur.move(Direction.Right);
            } else if (prev.col < cur.col) {
                cur.move(Direction.Left);
            } else {
                cur.move(direction);
            }
        }
    }
}

class GridMap {
    public grid: string[][] = [];
    private minRow: number = 0;
    private minCol: number = 0;

    static fromMap(map: { [key: string]: string }) {
        let minRow = Number.MAX_SAFE_INTEGER;
        let minCol = Number.MAX_SAFE_INTEGER;
        let maxRow = Number.MIN_SAFE_INTEGER;
        let maxCol = Number.MIN_SAFE_INTEGER;

        Object.keys(map).forEach((coordinates) => {
            const { row, col } = Position.fromString(coordinates);
            minRow = Math.min(minRow, row);
            minCol = Math.min(minCol, col);
            maxRow = Math.max(maxRow, row);
            maxCol = Math.max(maxCol, col);
        });

        const width = Math.abs(minCol) + Math.abs(maxCol) + 1;
        const height = Math.abs(minRow) + Math.abs(maxRow) + 1;

        const grid: string[][] = [];
        for (let row = 0; row < height; row++) {
            grid.push([]);
            const rowDiff = minRow + row;

            for (let col = 0; col < width; col++) {
                const colDiff = minCol + col;
                const char = map[`${rowDiff},${colDiff}`];
                if (char) {
                    grid[row][col] = char;
                } else {
                    grid[row].push('.');
                }
            }
        }

        const gridMap = new GridMap();
        gridMap.grid = grid;
        gridMap.minRow = minRow;
        gridMap.minCol = minCol;
        return gridMap;
    }

    toString() {
        return gridToString(this.grid);
    }
}

enum Direction {
    UpLeft = 'UpLeft',
    Up = 'Up',
    UpRight = 'UpRight',
    Right = 'Right',
    DownRight = 'DownRight',
    Down = 'Down',
    DownLeft = 'DownLeft',
    Left = 'Left',
}

const MAP_DIRECTION: { [key: string]: Direction } = {
    U: Direction.Up,
    D: Direction.Down,
    L: Direction.Left,
    R: Direction.Right,
};

const DIRECTION_MODIFIERS = {
    [Direction.UpLeft]: { row: -1, col: -1 },
    [Direction.Up]: { row: -1, col: 0 },
    [Direction.UpRight]: { row: -1, col: 1 },
    [Direction.Right]: { row: 0, col: 1 },
    [Direction.DownRight]: { row: 1, col: 1 },
    [Direction.Down]: { row: 1, col: 0 },
    [Direction.DownLeft]: { row: 1, col: -1 },
    [Direction.Left]: { row: 0, col: -1 },
};

class Position {
    constructor(public row = 0, public col = 0) {}

    move(direction: Direction, amount = 1) {
        const { row, col } = DIRECTION_MODIFIERS[direction];
        this.row += row * amount;
        this.col += col * amount;
    }

    up(amount = 1) {
        this.row -= amount;
    }

    down(amount = 1) {
        this.row += amount;
    }

    left(amount = 1) {
        this.col -= amount;
    }

    right(amount = 1) {
        this.col += amount;
    }

    manhattanDistance({ row, col }: Position) {
        return Math.abs(this.row - row) + Math.abs(this.col - col);
    }

    stepDistance({ row, col }: Position) {
        return Math.max(Math.abs(this.row - row), Math.abs(this.col - col));
    }

    equals({ row, col }: Position) {
        return this.row === row && this.col === col;
    }

    clone() {
        return new Position(this.row, this.col);
    }

    static fromString(coordinates: string) {
        const [row, col] = coordinates.split(',')
            .map((v) => Number.parseInt(v, 10));
        return new Position(row, col);
    }

    toString() {
        return `${this.row},${this.col}`;
    }
}
