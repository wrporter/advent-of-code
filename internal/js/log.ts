import chalk from 'chalk';

export function logDay(year: number, day: number) {
    console.log(`🎄 ${chalk.green.underline(`${year}: Day ${day}`)}`);
}

export function logPart1(answer: number | string) {
    console.log(`⭐  ${chalk.green('Part 1:')} ${chalk.red(answer)}`);
}

export function logPart2(answer: number | string) {
    console.log(`⭐  ${chalk.green('Part 2:')} ${chalk.red(answer)}`);
}
