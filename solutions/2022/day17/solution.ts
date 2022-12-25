import { AbstractSolution } from '~/solution';
import { Direction, gridToString, Point } from '~/';

const level: string = process.env.NODE_ENV === 'test' ? 'info' : 'info';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 17;
    filename = 'input.txt';

    part1(jets: string, ...args: unknown[]): string | number {
        return simulateRocksFalling(jets, 2022);
    }

    part2(jets: string, ...args: unknown[]): string | number {
        return simulateRocksFalling(jets, 1000000000000);
    }
}

function simulateRocksFalling(jets: string, maxRocks: number) {
    const chamber = initChamber();
    let pattern = 0;
    let jet: Jet = { direction: jets[0] as JetDirection, index: 0 };
    let tall = 0;

    const seen: { [key: string]: { prevRockNum: number, prevTall: number } } = {};
    let jump = 0;

    for (let rockNum = 1; rockNum <= maxRocks; rockNum++) {
        const rock: Rock = {
            id: pattern,
            pattern: patterns[pattern],
            point: new Point(
                appear.x + leftWall,
                appear.y + tall - patterns[pattern].length
            ),
        };
        let resting = false;

        // debug(`Rock ${rockNum} begins falling`);
        while (!resting) {
            debug(chamberToString(chamber, rock));

            if (jet.direction === '>' && rock.point.x + rock.pattern[0].length !== rightWall && !willCollide(chamber, rock, Direction.Right)) {
                // debug('Jet of gas pushes rock right >');
                rock.point = rock.point.right();
            } else if (jet.direction === '<' && rock.point.x - 1 !== leftWall && !willCollide(chamber, rock, Direction.Left)) {
                // debug('Jet of gas pushes rock left <');
                rock.point = rock.point.left();
            } else {
                // debug(`Jet of gas pushes rock ${jet.direction}, but nothing happens`);
            }
            jet = getNextJet(jets, jet);

            debug(chamberToString(chamber, rock));

            if (rock.point.y + rock.pattern.length === bottom || willCollide(chamber, rock, Direction.Down)) {
                addRock(chamber, rock);
                resting = true;
                tall = Math.min(tall, rock.point.y);

                const top = fingerprintTopRows(chamber, tall);
                const key = `${top}-${rock.id}-${jet.index}`;
                if (seen[key]) {
                    const { prevRockNum, prevTall } = seen[key];
                    const repeat = Math.floor((maxRocks - rockNum) / (rockNum - prevRockNum));
                    rockNum += (rockNum - prevRockNum) * repeat;
                    jump += repeat * (-tall + prevTall);
                }
                seen[key] = { prevRockNum: rockNum, prevTall: tall };

                // debug('Rock falls 1 unit, causing it to come to rest');
                debug(chamberToString(chamber, rock));
            } else {
                rock.point = rock.point.down();

                // debug('Rock falls 1 unit');
            }
        }

        pattern = (pattern + 1) % patterns.length;
    }

    return -tall + jump;
}

function fingerprintTopRows(chamber: { [p: string]: string }, maxY: number) {
    let result = '';
    for (let y = maxY; y < maxY + 7; y++) {
        for (let x = leftWall + 1; x < rightWall; x++) {
            result += chamber[`${x},${y}`] ?? '.';
        }
    }
    return result;
}

function addRock(chamber: { [p: string]: string }, rock: Rock) {
    for (let y = 0; y < rock.pattern.length; y++) {
        for (let x = 0; x < rock.pattern[y].length; x++) {
            const point = new Point(rock.point.x + x, rock.point.y + y);
            if (rock.pattern[y][x] === '@') {
                chamber[point.toString()] = '#';
            }
        }
    }
}

function initChamber() {
    const chamber: { [key: string]: string } = {};
    for (let y = -4; y <= 0; y++) {
        for (let x = 0; x <= 8; x++) {
            const p = `${x},${y}`;
            let char = '.';
            if (x === leftWall || x === rightWall) {
                char = '|';
                if (y === bottom) {
                    char = '+';
                }
            }
            chamber[p] = char;
        }
    }
    return chamber;
}

function willCollide(chamber: { [p: string]: string }, rock: Rock, direction: Direction): boolean {
    const { x: rx, y: ry } = rock.point.move(direction);

    for (let y = 0; y < rock.pattern.length; y++) {
        for (let x = 0; x < rock.pattern[y].length; x++) {
            if (rock.pattern[y][x] !== '@') {
                continue;
            }

            // cx,cy represent the position of a rock chunk
            const cx = x + rx;
            const cy = y + ry;
            const key = `${cx},${cy}`;

            if (chamber[key] === '#') {
                return true;
            }
        }
    }

    return false;
}

const leftWall = 0;
const rightWall = 8;
const bottom = 0;

type JetDirection = '<' | '>';

interface Jet {
    direction: JetDirection;
    index: number;
}

function getNextJet(jets: string, jet: Jet): Jet {
    const index = (jet.index + 1) % jets.length;
    const direction = jets[index] as JetDirection;
    return { direction, index };
}

interface Rock {
    pattern: string[];
    id: number;
    point: Point;
}

const appear = { x: 3, y: -3 };
const patterns = [
    [
        '@@@@',
    ],
    [
        '.@.',
        '@@@',
        '.@.',
    ],
    [
        '..@',
        '..@',
        '@@@',
    ],
    [
        '@',
        '@',
        '@',
        '@',
    ],
    [
        '@@',
        '@@',
    ],
];

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
            const { x, y } = Point.fromString(coordinates);
            minRow = Math.min(minRow, y);
            minCol = Math.min(minCol, x);
            maxRow = Math.max(maxRow, y);
            maxCol = Math.max(maxCol, x);
        });

        const width = Math.abs(minCol) + Math.abs(maxCol) + 1;
        let height = Math.abs(minRow) + Math.abs(maxRow) + 1;

        // Account for top of chamber. Height of the dropping rock + the tallest rock
        height += 7;
        minRow -= 7;

        const grid: string[][] = [];
        for (let row = 0; row < height; row++) {
            grid.push([]);
            const rowDiff = minRow + row;

            for (let col = 0; col < width; col++) {
                const colDiff = minCol + col;
                const spot = map[`${colDiff},${rowDiff}`];

                if (col === 0 || col === width - 1) {
                    grid[row][col] = '|';
                } else if (row === height - 1) {
                    grid[row][col] = '_';
                } else if (spot === '#') {
                    grid[row][col] = '#';
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

    imprint(point: Point, char: string) {
        const y = point.y - this.minRow;
        const x = point.x - this.minCol;
        this.grid[y][x] = char;
    }

    toString() {
        return gridToString(this.grid);
    }
}

function chamberToString(chamber: { [p: string]: string }, rock: Rock) {
    if (level !== 'debug') {
        return '';
    }

    const map = GridMap.fromMap(chamber);

    for (let y = 0; y < rock.pattern.length; y++) {
        for (let x = 0; x < rock.pattern[y].length; x++) {
            const point = new Point(rock.point.x + x, rock.point.y + y);
            const existing = chamber[point.toString()];
            if (rock.pattern[y][x] === '@' && (existing === undefined || existing === '.')) {
                map.imprint(point, '@');
            }
        }
    }

    return map.toString() + '\n';
}

function debug(message: string) {
    if (level === 'debug') {
        // console.clear();
        console.log(message);
        sleep(200);
    }
}

function sleep(millis: number) {
    const waitTill = new Date(new Date().getTime() + millis);
    while (waitTill > new Date()) {
    }
}
