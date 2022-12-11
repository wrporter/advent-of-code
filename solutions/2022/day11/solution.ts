import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 11;
    filename = 'input.txt';

    part1(input: string): string | number {
        return calculateMonkeyBusiness(input, 20, 3);
    }

    part2(input: string): string | number {
        return calculateMonkeyBusiness(input, 10_000);
    }
}

// const startingItemsRegex = /^  Starting items: (\d+(?:, \d+)*)$/;
const START = '  Starting items: ';
const OPERATION = '  Operation: new = old ';

function calculateMonkeyBusiness(input: string, maxRounds: number, worryDivider?: number) {
    const states = input.split('\n\n');
    let modulus = 1;

    const monkeys: Monkey[] = states.map((state, index) => {
        const situation = state.split('\n');
        const modifier = situation[2][OPERATION.length];
        const by = situation[2].slice(OPERATION.length + 2);
        const divisibleBy = Number.parseInt(situation[3].slice(situation[3].length - 2), 10);
        modulus *= divisibleBy;

        // const items = state.match(startingItemsRegex);

        return {
            id: index,
            items: situation[1]
                .slice(START.length)
                .split(', ')
                .map((v) => Number.parseInt(v, 10)),
            operation: (level: number) => {
                let amount = Number.parseInt(by);
                if (by === 'old') {
                    amount = level;
                }

                if (modifier === '+') {
                    return level + amount;
                } else if (modifier === '*') {
                    return level * amount;
                }
                throw new Error(`Unknown modifier '${modifier}'`);
            },
            test: {
                divisibleBy,
                ifTrue: Number.parseInt(situation[4].slice(situation[4].length - 1), 10),
                ifFalse: Number.parseInt(situation[5].slice(situation[5].length - 1), 10),
            },
            numItemsInspected: 0,
        };
    });

    for (let round = 0; round < maxRounds; round++) {
        for (let turn = 0; turn < monkeys.length; turn++) {
            const monkey = monkeys[turn];
            if (monkey.items.length === 0) {
                continue;
            }

            monkey.items = monkey.items.reverse();
            for (let i = monkey.items.length - 1; i >= 0; i--) {
                const item = monkey.items[i];
                const { divisibleBy, ifTrue, ifFalse } = monkey.test;
                monkey.numItemsInspected += 1;

                let level = monkey.operation(item);
                if (worryDivider) {
                    level = Math.floor(level / worryDivider);
                } else {
                    level = level % modulus;
                }

                if (level % divisibleBy === 0) {
                    monkeys[ifTrue].items.push(level);
                } else {
                    monkeys[ifFalse].items.push(level);
                }
                monkey.items.splice(i, 1);
            }
        }
    }

    monkeys.sort((a, b) => b.numItemsInspected - a.numItemsInspected);
    const monkeyBusiness = monkeys[0].numItemsInspected * monkeys[1].numItemsInspected;

    return monkeyBusiness;
}

interface Monkey {
    id: number;
    items: number[];
    operation: (level: number) => number;
    test: {
        divisibleBy: number;
        ifTrue: number;
        ifFalse: number;
    };
    numItemsInspected: number;
}
