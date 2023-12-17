import { AbstractSolution } from '~/solution';
import { Point } from '~/';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 14;
    filename = 'input.txt';

    part1(input: string): string | number {
        const { source, scan, bottom } = scanRocks(input);
        const hasEnteredVoid = (unit: Point) => unit.y > bottom;
        const shouldFallTo = (position: Point) => !scan[position.toString()];
        return sumRestedSand(scan, source, shouldFallTo, hasEnteredVoid);
    }

    part2(input: string): string | number {
        let { source, scan, bottom } = scanRocks(input);
        bottom += 2;
        const hasReachedSource = (unit: Point) => unit.equals(source);
        const shouldFallTo = (position: Point) => !scan[position.toString()] && position.y !== bottom;
        return sumRestedSand(scan, source, shouldFallTo, hasReachedSource);
    }
}

function sumRestedSand(
    scan: { [key: string]: boolean },
    source: Point,
    shouldFallTo: (position: Point) => boolean,
    shouldExit: (unit: Point) => boolean
) {
    let exitCondition = false;
    let sumSandComeToRest = 0;
    let next: Point;

    while (!exitCondition) {
        let unit = source.clone();

        let hasComeToRest = false;
        while (!hasComeToRest && !exitCondition) {
            if ((next = unit.down()) && shouldFallTo(next)) {
                unit = next;
            } else if ((next = unit.downLeft()) && shouldFallTo(next)) {
                unit = next;
            } else if ((next = unit.downRight()) && shouldFallTo(next)) {
                unit = next;
            } else {
                scan[unit.toString()] = true;
                sumSandComeToRest++;
                hasComeToRest = true;
            }

            if (shouldExit(unit)) {
                exitCondition = true;
            }
        }
    }

    return sumSandComeToRest;
}

function scanRocks(input: string) {
    const rockPaths = input.split('\n');
    const source = new Point(500, 0);
    const scan: { [key: string]: boolean } = {};
    let bottom = Number.MIN_SAFE_INTEGER;

    rockPaths.forEach((path) => {
        const points = path.split(' -> ').map(Point.fromString);
        for (let i = 1; i < points.length; i++) {
            const a = points[i - 1];
            const b = points[i];

            if (a.x === b.x) {
                for (let y = Math.min(a.y, b.y); y <= Math.max(a.y, b.y); y++) {
                    scan[new Point(a.x, y).toString()] = true;
                }
            } else if (a.y === b.y) {
                for (let x = Math.min(a.x, b.x); x <= Math.max(a.x, b.x); x++) {
                    scan[new Point(x, a.y).toString()] = true;
                }
            }

            bottom = Math.max(bottom, a.y, b.y);
        }
    });
    return { source, scan, bottom };
}
