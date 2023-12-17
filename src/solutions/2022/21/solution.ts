import { AbstractSolution } from '~/solution';
import algebra, { Equation } from 'algebra.js';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 21;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        const monkeys = parseMonkeys(input);
        return findNumber(monkeys);
    }

    part2(input: string, ...args: unknown[]): string | number {
        const monkeys = parseMonkeys(input);
        (monkeys.root as OperationMonkey).operator = '=';
        const equation = expandEquation(monkeys);

        const solution = (algebra.parse(equation) as Equation).solveFor('humn');
        return (solution as { numer: number }).numer;
    }
}

type Dictionary<T> = { [key: string]: T }

interface NumberMonkey {
    number: number;
}

interface OperationMonkey {
    left: string;
    operator: string;
    right: string;
}

type Monkey = NumberMonkey | OperationMonkey

function findNumber(monkeys: Dictionary<Monkey>, name = 'root'): number {
    if ((monkeys[name] as NumberMonkey).number) {
        return (monkeys[name] as NumberMonkey).number;
    }

    const monkey = monkeys[name] as OperationMonkey;
    const left = findNumber(monkeys, monkey.left);
    const right = findNumber(monkeys, monkey.right);

    if (monkey.operator === '=') {
        return left === right ? 1 : 0;
    } else if (monkey.operator === '+') {
        return left + right;
    } else if (monkey.operator === '-') {
        return left - right;
    } else if (monkey.operator === '*') {
        return left * right;
    } else {
        return left / right;
    }
}

function expandEquation(monkeys: Dictionary<Monkey>, name = 'root'): string {
    if (name === 'humn') {
        return 'humn';
    }
    if ((monkeys[name] as NumberMonkey).number) {
        return (monkeys[name] as NumberMonkey).number.toString();
    }

    const monkey = monkeys[name] as OperationMonkey;
    const left = expandEquation(monkeys, monkey.left);
    const right = expandEquation(monkeys, monkey.right);
    return `(${left}) ${monkey.operator} (${right})`;
}

function parseMonkeys(input: string) {
    return input.split('\n').reduce((monkeys, line) => {
        const parts = line.split(': ');
        const name = parts[0];
        if (isNumeric(parts[1])) {
            monkeys[name] = {
                number: Number.parseInt(parts[1], 10),
            };
        } else {
            const operation = parts[1].split(' ');
            monkeys[name] = {
                left: operation[0],
                operator: operation[1],
                right: operation[2],
            };
        }
        return monkeys;
    }, {} as Dictionary<Monkey>);
}

function isNumeric(str: string) {
    return !isNaN(parseInt(str));
}
