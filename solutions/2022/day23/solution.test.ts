import { Solution } from './solution';

const solution = new Solution();
const input = solution.readInput();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `.....
..##.
..#..
.....
..##.
.....`,
            args1: [],
            args2: [],
            want1: 25,
            want2: 4,
        },
        {
            input: `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`,
            args1: [],
            args2: [],
            want1: 110,
            want2: 20,
        },
        {
            input,
            args1: [],
            args2: [],
            want1: 3906,
            want2: 895,
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1, args1 }) => {
        expect(solution.part1(input, ...args1)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2, args2 }) => {
        expect(solution.part2(input, ...args2)).toEqual(want2);
    });
});
