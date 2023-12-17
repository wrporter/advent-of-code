import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 25;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        const snafus = input.split('\n');
        const sum = snafus.reduce((sum, snafu) => sum + toDecimal(snafu), 0);
        return toSnafu(sum);
    }

    part2(input: string, ...args: unknown[]): string | number {
        return 'Merry Christmas! 🎄';
    }
}

const SNAFU_DIGITS = '=-012';

function toDecimal(snafu: string): number {
    if (snafu.length === 0) {
        return 0;
    }
    const digit = snafu[snafu.length - 1];
    const remainder = snafu.slice(0, snafu.length - 1);
    const decimal = SNAFU_DIGITS.indexOf(digit) - 2;
    return toDecimal(remainder) * 5 + decimal;
}

function toSnafu(decimal: number): string {
    if (decimal === 0) {
        return '';
    }
    const quotient = Math.floor((decimal + 2) / 5);
    const remainder = (decimal + 2) % 5;
    return toSnafu(quotient) + SNAFU_DIGITS[remainder];
}
