import { Solution } from './solution';

const solution = new Solution();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`,
            want1: 157,
            want2: 70,
        }
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1 }) => {
        expect(solution.part1(input)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2 }) => {
        expect(solution.part2(input)).toEqual(want2);
    });
});