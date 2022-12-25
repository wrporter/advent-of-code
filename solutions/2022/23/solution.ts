import { AbstractSolution } from '~/solution';
import { gridToString } from '~/';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 23;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        const map = parseElves(input);
        const { sumEmptyTiles } = moveElves(map, 10);
        return sumEmptyTiles;
    }

    part2(input: string, ...args: unknown[]): string | number {
        const map = parseElves(input);
        const { round } = moveElves(map, 10_000);
        return round;
    }
}

function moveElves(map: Dictionary<boolean>, maxRounds: number) {
    // console.log('== Initial State ==');
    // console.log(GridMap.fromMap(map).toString());
    // console.log();

    let firstDirection = 0;
    let anyElfHasMoved = false;
    let round = 1;

    for (; round <= maxRounds; round++) {
        const firstHalf: Dictionary<{ from: Point, to: Point }[]> = {};

        for (const elf in map) {
            const { x, y } = fromKey(elf);
            const isAdjacentElf = Object.values(directions)
                .some(({ dx, dy }) => map[toKey({ x: x + dx, y: y + dy })]);
            let hasMoved = false;

            if (isAdjacentElf) {
                for (let group = firstDirection; group < firstDirection + directionGroups.length && !hasMoved; group++) {
                    const directionGroup = directionGroups[group % directionGroups.length];
                    const noElfInDirection = directionGroup.every(({ dx, dy }) => !map[toKey({ x: x + dx, y: y + dy })]);

                    if (noElfInDirection) {
                        hasMoved = true;
                        anyElfHasMoved = true;

                        // The first direction in a group is where we want to move to
                        const { dx, dy } = directionGroup[0];
                        const to = { x: x + dx, y: y + dy };
                        const key = toKey(to);
                        const entry = { from: { x, y }, to };

                        if (firstHalf[key]) {
                            firstHalf[key].push(entry);
                        } else {
                            firstHalf[key] = [entry];
                        }
                    }
                }
            }

            if (!hasMoved) {
                // The elf should not move, keep it where it is.
                const from = fromKey(elf);
                firstHalf[elf] = [{ from, to: from }];
            }
        }

        const secondHalf: Dictionary<boolean> = {};
        for (const spot in firstHalf) {
            const elves = firstHalf[spot];
            if (elves.length === 1) {
                // There is only 1 elf that wants to move to this location, so it's okay.
                secondHalf[toKey(elves[0].to)] = true;
            } else {
                // The elves either couldn't move or there were too many of them. Keep them where they were.
                elves.forEach((elf) => {
                    secondHalf[toKey(elf.from)] = true;
                });
            }
        }

        // Update the current view of the map and the first direction.
        map = secondHalf;
        firstDirection = (firstDirection + 1) % directionGroups.length;

        if (!anyElfHasMoved) {
            break;
        }
        anyElfHasMoved = false;

        // console.log(`== End of Round ${round} ==`);
        // console.log(GridMap.fromMap(map).toString());
        // console.log();
    }

    const rectangle = GridMap.fromMap(map).grid;
    let sumEmptyTiles = 0;
    for (let y = 0; y < rectangle.length; y++) {
        for (let x = 0; x < rectangle[y].length; x++) {
            if (rectangle[y][x] === '.') {
                sumEmptyTiles += 1;
            }
        }
    }

    return { sumEmptyTiles, round };
}

function parseElves(input: string) {
    return input.split('\n').reduce((result, line, y) => {
        for (let x = 0; x < line.length; x++) {
            if (line[x] === '#') {
                result[toKey({ x, y })] = true;
            }
        }
        return result;
    }, {} as Dictionary<boolean>);
}

const directions = {
    North: { name: 'North', dx: 0, dy: -1 },
    NorthEast: { name: 'NorthEast', dx: 1, dy: -1 },
    NorthWest: { name: 'NorthWest', dx: -1, dy: -1 },
    South: { name: 'South', dx: 0, dy: 1 },
    SouthEast: { name: 'SouthEast', dx: 1, dy: 1 },
    SouthWest: { name: 'SouthWest', dx: -1, dy: 1 },
    West: { name: 'West', dx: -1, dy: 0 },
    East: { name: 'East', dx: 1, dy: 0 },
};

const directionGroups = [
    [directions.North, directions.NorthEast, directions.NorthWest],
    [directions.South, directions.SouthEast, directions.SouthWest],
    [directions.West, directions.NorthWest, directions.SouthWest],
    [directions.East, directions.NorthEast, directions.SouthEast],
];

interface Point {
    x: number,
    y: number;
}

function toKey({ x, y }: Point) {
    return `${x},${y}`;
}

function fromKey(coordinates: string): Point {
    const [x, y] = coordinates.split(',')
        .map((v) => Number.parseInt(v, 10));
    return { x, y };
}

type Dictionary<T> = { [key: string]: T };

class GridMap {
    public grid: string[][] = [];
    private minY: number = 0;
    private minX: number = 0;

    static fromMap(map: Dictionary<boolean>) {
        let minY = Number.MAX_SAFE_INTEGER;
        let minX = Number.MAX_SAFE_INTEGER;
        let maxY = Number.MIN_SAFE_INTEGER;
        let maxX = Number.MIN_SAFE_INTEGER;

        Object.keys(map).forEach((coordinates) => {
            const { x, y } = fromKey(coordinates);
            minY = Math.min(minY, y);
            minX = Math.min(minX, x);
            maxY = Math.max(maxY, y);
            maxX = Math.max(maxX, x);
        });

        const width = Math.abs(minX - maxX) + 1;
        let height = Math.abs(minY - maxY) + 1;

        const grid: string[][] = [];
        for (let y = 0; y < height; y++) {
            grid.push([]);
            const rowDiff = minY + y;

            for (let x = 0; x < width; x++) {
                const dx = minX + x;

                if (map[`${dx},${rowDiff}`]) {
                    grid[y][x] = '#';
                } else {
                    grid[y][x] = '.';
                }
            }
        }

        const gridMap = new GridMap();
        gridMap.grid = grid;
        gridMap.minY = minY;
        gridMap.minX = minX;
        return gridMap;
    }

    toString() {
        return gridToString(this.grid);
    }
}
