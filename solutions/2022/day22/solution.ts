import { AbstractSolution } from '~/solution';
import { isNum, toInt } from '~/math';
import { Direction, DIRECTION_STRINGS, Point, Vector } from '~/';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 22;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        let { map, path, me } = parseState(input);
        const trail: { [key: string]: Direction } = { [me.point.toString()]: me.direction };

        // console.log(me);

        for (let i = 0; i < path.length; i++) {
            const instruction = path[i];

            // console.log(`Instruction: ${instruction}`);

            if (typeof instruction === 'string') {
                me.rotate(getDegrees(instruction));
                trail[me.point.toString()] = me.direction;
                continue;
            }

            let hasHitWall = false;
            for (let move = 0; move < instruction && !hasHitWall; move++) {
                const next = me.point.move(me.direction);
                const { x, y } = next;

                if (isOutOfBounds(map, next)) {
                    let hasWrapped = false;
                    const wrapper = me.clone();
                    wrapper.rotate(180);
                    while (!hasWrapped) {
                        wrapper.move();
                        if (isOutOfBounds(map, wrapper.point)) {
                            hasWrapped = true;
                            wrapper.rotate(180);
                            wrapper.move();

                            // console.log(`WRAP -- ${wrapper.point}`);

                            if (map[wrapper.point.y][wrapper.point.x] === '.') {
                                me = wrapper.clone();
                                trail[me.point.toString()] = me.direction;

                                // console.log(me);
                            } else {
                                hasHitWall = true;

                                // console.log('WRAP -- hit wall');
                            }
                        }
                    }
                } else if (map[y][x] === '.') {
                    me.move();
                    trail[me.point.toString()] = me.direction;

                    // console.log(me);
                } else if (map[y][x] === '#') {
                    hasHitWall = true;

                    // console.log('hit wall');
                }
            }
        }

        // printTrail(map, trail);
        return 1000 * (me.point.y + 1) + 4 * (me.point.x + 1) + FACING[me.direction];
    }

    part2(input: string, ...args: unknown[]): string | number {
        const lines = input.split('\n');
        return 'TBD';
    }
}

function printTrail(map: string[][], trail: { [p: string]: Direction }) {
    const copy: string[][] = JSON.parse(JSON.stringify(map));
    Object.entries(trail).forEach(([pointKey, direction]) => {
        const { x, y } = Point.fromString(pointKey);
        copy[y][x] = DIRECTION_STRINGS[direction];
    });
    // console.log(copy.map((row) => row.join('')).join('\n'));
}

const FACING = {
    [Direction.Right]: 0,
    [Direction.Down]: 1,
    [Direction.Left]: 2,
    [Direction.Up]: 3,
};

function getDegrees(direction: string) {
    if (direction === 'L') {
        return -90;
    }
    return 90;
}

function isOutOfBounds(map: string[][], { x, y }: Point) {
    return y < 0 || y >= map.length || x < 0 || x >= map[y].length || map[y][x] === ' ';
}

function parseState(input: string) {
    const [mapStr, pathStr] = input.split('\n\n');

    const me = new Vector(Direction.Right, new Point());

    const map: string[][] = [];
    let isStartSet = false;
    const lines = mapStr.split('\n');
    for (let y = 0; y < lines.length; y++) {
        map.push([]);
        for (let x = 0; x < lines[y].length; x++) {
            const char = lines[y][x];
            if (y === 0 && char === '.' && !isStartSet) {
                isStartSet = true;
                me.point.x = x;
                me.point.y = y;
            }
            map[y][x] = char;
        }
    }

    const path: Array<string | number> = [];
    let currentNum = '';
    for (let i = 0; i < pathStr.length; i++) {
        const char = pathStr[i];
        if (isNum(char)) {
            currentNum += char;
        } else {
            if (currentNum) {
                path.push(toInt(currentNum));
                currentNum = '';
            }
            path.push(char);
        }
    }
    if (currentNum) {
        path.push(toInt(currentNum));
    }

    return { map, path, me };
}
