import { AbstractSolution } from '~/solution';
import { mod } from '~/math';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 24;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        const valley = parseValley(input);
        const goals = [valley.end];
        return findMinTime(valley, goals);
    }

    part2(input: string, ...args: unknown[]): string | number {
        const valley = parseValley(input);
        const goals = [valley.end, valley.start, valley.end];
        return findMinTime(valley, goals);
    }
}

function findMinTime(valley: Valley, goals: Point[]) {
    const { walls, blizzardStarts, start, height, width } = valley;
    let time = 0;
    let queue = new MySet<Point>([start]);

    while (goals.length > 0) {
        time += 1;

        const blizzards = new MySet<Point>();
        blizzardStarts.forEach(({ x, y, dx, dy }) => {
            const next = { x: 1 + mod(x + time * dx - 1, width - 2), y: 1 + mod(y + time * dy - 1, height - 2) };
            blizzards.add(next);
        });

        const nextPositions = new MySet<Point>();
        queue.forEach((valueStr) => {
            const { x, y } = JSON.parse(valueStr) as Point;
            directions.forEach(({ dx, dy }) => {
                const nx = x + dx;
                const ny = y + dy;
                const next = { x: nx, y: ny };

                if (
                    nx >= 0 && nx < width && ny >= 0 && ny < height &&
                    !walls.has(next) &&
                    !blizzards.has(next)
                ) {
                    nextPositions.add(next);
                }
            });
        });

        queue = nextPositions;

        if (queue.has(goals[0])) {
            // console.log(`Goal ${JSON.stringify(goals[0])} reached after ${time} steps (queue size: ${queue.size})`);
            queue = new MySet<Point>([goals.shift()]);
        }
    }

    return time;
}

function parseValley(input: string): Valley {
    const lines = input.split('\n');

    const walls = new MySet<Point>();
    const blizzardStarts = new Set<Vector>();
    const height = lines.length;
    const width = lines[0].length;
    let start = { x: 0, y: 0 };
    let end = { x: 0, y: 0 };

    lines.forEach((line, y) => {
        for (let x = 0; x < line.length; x++) {
            const char = line[x];
            if (y === 0 && char === '.') {
                start = { x, y };
            }
            if (y === lines.length - 1 && char === '.') {
                end = { x, y };
            }

            if (char === '#') {
                walls.add({ x, y });
            } else if (char === '>') {
                blizzardStarts.add({ x, y, dx: 1, dy: 0 });
            } else if (char === '<') {
                blizzardStarts.add({ x, y, dx: -1, dy: 0 });
            } else if (char === '^') {
                blizzardStarts.add({ x, y, dx: 0, dy: -1 });
            } else if (char === 'v') {
                blizzardStarts.add({ x, y, dx: 0, dy: 1 });
            }
        }
    });

    // console.log(`Maze size: ${height}x${width}, ${walls.size} walls, ${blizzardStarts.size} blizzards`);

    return { walls, blizzardStarts, start, end, height, width };
}

class MySet<T> extends Set {
    add(value: T): this {
        super.add(JSON.stringify(value));
        return this;
    }

    has(value: T): boolean {
        return super.has(JSON.stringify(value));
    }

    delete(value: T): boolean {
        return super.delete(JSON.stringify(value));
    }
}

const directions = [
    { dx: 0, dy: 0 },
    { dx: 1, dy: 0 },
    { dx: -1, dy: 0 },
    { dx: 0, dy: 1 },
    { dx: 0, dy: -1 },
];

interface Valley {
    start: Point;
    end: Point;
    blizzardStarts: Set<Vector>;
    width: number;
    height: number;
    walls: MySet<Point>;
}

interface Point {
    x: number,
    y: number;
}

interface Vector extends Point {
    dx: number,
    dy: number;
}