import { Solution } from './solution';

const solution = new Solution();
const input = solution.readInput();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `1
2
-3
3
-2
0
4`,
            args1: [],
            args2: [],
            want1: 3,
            want2: 1623178306,
        },
        {
            input,
            args1: [],
            args2: [],
            want1: 27726,
            want2: 4275451658004,
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1, args1 }) => {
        expect(solution.part1(input, ...args1)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2, args2 }) => {
        expect(solution.part2(input, ...args2)).toEqual(want2);
    });
});
