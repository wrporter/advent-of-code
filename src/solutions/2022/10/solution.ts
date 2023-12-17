import { AbstractSolution } from '~/solution';
import { gridToString } from '~/';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 10;
    filename = 'input.txt';

    part1(input: string): string | number {
        const instructions = input.split('\n');
        let signalStrength = 0;

        for (const { cycle, registerX } of cpu(instructions, 240)) {
            if ((cycle - 20) % 40 === 0) {
                signalStrength += cycle * registerX;
            }
        }

        return signalStrength;
    }

    part2(input: string): string | number {
        const instructions = input.split('\n');
        const crt: string[][] = Array.from({ length: 6 }, () => Array.from({ length: 40 }, () => '.'));

        for (const { cycle, registerX } of cpu(instructions, 240)) {
            const row = Math.floor((cycle - 1) / 40);
            const col = (cycle - 1) % 40;
            const isWithinSprite = col >= registerX - 1 && col <= registerX + 1;
            crt[row][col] = isWithinSprite ? '#' : '.';
        }

        return `\n${gridToString(crt)}`;
    }
}

function* cpu(instructions: string[], maxCycles: number) {
    let address = 0;
    let registerX = 1;
    let operation = parseInstruction(instructions, address);

    for (let cycle = 1; cycle <= maxCycles && address < instructions.length; cycle++) {
        if (operation.cycles === 0) {
            registerX = operation.execute(registerX, ...operation.args);
            address += 1;
            operation = parseInstruction(instructions, address);
        }
        operation.cycles -= 1;

        yield { cycle, registerX };
    }
}

function parseInstruction(instructions: string[], address: number): Operation {
    const [opCode, ...args] = instructions[address].split(' ');
    const { execute, ...fields } = Operations[opCode as OpCode];
    return {
        ...fields,
        execute,
        args: (args ?? []).map((v) => Number.parseInt(v, 10)),
    };
}

enum OpCode {
    noop = 'noop',
    addx = 'addx',
}

interface Operation {
    opCode: OpCode;
    cycles: number;
    args: number[];
    execute: (registerX: number, ...args: number[]) => number;
}

const Operations = {
    [OpCode.noop]: { opCode: OpCode.noop, cycles: 1, execute: (registerX: number) => registerX },
    [OpCode.addx]: { opCode: OpCode.addx, cycles: 2, execute: (registerX: number, arg1: number) => registerX + arg1 },
};
