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
            input: `32T3K 765
	T55J5 684
	KK677 28
	KTJJT 220
	QQQJA 483`,
            want1: 6440,
            want2: 5905,
        },
        {
            input,
            want1: 253866470,
            want2: 254494947,
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1, args1 }) => {
        expect(solution.part1(input, args1)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2, args2 }) => {
        expect(solution.part2(input, args2)).toEqual(want2);
    });
});
