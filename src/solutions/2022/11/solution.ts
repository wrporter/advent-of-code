import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 11;
    filename = 'input.txt';

    part1(input: string): string | number {
        const monkeys = parseMonkeyStates(input);
        const manageWorry = (worry: number) => Math.floor(worry / 3);
        return calculateMonkeyBusiness(monkeys, 20, manageWorry);
    }

    part2(input: string): string | number {
        const monkeys = parseMonkeyStates(input);

        // Chinese Remainder Theorem: Get product of all prime moduli
        const modulus = monkeys.reduce((product, monkey) => product * monkey.test.divisibleBy, 1);

        const manageWorry = (worry: number) => worry % modulus;
        return calculateMonkeyBusiness(monkeys, 10_000, manageWorry);
    }
}

function calculateMonkeyBusiness(monkeys: Monkey[], maxRounds: number, manageWorry: (worry: number) => number) {
    for (let round = 0; round < maxRounds; round++) {
        monkeys.forEach((monkey) => {
            if (monkey.items.length === 0) {
                return; // skip monkeys with no items
            }

            const { divisibleBy, trueMonkey, falseMonkey } = monkey.test;
            monkey.items.forEach((item) => {
                monkey.numItemsInspected += 1;

                let worry = monkey.operation(item);
                worry = manageWorry(worry);

                if (worry % divisibleBy === 0) {
                    monkeys[trueMonkey].items.push(worry);
                } else {
                    monkeys[falseMonkey].items.push(worry);
                }
            });
            monkey.items = [];
        });
    }

    monkeys.sort((a, b) => b.numItemsInspected - a.numItemsInspected);
    return monkeys[0].numItemsInspected * monkeys[1].numItemsInspected;
}

const startingItemsRegex = / {2}Starting items: (\d+(?:, \d+)*)/;
const operationRegex = / {2}Operation: new = old ([+*]) (old|\d+)/;
const testRegex = / {2}Test: divisible by (\d+)/;
const testTrueRegex = / {4}If true: throw to monkey (\d+)/;
const testFalseRegex = / {4}If false: throw to monkey (\d+)/;

function parseMonkeyStates(input: string) {
    const states = input.split('\n\n');
    const monkeys: Monkey[] = states.map((state, id) => {
        const items = state.match(startingItemsRegex)?.[1]
            .split(', ').map(toInt) ?? [];

        const operationMatch = state.match(operationRegex);
        const operator = operationMatch?.[1];
        const modifier = operationMatch?.[2];
        const modifierInt = toInt(modifier);

        const divisibleBy = toInt(state.match(testRegex)?.[1]);

        return {
            id,
            items,
            operation: (level: number) => {
                let amount = modifierInt;
                if (modifier === 'old') {
                    amount = level;
                }

                if (operator === '+') {
                    return level + amount;
                } else if (operator === '*') {
                    return level * amount;
                }
                throw new Error(`Unknown operator '${operator}'`);
            },
            test: {
                divisibleBy,
                trueMonkey: toInt(state.match(testTrueRegex)?.[1]),
                falseMonkey: toInt(state.match(testFalseRegex)?.[1]),
            },
            numItemsInspected: 0,
        };
    });
    return monkeys;
}

interface Monkey {
    id: number;
    items: number[];
    operation: (level: number) => number;
    test: {
        divisibleBy: number;
        trueMonkey: number;
        falseMonkey: number;
    };
    numItemsInspected: number;
}

function toInt(value: string = '') {
    return Number.parseInt(value, 10);
}
