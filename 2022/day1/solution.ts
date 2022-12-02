import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 1;
    filename = 'input.txt';

    part1(input: string): string | number {
        const lists = input.split('\n\n');
        let max = 0;

        lists.forEach((listString) => {
            const total = listString.split('\n').reduce((sum, calorieString) => {
                sum += Number.parseInt(calorieString, 10);
                return sum;
            }, 0);

            if (total > max) {
                max = total;
            }
        });

        return max;
    }

    part2(input: string): string | number {
        const lists = input.split('\n\n');

        const totals = lists.reduce((result: number[], listString) => {
            const total = listString.split('\n').reduce((sum, calorieString) => {
                sum += Number.parseInt(calorieString, 10);
                return sum;
            }, 0);

            result.push(total);
            return result;
        }, []);

        totals.sort((a, b) => b - a);
        return totals[0] + totals[1] + totals[2];
    }
}
