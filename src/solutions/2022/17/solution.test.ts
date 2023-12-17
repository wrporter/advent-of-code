import { Solution } from './solution';

const solution = new Solution();
const input = solution.readInput();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`,
            args1: [],
            args2: [],
            want1: 3068,
            want2: 1_514_285_714_288,
        },
        {
            input,
            args1: [],
            args2: [],
            want1: 3168,
            want2: 1554117647070,
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1, args1 }) => {
        expect(solution.part1(input, ...args1)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2, args2 }) => {
        expect(solution.part2(input, ...args2)).toEqual(want2);
    });
});
