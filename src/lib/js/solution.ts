import { logDay, logPart1, logPart2 } from '~/log';
import { read } from '~/file';
import { Time } from '~/time';
import chalk from 'chalk';

export abstract class AbstractSolution {
    protected filename = 'input.txt';

    readInput() {
        return read(this.year, this.day, this.filename);
    }

    @Time(chalk.bold.bgBlue('Total:'))
    run(args1?: unknown[], args2?: unknown[]) {
        logDay(this.year, this.day);
        const input = this.readInput();
        this.solvePart1(input, args1);
        this.solvePart2(input, args2);
    }

    @Time('Part 1:')
    protected solvePart1(input: string, args?: unknown[]) {
        const answer1 = this.part1(input, args);
        logPart1(answer1);
    }

    @Time('Part 2:')
    protected solvePart2(input: string, args?: unknown[]) {
        const answer2 = this.part2(input, args);
        logPart2(answer2);
    }

    abstract get year(): number;
    abstract get day(): number;

    abstract part1(input: string, ...args: unknown[]): string | number;

    abstract part2(input: string, ...args: unknown[]): string | number;
}

export interface TestCase {
    input: string;
    args1?: any[];
    args2?: any[];
    want1: any;
    want2: any;
}
