import { AbstractSolution } from '~/solution';

const MARKER_SIZE = 4;
const MESSAGE_SIZE = 14;

export class Solution extends AbstractSolution {
    year = 2022;
    day = 6;
    filename = 'input.txt';

    part1(input: string): string | number {
        return this.lastIndexOfPacket(input, MARKER_SIZE);
    }

    part2(input: string): string | number {
        return this.lastIndexOfPacket(input, MESSAGE_SIZE);
    }

    private lastIndexOfPacket(input: string, size: number) {
        for (let i = 0; i < input.length - size; i++) {
            const marker = input.slice(i, i + size);
            const unique = new Set(marker.split(''));
            if (unique.size === size) {
                return i + size;
            }
        }
        return -1;
    }
}
