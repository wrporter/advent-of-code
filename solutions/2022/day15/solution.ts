import { AbstractSolution } from '~/solution';
import { Direction, DIRECTION_MODIFIERS, DIRECTIONS, Point } from '~/';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 15;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        let [targetRow] = args;
        const sensors = parseSensors(input);
        const targetBeacon = sensors.find((sensor) => sensor.beacon.y === targetRow)
            ?.beacon ?? new Point();

        let sum = 0;
        sum += sumReachedBySensorInRow(targetBeacon, sensors, Direction.Left);
        sum += sumReachedBySensorInRow(targetBeacon, sensors, Direction.Right);
        return sum;
    }

    part2(input: string, ...args: unknown[]): string | number {
        const sensors = parseSensors(input);

        const min = 0;
        const max = 4000000;

        for (const sensor of sensors) {
            const edges = getDiamond(sensor, min, max);
            for (const edge of edges) {
                if (!isReachedBySensor(sensors, edge)) {
                    const isHole = DIRECTIONS.every((dir) => isReachedBySensor(sensors, edge.move(dir)));
                    if (isHole) {
                        return edge.x * 4000000 + edge.y;
                    }
                }
            }
        }

        return -1;
    }
}

function sumReachedBySensorInRow(beacon: Point, sensors: Sensor[], direction: Direction) {
    const { x: dx } = DIRECTION_MODIFIERS[direction];
    let sum = 0;

    for (let x = beacon.x + dx; ; x += dx) {
        const point = new Point(x, beacon.y);
        if (!isReachedBySensor(sensors, point)) {
            break;
        }
        sum += 1;
    }

    return sum;
}

function isReachedBySensor(sensors: Sensor[], point: Point) {
    for (const sensor of sensors) {
        const distance = sensor.point.manhattanDistance(point);
        if (sensor.range >= distance) {
            return true;
        }
    }
    return false;
}

function parseSensors(input: string) {
    const sensorRegex = /Sensor at x=(\d+), y=(\d+): closest beacon is at x=(\d+), y=(\d+)/;
    const sensors: Sensor[] = [];

    input.split('\n').forEach((line) => {
        const match = line.match(sensorRegex);
        if (match) {
            const beacon = new Point(toInt(match[3]), toInt(match[4]));
            const sensor = new Point(toInt(match[1]), toInt(match[2]));
            sensors.push({
                point: sensor,
                beacon: beacon,
                range: sensor.manhattanDistance(beacon),
            });
        }
    });

    return sensors;
}

interface Sensor {
    point: Point;
    beacon: Point;
    range: number;
}

function getDiamond(sensor: Sensor, min: number, max: number) {
    const radius = sensor.range + 1;
    const center = sensor.point;

    const points: Point[] = [
        new Point(center.x - radius), // left
        new Point(center.x + radius), // right
    ];

    const y = center.y;
    let dy = 1;
    let x = Math.max(min, center.x - radius) + 1;
    const maxX = Math.min(max, center.x + radius) - 1;

    while (x <= maxX) {
        points.push(new Point(x, y - dy)); // upper
        points.push(new Point(x, y + dy)); // lower
        dy += 1;
        x += 1;
    }

    return points;
}

function toInt(value: string = '') {
    return Number.parseInt(value, 10);
}