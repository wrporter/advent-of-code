import { Solution } from './solution';

const solution = new Solution();
const input = solution.readInput();

describe(`Day ${solution.day}`, () => {
    const tests = [
        {
            input: `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`,
            args1: [],
            args2: [],
            want1: 1651,
            want2: 1707,
        },
        {
            input,
            args1: [],
            args2: [],
            want1: 1724,
            want2: 2283,
        },
    ];

    test.each(tests)('Part 1 - Test %#', ({ input, want1, args1 }) => {
        expect(solution.part1(input, ...args1)).toEqual(want1);
    });

    test.each(tests)('Part 2 - Test %#', ({ input, want2, args2 }) => {
        expect(solution.part2(input, ...args2)).toEqual(want2);
    });
});
