import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 8;
    filename = 'input.txt';

    part1(input: string): string | number {
        const grid = this.parseGrid(input);

        // count edges
        let visible = (grid.length * 2) + (grid[0].length * 2) - 4;
        const counted: { [key: string]: boolean } = {};

        const count = (tallest: number, row: number, col: number) => {
            const key = `${row},${col}`;
            const height = grid[row][col];
            if (height > tallest && !counted[key]) {
                visible += 1;
                counted[key] = true;
            }
            return Math.max(tallest, height);
        };

        // top
        for (let col = 1; col < grid[0].length - 1; col++) {
            let tallest = grid[0][col];
            for (let row = 1; row < grid.length - 1; row++) {
                tallest = count(tallest, row, col);
            }
        }
        // left
        for (let row = 1; row < grid.length - 1; row++) {
            let tallest = grid[row][0];
            for (let col = 1; col < grid[row].length - 1; col++) {
                tallest = count(tallest, row, col);
            }
        }
        // right
        for (let row = 1; row < grid.length - 1; row++) {
            let tallest = grid[row][grid[row].length - 1];
            for (let col = grid[row].length - 2; col > 0; col--) {
                tallest = count(tallest, row, col);
            }
        }
        // bottom
        for (let col = 1; col < grid[0].length - 1; col++) {
            let tallest = grid[grid.length - 1][col];
            for (let row = grid.length - 2; row > 0; row--) {
                tallest = count(tallest, row, col);
            }
        }

        return visible;
    }

    part2(input: string): string | number {
        const grid = this.parseGrid(input);
        let maxScore = 0;

        for (let treeRow = 1; treeRow < grid.length - 1; treeRow++) {
            for (let treeCol = 1; treeCol < grid.length - 1; treeCol++) {
                let score = 1;
                let visible = 0;
                const height = grid[treeRow][treeCol];

                // down
                visible = 1;
                for (let row = treeRow + 1; row < grid.length-1 && grid[row][treeCol] < height; row++) {
                    visible += 1;
                }
                score = score * (visible);
                // up
                visible = 1;
                for (let row = treeRow - 1; row > 0 && grid[row][treeCol] < height; row--) {
                    visible += 1;
                }
                score = score * (visible);
                // right
                visible = 1;
                for (let col = treeCol + 1; col < grid[treeRow].length-1 && grid[treeRow][col] < height; col++) {
                    visible += 1;
                }
                score = score * (visible);
                // left
                visible = 1;
                for (let col = treeCol - 1; col > 0 && grid[treeRow][col] < height; col--) {
                    visible += 1;
                }
                score = score * (visible);

                maxScore = Math.max(maxScore, score);
            }
        }

        return maxScore;
    }

    private parseGrid(input: string) {
        const lines = input.split('\n');
        const grid: number[][] = [];
        lines.map((line) => {
            grid.push(
                line.split('')
                    .map((value) => Number.parseInt(value, 10))
            );
        });
        return grid;
    }
}
