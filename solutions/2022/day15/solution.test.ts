import { Solution } from './solution';
import * as fs from 'fs';

const solution = new Solution();
const input = fs.readFileSync(`solutions/${solution.year}/day${solution.day}/${solution.filename}`, 'utf-8')

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`,
            args: [10],
            want1: 26,
            want2: 56000011,
        },
        {
            input,
            args: [2000000],
            want1: 4919281,
            want2: 12630143363767,
        }
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1, args }) => {
        expect(solution.part1(input, ...args)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2, args }) => {
        expect(solution.part2(input, ...args)).toEqual(want2);
    });
});
