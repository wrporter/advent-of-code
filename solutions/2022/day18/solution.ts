import { AbstractSolution } from '~/solution';
import { Dictionary, Point } from '~/';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 18;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        const { cubes } = parseLava(input);
        let surfaceArea = 0;

        for (const cube of Object.values(cubes)) {
            for (const side of getSides(cube)) {
                if (!cubes[toKey(side)]) {
                    surfaceArea = surfaceArea + 1;
                }
            }
        }

        return surfaceArea;
    }

    part2(input: string, ...args: unknown[]): string | number {
        const { cubes, min, max } = parseLava(input);
        const visitedAir: Dictionary<boolean> = {};
        const queue = [{ x: min.x, y: min.y, z: min.z }];

        while (queue.length > 0) {
            const point = queue.pop() as Point3D;

            for (const side of getSides(point)) {
                const key = toKey(side);
                if (isInBounds(side, min, max) && !visitedAir[key] && !cubes[key]) {
                    queue.push(side);
                }
            }

            visitedAir[toKey(point)] = true;
        }

        let surfaceArea = 0;
        for (const air of Object.keys(visitedAir)) {
            for (const side of getSides(fromKey(air))) {
                if (cubes[toKey(side)]) {
                    surfaceArea = surfaceArea + 1;
                }
            }
        }
        return surfaceArea;
    }
}

function isInBounds({ x, y, z }: Point3D, min: Point3D, max: Point3D) {
    return x >= min.x && y >= min.y && z >= min.z &&
        x <= max.x && y <= max.y && z <= max.z;
}

function getSides({ x, y, z }: Point3D) {
    return DIRECTIONS.reduce((sides, { dx, dy, dz }) => {
        sides.push({ x: x + dx, y: y + dy, z: z + dz });
        return sides;
    }, [] as Point3D[]);
}

function parseLava(input: string) {
    const min = { x: 0, y: 0, z: 0 };
    const max = { x: 0, y: 0, z: 0 };

    const cubes = input.split('\n')
        .reduce((result, list) => {
            const point = fromKey(list);
            result[list] = point;
            const { x, y, z } = point;

            min.x = Math.min(min.x, point.x - 1);
            min.y = Math.min(min.y, point.y - 1);
            min.z = Math.min(min.z, point.z - 1);
            max.x = Math.max(max.x, point.x + 1);
            max.y = Math.max(max.y, point.y + 1);
            max.z = Math.max(max.z, point.z + 1);

            return result;
        }, {} as { [key: string]: Point3D });

    return { cubes, min, max };
}

export const DIRECTIONS = [
    { dx: 0, dy: -1, dz: 0 },
    { dx: 1, dy: 0, dz: 0 },
    { dx: 0, dy: 1, dz: 0 },
    { dx: -1, dy: 0, dz: 0 },
    { dx: 0, dy: 0, dz: -1 },
    { dx: 0, dy: 0, dz: 1 },
];

interface Point3D {
    x: number;
    y: number;
    z: number;
}

function toKey({ x, y, z }: Point3D): string {
    return `${x},${y},${z}`;
}

function fromKey(coordinates: string): Point3D {
    const [x, y, z] = coordinates.split(',')
        .map((v) => Number.parseInt(v, 10));
    return { x, y, z };
}
