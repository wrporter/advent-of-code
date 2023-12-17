import { Solution } from './solution';

const solution = new Solution();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `mjqjpqmgbljsphdztnvjfqwrcgsmlb`,
            want1: 7,
            want2: 19,
        },
        {
            input: `bvwbjplbgvbhsrlpgdmjqwftvncz`,
            want1: 5,
            want2: 23,
        },
        {
            input: `nppdvjthqldpwncqszvftbrmjlhg`,
            want1: 6,
            want2: 23,
        },
        {
            input: `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`,
            want1: 10,
            want2: 29,
        },
        {
            input: `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`,
            want1: 11,
            want2: 26,
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1 }) => {
        expect(solution.part1(input)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2 }) => {
        expect(solution.part2(input)).toEqual(want2);
    });
});
