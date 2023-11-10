import { Solution } from './solution';

const solution = new Solution();
const input = solution.readInput();
interface TestCase {
    input: string;
    args1?: any[];
    args2?: any[];
    want1: any;
    want2: any;
}

describe(`Day ${solution.day}`, () => {
    const tests: TestCase[] = [
        {
            input: `1000`,
            want1: 'TBD',
            want2: 'TBD',
        },
        {
            input,
            want1: 'TBD',
            want2: 'TBD',
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1, args1 }) => {
        expect(solution.part1(input, args1)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2, args2 }) => {
        expect(solution.part2(input, args2)).toEqual(want2);
    });
});
