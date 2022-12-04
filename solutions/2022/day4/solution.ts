import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 4;
    filename = 'input.txt';

    part1(input: string): string | number {
        const lines = input.split('\n');

        return lines.reduce((numOverlappingPairs, line) => {
            const { low1, high1, low2, high2 } = this.parseLine(line);

            if ((low1 >= low2 && high1 <= high2) || (low2 >= low1 && high2 <= high1)) {
                numOverlappingPairs += 1;
            }
            return numOverlappingPairs
        }, 0);
    }

    part2(input: string): string | number {
        const lines = input.split('\n');

        return lines.reduce((numOverlappingPairs, line) => {
            const { low1, high1, low2, high2 } = this.parseLine(line);

            if ((low1 >= low2 && low1 <= high2) || (high1 >= low2 && high1 <= high2) ||
                (low2 >= low1 && low2 <= high1) || (high2 >= low1 && high2 <= high1)) {
                numOverlappingPairs += 1;
            }
            return numOverlappingPairs
        }, 0);
    }

    private parseLine(line: string) {
        const [section1, section2] = line.split(',');
        const [low1, high1] = section1.split('-').map(v => Number.parseInt(v, 10));
        const [low2, high2] = section2.split('-').map(v => Number.parseInt(v, 10));
        return { low1, high1, low2, high2 };
    }
}
