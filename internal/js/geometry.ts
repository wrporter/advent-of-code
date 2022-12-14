export enum Direction {
    Up = 'Up',
    Right = 'Right',
    Down = 'Down',
    Left = 'Left',
}

export const DIRECTION_MODIFIERS = {
    [Direction.Up]: { x: 0, y: -1 },
    [Direction.Right]: { x: 1, y: 0 },
    [Direction.Down]: { x: 0, y: 1 },
    [Direction.Left]: { x: -1, y: 0 },
};

export const DIRECTIONS = [
    Direction.Up,
    Direction.Down,
    Direction.Left,
    Direction.Right,
];

export class Point {
    constructor(public x = 0, public y = 0) {}

    move(direction: Direction, amount = 1) {
        const { x: mx, y: my } = DIRECTION_MODIFIERS[direction];
        const x = this.x + (mx * amount);
        const y = this.y + (my * amount);
        return new Point(x, y);
    }

    up(amount = 1) {
        return this.move(Direction.Up, amount);
    }

    down(amount = 1) {
        return this.move(Direction.Down, amount);
    }

    downLeft(amount = 1) {
        const x = this.x + (-1 * amount);
        const y = this.y + amount;
        return new Point(x, y);
    }

    downRight(amount = 1) {
        const x = this.x + amount;
        const y = this.y + amount;
        return new Point(x, y);
    }

    left(amount = 1) {
        return this.move(Direction.Left, amount);
    }

    right(amount = 1) {
        return this.move(Direction.Right, amount);
    }

    manhattanDistance({ x, y }: Point) {
        return Math.abs(this.x - x) + Math.abs(this.y - y);
    }

    stepDistance({ y, x }: Point) {
        return Math.max(Math.abs(this.y - y), Math.abs(this.x - x));
    }

    equals({ x, y }: Point) {
        return this.x === x && this.y === y;
    }

    clone() {
        return new Point(this.x, this.y);
    }

    static fromString(coordinates: string) {
        const [x, y] = coordinates.split(',')
            .map((v) => Number.parseInt(v, 10));
        return new Point(x, y);
    }

    toString() {
        return `${this.x},${this.y}`;
    }
}
