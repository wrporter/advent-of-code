import { AbstractSolution } from '~/solution';

const items = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
const priority = items.split('')
    .reduce((priority: { [key: string]: number }, item, index) => {
        priority[item] = index + 1;
        return priority;
    }, {})

export class Solution extends AbstractSolution {
    year = 2022;
    day = 3;
    filename = 'input.txt';

    part1(input: string): string | number {
        const rucksacks = input.split('\n');

        const commonItems = rucksacks.reduce((commonItems: string[], rucksack) => {
            const middle = rucksack.length / 2;
            const [compartment1, compartment2] = [rucksack.slice(0, middle), rucksack.slice(middle)];

            let commonItem = '';
            for (let i = 0; i < compartment1.length && !commonItem; i++) {
                const char1 = compartment1[i];
                for (let j = 0; j < compartment2.length && !commonItem; j++) {
                    const char2 = compartment2[j];
                    if (char1 === char2) {
                        commonItem = char1;
                    }
                }
            }

            commonItems.push(commonItem);
            return commonItems;
        }, []);

        return commonItems.reduce(
            (sum, item) => sum + priority[item],
            0);
    }

    part2(input: string): string | number {
        const rucksacks = input.split('\n');

        let sum = 0;

        for (let i = 0; i < rucksacks.length; i += 3) {
            const rucksack1 = rucksacks[i];
            const rucksack2 = rucksacks[i + 1];
            const rucksack3 = rucksacks[i + 2];

            let commonItem = '';
            for (let i1 = 0; i1 < rucksack1.length && !commonItem; i1++) {
                const item1 = rucksack1[i1];
                for (let i2 = 0; i2 < rucksack2.length && !commonItem; i2++) {
                    const item2 = rucksack2[i2];
                    for (let i3 = 0; i3 < rucksack3.length && !commonItem; i3++) {
                        const item3 = rucksack3[i3];

                        if (item1 === item2 && item1 === item3) {
                            commonItem = item1;
                        }
                    }
                }
            }

            sum += priority[commonItem];
        }

        return sum;
    }
}
