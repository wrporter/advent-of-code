import { Solution } from './solution';

const solution = new Solution();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`,
            want1: 31,
            want2: 29,
        },
        {
            input: `Szzzzzqp
bzzzzzro
cEzzzzsn
dyxwvutm
efghijkl`,
            want1: 25,
            want2: 25,
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1 }) => {
        expect(solution.part1(input)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2 }) => {
        expect(solution.part2(input)).toEqual(want2);
    });
});
