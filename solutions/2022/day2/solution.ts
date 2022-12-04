import { AbstractSolution } from '~/';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 2;
    filename = 'input.txt';

    part1(input: string): string | number {
        const lines = input.split('\n');

        const shape: { [key: string]: number } = {
            // Opponent
            A: 1,
            B: 2,
            C: 3,

            // Self
            X: 1,
            Y: 2,
            Z: 3,
        };

        const outcome: { [key: string]: number } = {
            // Lose
            'A Z': 0,
            'B X': 0,
            'C Y': 0,

            // Draw
            'A X': 3,
            'B Y': 3,
            'C Z': 3,

            // Win
            'A Y': 6,
            'B Z': 6,
            'C X': 6,
        };

        const totalScore = lines.reduce((score, line) => {
            const [opponent, self] = line.split(' ');
            score += shape[self];
            score += outcome[`${opponent} ${self}`];
            return score;
        }, 0);

        return totalScore;
    }

    part2(input: string): string | number {
        const lines = input.split('\n');

        enum Shape {
            Rock = 1,
            Paper = 2,
            Scissors = 3,
        }

        const shapes: { [key: string]: number } = {
            A: Shape.Rock,
            B: Shape.Paper,
            C: Shape.Scissors,
        };

        const choices: { [key: string]: { shape: { [key: number]: Shape }, score: number } } = {
            X: { // Lose
                shape: {
                    [Shape.Rock]: Shape.Scissors,
                    [Shape.Paper]: Shape.Rock,
                    [Shape.Scissors]: Shape.Paper,
                },
                score: 0,
            },
            Y: { // Draw
                shape: {
                    [Shape.Rock]: Shape.Rock,
                    [Shape.Paper]: Shape.Paper,
                    [Shape.Scissors]: Shape.Scissors,
                },
                score: 3,
            },
            Z: { // Win
                shape: {
                    [Shape.Rock]: Shape.Paper,
                    [Shape.Paper]: Shape.Scissors,
                    [Shape.Scissors]: Shape.Rock,
                },
                score: 6,
            },
        };

        const totalScore = lines.reduce((score, line) => {
            const [opponent, end] = line.split(' ');
            const shape = shapes[opponent];
            const outcome = choices[end];

            score += outcome.shape[shape];
            score += outcome.score;
            return score;
        }, 0);

        return totalScore;
    }
}
