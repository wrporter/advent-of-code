import { Solution } from './solution';

const solution = new Solution();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
            want1: 13,
            want2: 1,
        },
        {
            input: `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`,
            want1: 88,
            want2: 36,
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1 }) => {
        expect(solution.part1(input)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2 }) => {
        expect(solution.part2(input)).toEqual(want2);
    });
});
