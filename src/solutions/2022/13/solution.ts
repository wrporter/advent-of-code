import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 13;
    filename = 'input.txt';

    part1(input: string): string | number {
        const pairs = input.split('\n\n')
            .map((group) => group
                .split('\n')
                .map((v) => JSON.parse(v))
            );

        return pairs.reduce((sum, [a, b], i) => {
            if (compare(a, b) > 0) {
                return sum;
            }
            return sum + i + 1;
        }, 0);
    }

    part2(input: string): string | number {
        let packets = input.replaceAll('\n\n', '\n')
            .split('\n')
            .map((packet) => JSON.parse(packet));

        const dividerPackets = [[[2]], [[6]]];
        packets.push(...dividerPackets);
        packets = packets.sort(compare).map((packet) => JSON.stringify(packet));
        const [divider1, divider2] = dividerPackets.map((p) => JSON.stringify(p))

        return packets.reduce((result, packet, i) => {
            if (packet === divider1 || packet === divider2) {
                return result * (i + 1);
            }
            return result;
        }, 1);
    }
}

function compare(a: number | number[], b: number | number[]): number {
    if (typeof a === 'number' && typeof b === 'number') {
        return a - b;
    } else if (typeof a === 'number') {
        return compare([a], b);
    } else if (typeof b === 'number') {
        return compare(a, [b]);
    } else {
        for (let i = 0; i < a.length && i < b.length; i++) {
            const result = compare(a[i], b[i]);
            if (result !== 0) {
                return result;
            }
        }
        return a.length - b.length;
    }
}
