import { Solution } from './solution';

const solution = new Solution();
const input = solution.readInput();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`,
            args1: [],
            args2: [],
            want1: 18,
            want2: 54,
        },
        {
            input,
            args1: [],
            args2: [],
            want1: 292,
            want2: 816,
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1, args1 }) => {
        expect(solution.part1(input, ...args1)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2, args2 }) => {
        expect(solution.part2(input, ...args2)).toEqual(want2);
    });
});
