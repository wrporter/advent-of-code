import { Solution } from './solution';

const solution = new Solution();
const input = solution.readInput();

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
            args1: [10],
            args2: [20],
            want1: 26,
            want2: 56000011,
        },
        {
            input,
            args1: [2_000_000],
            args2: [4_000_000],
            want1: 4919281,
            want2: 12630143363767,
        }
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1, args1 }) => {
        expect(solution.part1(input, args1)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2, args2 }) => {
        expect(solution.part2(input, args2)).toEqual(want2);
    });
});
