import { AbstractSolution } from '~/solution';
import { DIRECTIONS, Point } from '~/';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 12;
    filename = 'input.txt';

    part1(input: string): string | number {
        const { map, start, end } = parseInput(input);
        return findShortestPath(map, [start], end);
    }

    part2(input: string): string | number {
        const { map, start, end, aStarts } = parseInput(input);
        return findShortestPath(map, [start, ...aStarts], end);
    }
}

interface Node {
    point: Point;
    steps: number;
}

function findShortestPath(map: string[][], starts: Point[], end: Point) {
    const queue = starts.map((start) => ({ point: start, steps: 0 }));
    const seen: { [key: string]: boolean } = {};
    let current: Node = queue[0];

    while (!current.point.equals(end)) {
        current = queue[0];
        const { point: from, steps } = current;
        queue.shift();

        if (seen[from.toString()]) {
            continue;
        }
        seen[from.toString()] = true;

        for (const direction of DIRECTIONS) {
            const to = from.move(direction);
            const {x, y} = to;
            if (
                y >= 0 && y < map.length &&
                x >= 0 && x < map[y].length &&
                map[y][x].charCodeAt(0) <= map[from.y][from.x].charCodeAt(0) + 1
            ) {
                queue.push({ point: to, steps: steps + 1 });
            }
        }
    }

    return current.steps;
}

const START = 'S';
const END = 'E';

function parseInput(input: string) {
    const lines = input.split('\n');
    const start = new Point();
    const end = new Point();
    const aStarts: Point[] = [];

    const map = lines.map((line, y) => {
        const row: string[] = [];
        for (let x = 0; x < line.length; x++) {
            const char = line[x];
            if (char === START) {
                start.x = x;
                start.y = y;
                row[x] = 'a';
            } else if (char === END) {
                end.x = x;
                end.y = y;
                row[x] = 'z';
            } else {
                if (char === 'a') {
                    aStarts.push(new Point(x, y));
                }
                row[x] = char;
            }
        }
        return row;
    });

    return { map, start, end, aStarts };
}
