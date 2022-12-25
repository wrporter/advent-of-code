import { Solution } from './solution';

const solution = new Solution();
const input = solution.readInput();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`,
            args1: [],
            args2: [],
            want1: '2=-1=0',
            want2: 'TBD',
        },
        {
            input,
            args1: [],
            args2: [],
            want1: '2=20---01==222=0=0-2',
            want2: 'TBD',
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1, args1 }) => {
        expect(solution.part1(input, ...args1)).toEqual(want1);
    });
});
