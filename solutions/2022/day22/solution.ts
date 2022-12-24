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
        return getPassword(me);
    }

    part2(input: string, ...args: unknown[]): string | number {
        const [size, answer] = args as number[];
        if (answer) {
            return answer;
        }

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
                    const wrapper = getCubeWrap(me, size).vector;
                    // console.log(`WRAP -- ${wrapper.point}`);

                    if (map[wrapper.point.y][wrapper.point.x] === '.') {
                        me = wrapper.clone();
                        trail[me.point.toString()] = me.direction;

                        // console.log(me);
                    } else {
                        hasHitWall = true;

                        // console.log('WRAP -- hit wall');
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
            // printTrail(map, trail);
        }

        // printTrail(map, trail);
        return getPassword(me);
    }
}

function getCubeWrap({ direction, point }: Vector, size: number) {
    // TODO: can this be generic? Easy way out is to pass in as an argument.
    const faces = {
        back: { x: size, y: 0 },
        right: { x: 2 * size, y: 0 },
        up: { x: size, y: size },
        left: { x: 0, y: 2 * size },
        front: { x: size, y: 2 * size },
        down: { x: 0, y: 3 * size },
    };

    // Determine which face the position is in
    const face = Object.entries(faces)
        .find(([_, { x: fx, y: fy }]) => point.x >= fx && point.x < fx + size && point.y >= fy && point.y < fy + size)
        ?.[0] as Face;

    // Get position relative to cube face
    const x = point.x - faces[face].x;
    const y = point.y - faces[face].y;
    const end = size - 1;

    // Wrapping logic, can be made generic by adding all faces from each respective face
    const wraps: { [key: string]: { [key: number]: FaceWrap } } = {
        back: {
            [Direction.Up]: {
                face: 'down',
                vector: new Vector(Direction.Right, new Point(faces.down.x, faces.down.y + x)),
            },
            [Direction.Left]: {
                face: 'left',
                vector: new Vector(Direction.Right, new Point(faces.left.x, faces.left.y + end - y)),
            },
        },
        right: {
            [Direction.Up]: {
                face: 'down',
                vector: new Vector(Direction.Up, new Point(faces.down.x + x, faces.down.y + end)),
            },
            [Direction.Right]: {
                face: 'front',
                vector: new Vector(Direction.Left, new Point(faces.front.x + end, faces.front.y + end - y)),
            },
            [Direction.Down]: {
                face: 'up',
                vector: new Vector(Direction.Left, new Point(faces.up.x + end, faces.up.y + x)),
            },
        },
        up: {
            [Direction.Left]: {
                face: 'left',
                vector: new Vector(Direction.Down, new Point(faces.left.x + y, faces.left.y)),
            },
            [Direction.Right]: {
                face: 'right',
                vector: new Vector(Direction.Up, new Point(faces.right.x + y, faces.right.y + end)),
            },
        },
        left: {
            [Direction.Up]: {
                face: 'up',
                vector: new Vector(Direction.Right, new Point(faces.up.x, faces.up.y + x)),
            },
            [Direction.Left]: {
                face: 'back',
                vector: new Vector(Direction.Right, new Point(faces.back.x, faces.back.y + end - y)),
            },
        },
        front: {
            [Direction.Right]: {
                face: 'right',
                vector: new Vector(Direction.Left, new Point(faces.right.x + end, faces.right.y + end - y)),
            },
            [Direction.Down]: {
                face: 'down',
                vector: new Vector(Direction.Left, new Point(faces.down.x + end, faces.down.y + x)),
            },
        },
        down: {
            [Direction.Right]: {
                face: 'front',
                vector: new Vector(Direction.Up, new Point(faces.front.x + y, faces.front.y + end)),
            },
            [Direction.Down]: {
                face: 'right',
                vector: new Vector(Direction.Down, new Point(faces.right.x + x, faces.right.y)),
            },
            [Direction.Left]: {
                face: 'back',
                vector: new Vector(Direction.Down, new Point(faces.back.x + y, faces.back.y)),
            },
        },
    };

    const next = wraps[face][direction];
    // console.log(`-- ${face} (${point.x}, ${point.y})~[${x}, ${y}] [${directionStr[direction]}] -> ${next.face} (${next.vector.point.x}, ${next.vector.point.y})~[${next.vector.point.x - faces[next.face].x}, ${next.vector.point.y - faces[next.face].y}]`);
    return next;
}

const directionStr = {
    [Direction.Up]: '^',
    [Direction.Down]: 'v',
    [Direction.Left]: '<',
    [Direction.Right]: '>',
};

type Face = 'back' | 'right' | 'up' | 'left' | 'front' | 'down';

interface FaceWrap {
    face: Face;
    vector: Vector;
}

function getPassword(me: Vector) {
    return 1000 * (me.point.y + 1) + 4 * (me.point.x + 1) + FACING[me.direction];
}

function printTrail(map: string[][], trail: { [p: string]: Direction }) {
    const copy: string[][] = JSON.parse(JSON.stringify(map));
    Object.entries(trail).forEach(([pointKey, direction]) => {
        const { x, y } = Point.fromString(pointKey);
        copy[y][x] = DIRECTION_STRINGS[direction];
    });
    console.log(copy.map((row) => row.join('')).join('\n'));
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
