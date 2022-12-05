import { AbstractSolution } from '~/solution';

const moveRegex = /move (\d+) from (\d+) to (\d+)/;
const rowRegex = /.{4}/g
const whitespaceRegex = /^\s+$/

interface Move {
    amount: number;
    from: number;
    to: number;
}

export class Solution extends AbstractSolution {
    year = 2022;
    day = 5;
    filename = 'input.txt';

    part1(input: string): string | number {
        const { moves, crates } = this.parseInput(input);

        moves.forEach(({ amount, from, to }) => {
            for (let moved = 0; moved < amount; moved++) {
                const crate = crates[from].pop();
                if (crate) {
                    crates[to].push(crate);
                }
            }
        });

        return this.joinTopCrates(crates);
    }

    part2(input: string): string | number {
        const { moves, crates } = this.parseInput(input);

        moves.forEach(({ amount, from, to }) => {
            const cratesToMove = crates[from].splice(crates[from].length - amount);
            if (crates.length > 0) {
                crates[to].push(...cratesToMove);
            }
        });

        return this.joinTopCrates(crates);
    }

    private parseInput(input: string) {
        const [crateStr, moveStr] = input.split('\n\n');
        const crateLines = crateStr.split('\n');
        const moveStrings = moveStr.split('\n');

        crateLines.pop(); // remove the numbers
        const crates: string[][] = crateLines.reduce((crates, crateLine) => {
            const crateRow = (' ' + crateLine).match(rowRegex);
            if (!crateRow) {
                return crates;
            }
            crateRow.forEach((crate, column) => {
                if (!crates[column]) {
                    crates.push([]);
                }
                if (whitespaceRegex.test(crate)) {
                    return;
                }
                crates[column].unshift(crate[2]);
            });
            return crates;
        }, [] as string[][]);

        const moves = this.parseMoves(moveStrings);

        return { moves, crates };
    }

    private parseMoves(moves: string[]): Move[]  {
        return moves.map((move) => {
            const match = move.match(moveRegex);
            if (!match) {
                return;
            }

            return {
                amount: Number.parseInt(match[1], 10),
                from: Number.parseInt(match[2], 10) - 1,
                to: Number.parseInt(match[3], 10) - 1,
            };
        });
    }

    private joinTopCrates(crates: string[][]) {
        return crates.reduce((result, crateColumn) => {
            if (crateColumn.length === 0) {
                return result;
            }
            return result + crateColumn[crateColumn.length - 1];
        }, '');
    }
}
