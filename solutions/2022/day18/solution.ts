import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 18;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        const { lava } = parseLava(input);
        let surfaceArea = 0;

        Object.entries(lava).forEach(([key, point]) => {
            DIRECTIONS.forEach((direction) => {
                if (!lava[point.move(direction).key()]) {
                    surfaceArea = surfaceArea + 1;
                }
            });
        });

        return surfaceArea;
    }

    part2(input: string, ...args: unknown[]): string | number {
        const { lava, min, max } = parseLava(input);
        let surfaceArea = 0;

        const seen: {[key: string]: boolean} = {};

        Object.entries(lava).forEach(([key, point]) => {
            // if (DIRECTIONS.every((direction) => lava[point.move(direction).toString()])) {
            //     return;
            // }

            const queue: Node[] = [{ point, score: 0 }];

            while (queue.length > 0) {
                let { point, score } = queue.shift() as Node;
                const key = point.key();

                // if (!lava[point.key()]) {
                //     DIRECTIONS.forEach((direction) => {
                //         if ()
                //     })
                // }

                if (seen[key]) {
                    continue;
                }
                seen[key] = true;

                if (
                    point.x === min.x || point.y === min.y || point.z === min.z ||
                    point.x === max.x || point.y === max.y || point.z === max.z
                ) {
                    surfaceArea += score;
                    continue;
                }

                DIRECTIONS.forEach((direction) => {
                    const next = point.move(direction);
                    if (!lava[next.key()] && !seen[next.key()]) {
                        queue.push({
                            point: next,
                            score: score + 1,
                        });
                    }
                    seen[key] = true;
                });
            }
        });

        return surfaceArea;
    }
}

interface Node {
    point: Point3D;
    score: number;
}

function parseLava(input: string) {
    const min = new Point3D();
    const max = new Point3D();

    const lava = input.split('\n')
        .reduce((result, list) => {
            const point = Point3D.fromKey(list);
            result[list] = point;

            min.x = Math.min(min.x, point.x);
            min.y = Math.min(min.y, point.y);
            min.z = Math.min(min.z, point.z);
            max.x = Math.max(max.x, point.x);
            max.y = Math.max(max.y, point.y);
            max.z = Math.max(max.z, point.z);

            return result;
        }, {} as { [key: string]: Point3D });

    return { lava, min, max };
}

export enum Direction {
    Up = 'Up',
    Right = 'Right',
    Down = 'Down',
    Left = 'Left',
    In = 'In',
    Out = 'Out',

    UpRightIn = 'UpRightIn',
    UpLeftIn = 'UpLeftIn',
    UpRightOut = 'UpRightOut',
    UpLeftOut = 'UpLeftOut',
    DownRightIn = 'DownRightIn',
    DownLeftIn = 'DownLeftIn',
    DownRightOut = 'DownRightOut',
    DownLeftOut = 'DownLeftOut',
}

export const DIRECTIONS = [
    Direction.Up,
    Direction.Down,
    Direction.Left,
    Direction.Right,
    Direction.In,
    Direction.Out,
];

export const DIAGONALS = [
    Direction.UpRightIn,
    Direction.UpLeftIn,
    Direction.UpRightOut,
    Direction.UpLeftOut,
    Direction.DownRightIn,
    Direction.DownLeftIn,
    Direction.DownRightOut,
    Direction.DownLeftOut,
];

export const DIRECTION_MODIFIERS = {
    [Direction.Up]: { x: 0, y: -1, z: 0 },
    [Direction.Right]: { x: 1, y: 0, z: 0 },
    [Direction.Down]: { x: 0, y: 1, z: 0 },
    [Direction.Left]: { x: -1, y: 0, z: 0 },
    [Direction.In]: { x: 0, y: 0, z: 1 },
    [Direction.Out]: { x: 0, y: 0, z: -1 },

    [Direction.UpRightIn]: { x: 1, y: 1, z: 1 },
    [Direction.UpLeftIn]: { x: -1, y: 1, z: 1 },
    [Direction.UpRightOut]: { x: 1, y: 1, z: -1 },
    [Direction.UpLeftOut]: { x: -1, y: 1, z: -1 },
    [Direction.DownRightIn]: { x: 1, y: -1, z: 1 },
    [Direction.DownLeftIn]: { x: -1, y: -1, z: 1 },
    [Direction.DownRightOut]: { x: 1, y: -1, z: -1 },
    [Direction.DownLeftOut]: { x: -1, y: -1, z: -1 },
};

class Point3D {
    constructor(
        public x: number = 0,
        public y: number = 0,
        public z: number = 0,
    ) {}

    move(direction: Direction, amount = 1) {
        const { x: dx, y: dy, z: dz } = DIRECTION_MODIFIERS[direction];
        const x = this.x + (dx * amount);
        const y = this.y + (dy * amount);
        const z = this.z + (dz * amount);
        return new Point3D(x, y, z);
    }

    static fromKey(coordinates: string) {
        const [x, y, z] = coordinates.split(',')
            .map((v) => Number.parseInt(v, 10));
        return new Point3D(x, y, z);
    }

    key() {
        return `${this.x},${this.y},${this.z}`;
    }
}
