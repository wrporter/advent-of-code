import { logDay, logPart1, logPart2 } from '~/log';
import { read } from '~/file';
import { Time } from '~/time';
import chalk from 'chalk';

export abstract class AbstractSolution {
    protected filename = 'input.txt';

    readInput(year: number, day: number) {
        return read(year, day, this.filename);
    }

    @Time(chalk.bold.bgBlue('Total:'))
    run() {
        logDay(this.year, this.day);
        const input = this.readInput(this.year, this.day);
        this.solvePart1(input);
        this.solvePart2(input);
    }

    @Time('Part 1:')
    protected solvePart1(input: string) {
        const answer1 = this.part1(input);
        logPart1(answer1);
    }

    @Time('Part 2:')
    protected solvePart2(input: string) {
        const answer2 = this.part2(input);
        logPart2(answer2);
    }

    abstract get year(): number;
    abstract get day(): number;

    abstract part1(input: string): string | number;

    abstract part2(input: string): string | number;
}
