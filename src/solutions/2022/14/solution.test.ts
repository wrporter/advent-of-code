import { Solution } from './solution';

const solution = new Solution();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`,
            want1: 24,
            want2: 93,
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1 }) => {
        expect(solution.part1(input)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2 }) => {
        expect(solution.part2(input)).toEqual(want2);
    });
});
