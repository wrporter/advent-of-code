import { permutations } from './probability';

const tests = [
    {
        args: [['a']],
        want: [['a']],
    },
    {
        args: [['a', 'b']],
        want: [['a'], ['b'], ['a', 'b'], ['b', 'a']],
    },
    {
        args: [['a', 'b', 'c']],
        want: [
            ['a'],
            ['b'],
            ['c'],
            ['a', 'b'],
            ['a', 'c'],
            ['b', 'a'],
            ['b', 'c'],
            ['c', 'a'],
            ['c', 'b'],
            ['a', 'b', 'c'],
            ['a', 'c', 'b'],
            ['b', 'a', 'c'],
            ['b', 'c', 'a'],
            ['c', 'a', 'b'],
            ['c', 'b', 'a'],
        ],
    },
];

test.each(tests)('permutations %#', ({ args, want }) => {
    expect(permutations(args[0])).toEqual(want);
});
