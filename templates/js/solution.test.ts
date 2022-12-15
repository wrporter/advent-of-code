import { Solution } from './solution';

const solution = new Solution();
const input = solution.readInput();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `1000`,
            args1: [],
            args2: [],
            want1: 0,
            want2: 'TBD',
        },
        {
            input,
            args1: [],
            args2: [],
            want1: 0,
            want2: 'TBD',
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1, args1 }) => {
        expect(solution.part1(input, ...args1)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2, args2 }) => {
        expect(solution.part2(input, ...args2)).toEqual(want2);
    });
});
