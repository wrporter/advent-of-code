import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 20;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        const list = parseList(input, 1);
        mix(list);
        return calculateCoordinates(list);
    }

    part2(input: string, ...args: unknown[]): string | number {
        const list = parseList(input, 811589153);
        for (let i = 0; i < 10; i++) {
            mix(list);
        }
        return calculateCoordinates(list);
    }
}

function calculateCoordinates(list: Node[]) {
    let sum = 0;
    let current = list.find((n) => n.value === 0) as Node;
    for (let i = 0; i < 3; i++) {
        for (let j = 0; j < 1000; j++) {
            current = current.next;
        }
        sum += current.value;
    }
    return sum;
}

function mix(list: Node[]) {
    for (const node of list) {
        let left = node.prev;
        let right = node.next;
        left.next = right;
        right.prev = left;

        const moves = mod(node.value, list.length - 1);
        for (let move = 0; move < moves; move++) {
            left = left.next;
            right = right.next;
        }

        left.next = node;
        right.prev = node;

        node.prev = left;
        node.next = right;
    }
}

function parseList(input: string, decryptionKey: number) {
    return input.split('\n')
        .map((value) => new Node(Number.parseInt(value, 10) * decryptionKey))
        .map((node, index, list) => {
            const prev = list[mod(index - 1, list.length)];
            const next = list[(index + 1) % list.length];
            node.prev = prev;
            node.next = next;
            return node;
        });
}

class Node {
    // Always assume these are correctly set.
    public prev!: Node;
    public next!: Node;

    constructor(public value: number) {}
}

// mod allows for modular arithmetic with negative numbers wrapping backwards.
function mod(value: number, modulo: number) {
    return ((value % modulo) + modulo) % modulo;
}

