import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2023;
    day = 7;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        const hands = parseInput(input);
        sortHands(hands);
        return sumWinnings(hands);
    }

    part2(input: string, ...args: unknown[]): string | number {
        const hands = parseInput(input.replace(/J/g, 'W'));
        sortHands(hands);
        return sumWinnings(hands);
    }
}

const strength = 'W23456789TJQKA'.split('').reduce((acc, card, i) => {acc[card]=i; return acc}, {} as Record<string, number>)

const highCard = 0;
const onePair = 1;
const twoPair = 2;
const threeOfAKind = 3;
const fullHouse = 4;
const fourOfAKind = 5;
const fiveOfAKind = 6;

interface CamelHand {
    cards: string;
    bid: number;
    type: number;
}

function parseInput(input: string): CamelHand[] {
    return input.split("\n").map((line) => {
        const parts = line.split(" ");
        const cards = parts[0];
        const bid = parseInt(parts[1]);

        const counts: Record<string, number> = {};
        let most = 0;
        let wilds = 0;

        for (const card of cards) {
            if (card !== 'W') {
                if (counts[card]) {
                    counts[card]++
                } else {
                    counts[card] = 1
                }
                most = Math.max(most, counts[card])
            } else {
                wilds++
            }
        }

        most = most + wilds;
        const countSize = Object.keys(counts).length
        let type: number;

        if (most === 5) {
            type = fiveOfAKind;
        } else if (most === 4) {
            type = fourOfAKind;
        } else if (countSize === 2 && most === 3) {
            type = fullHouse;
        } else if (most === 3) {
            type = threeOfAKind;
        } else if (countSize === 3 && most === 2) {
            type = twoPair;
        } else if (most === 2) {
            type = onePair;
        } else {
            type = highCard;
        }

        return {cards, bid, type};
    });
}

function sortHands(hands: CamelHand[]) {
    hands.sort((hand1, hand2) => {
        if (hand1.type === hand2.type) {
            for (let c = 0; c < hand1.cards.length; c++) {
                const strength1 = strength[hand1.cards[c]];
                const strength2 = strength[hand2.cards[c]];

                if (strength1 < strength2) {
                    return -1;
                } else if (strength1 > strength2) {
                    return 1;
                }
            }
            return 0;
        }

        return hand1.type - hand2.type;
    });
}

function sumWinnings(hands: CamelHand[]) {
    return hands.reduce((totalWinnings, hand, i) => {
        const rank = i + 1;
        return totalWinnings + hand.bid * rank;
    }, 0)
}