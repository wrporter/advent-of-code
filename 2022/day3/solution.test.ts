import { Solution } from './solution';

const solution = new Solution();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `1000`,
            want1: 24000,
            want2: 45000,
        }
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1 }) => {
        expect(solution.part1(input)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2 }) => {
        expect(solution.part2(input)).toEqual(want2);
    });
});
