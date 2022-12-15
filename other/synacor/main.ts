import * as fs from 'fs';
import readlineSync from 'readline-sync';

function main() {
    const filename = 'other/synacor/challenge.bin';

    const buffer = fs.readFileSync(filename);
    const program = [];
    for (let offset = 0; offset < buffer.length / 2; offset++) {
        program.push(buffer.readUInt16LE(offset * 2));
    }

    const cpu = new Computer(program);
    // const cpu = new Computer([9, 32768, 32769, 82, 19, 32768]);
    // cpu.debug = true;
    cpu.run();
    // cpu.disassemble();
}

const REGISTER_START = 32768;
const REGISTER_END = 32775;

class Computer {
    private registers = [0, 0, 0, 0, 0, 0, 0, 0];
    private stack: number[] = [];
    private readonly memory: number[];
    public debug = false;

    constructor(program: number[]) {
        this.memory = program;
    }

    run() {
        for (let address = 0; address < this.memory.length;) {
            const opCode = this.memory[address] as OpCode;
            const operation = this.operations[opCode];
            if (!operation) {
                console.log(`// Unknown operation '${opCode}' at address ${address}`);
                address++;
                continue;
            }
            const args = this.memory.slice(address + 1, address + operation.length);
            args.unshift(address);
            const opAddress = operation.apply(this, args);
            const nextAddress =  address + (operation.length || 1);
            address = opAddress ?? nextAddress;
        }
    }

    disassemble() {
        for (let address = 0; address < this.memory.length;) {
            const opCode = this.memory[address] as OpCode;
            const operation = this.operations[opCode];
            if (!operation) {
                // console.log(`// Unknown operation '${opCode}' at address ${address}`);
                address++;
                continue;
            }
            const args: Array<string | number> = this.memory.slice(address, address + operation.length);
            address += operation.length || 1;

            if (opCode === OpCode.out && args[1] < REGISTER_START) {
                args.push(String.fromCharCode(args[1] as number)
                    .replace(/\n/g, '\\n'));
            }

            console.log(`[${address}]: ${OpCode[opCode]} ${args.slice(1).join(' ')}`);
        }
    }

    private get(value: number) {
        if (value >= REGISTER_START && value <= REGISTER_END) {
            return this.registers[value - REGISTER_START];
        }
        return value;
    }

    private set(address: number, value: number) {
        this.registers[address - REGISTER_START] = value;
    }

    private operations: { [key in OpCode]: OpFunc } = {
        // halt: 0 - stop execution and terminate the program
        [OpCode.halt]: () => {
            if (this.debug) {
                console.log(`>> halt`);
            }
            return this.memory.length;
        },
        // set: 1 a b - set register <a> to the value of <b>
        [OpCode.set]: (address, a, b) => {
            if (this.debug) {
                console.log(`>> set [${a}] = [${b}:${this.get(b)}]`);
            }
            this.set(a, this.get(b));
        },
        // push: 2 a - push <a> onto the stack
        [OpCode.push]: (address, a) => {
            this.stack.push(this.get(a));
            if (this.debug) {
                console.log(`>> push ${this.get(a)}`);
            }
        },
        // pop: 3 a - remove the top element from the stack and write it into <a>; empty stack = error
        [OpCode.pop]: (address, a) => {
            const value = this.stack.pop();
            if (value === undefined) {
                throw new Error('Attempted to pop from an empty stack.');
            }

            if (this.debug) {
                console.log(`>> pop [${a}] = ${value}`);
            }
            this.set(a, value);
        },
        // eq: 4 a b c - set <a> to 1 if <b> is equal to <c>; set it to 0 otherwise
        [OpCode.eq]: (address, a, b, c) => {
            if (this.debug) {
                console.log(`>> eq [${a}] = [${b}:${this.get(b)}] == [${c}:${this.get(c)}]`);
            }
            if (this.get(b) === this.get(c)) {
                this.set(a, 1);
            } else {
                this.set(a, 0);
            }
        },
        // gt: 5 a b c - set <a> to 1 if <b> is greater than <c>; set it to 0 otherwise
        [OpCode.gt]: (address, a, b, c) => {
            if (this.debug) {
                console.log(`>> gt [${a}] = [${b}:${this.get(b)}] > [${c}:${this.get(c)}]`);
            }
            if (this.get(b) > this.get(c)) {
                this.set(a, 1);
            } else {
                this.set(a, 0);
            }
        },
        // jmp: 6 a - jump to <a>
        [OpCode.jmp]: (address, a) => {
            if (this.debug) {
                console.log(`>> jmp [${a}:${this.get(a)}]`);
            }
            return this.get(a);
        },
        // jt: 7 a b - if <a> is nonzero, jump to <b>
        [OpCode.jt]: (address, a, b) => {
            if (this.debug) {
                console.log(`>> jt [${a}:${this.get(a)}] != 0 [${b}:${this.get(b)}]`);
            }
            if (this.get(a) !== 0) {
                return this.get(b);
            }
        },
        // jf: 8 a b - if <a> is zero, jump to <b>
        [OpCode.jf]: (address, a, b) => {
            if (this.debug) {
                console.log(`>> jf [${a}:${this.get(a)}] == 0 [${b}:${this.get(b)}]`);
            }
            if (this.get(a) === 0) {
                return this.get(b);
            }
        },
        // add: 9 a b c - assign into <a> the sum of <b> and <c> (modulo 32768)
        [OpCode.add]: (address, a, b, c) => {
            if (this.debug) {
                console.log(`>> add [${a}] = [${b}:${this.get(b)}] + [${c}:${this.get(c)}]`);
            }
            this.set(a, to15Bits(this.get(b) + this.get(c)));
        },
        // mult: 10 a b c - store into <a> the product of <b> and <c> (modulo 32768)
        [OpCode.mult]: (address, a, b, c) => {
            if (this.debug) {
                console.log(`>> mult [${a}] = [${b}:${this.get(b)}] * [${c}:${this.get(c)}]`);
            }
            this.set(a, to15Bits(this.get(b) * this.get(c)));
        },
        // mod: 11 a b c - store into <a> the remainder of <b> divided by <c>
        [OpCode.mod]: (address, a, b, c) => {
            if (this.debug) {
                console.log(`>> mod [${a}] = [${b}:${this.get(b)}] % [${c}:${this.get(c)}]`);
            }
            this.set(a, this.get(b) % this.get(c));
        },
        // and: 12 a b c - stores into <a> the bitwise and of <b> and <c>
        [OpCode.and]: (address, a, b, c) => {
            if (this.debug) {
                console.log(`>> and [${a}] = [${b}:${this.get(b)}] & [${c}:${this.get(c)}]`);
            }
            this.set(a, this.get(b) & this.get(c));
        },
        // or: 13 a b c - stores into <a> the bitwise or of <b> and <c>
        [OpCode.or]: (address, a, b, c) => {
            if (this.debug) {
                console.log(`>> or [${a}] = [${b}:${this.get(b)}] | [${c}:${this.get(c)}]`);
            }
            this.set(a, this.get(b) | this.get(c));
        },
        // not: 14 a b - stores 15-bit bitwise inverse of <b> in <a>
        [OpCode.not]: (address, a, b) => {
            if (this.debug) {
                console.log(`>> add [${a}] = ^[${b}:${this.get(b)}]`);
            }
            this.set(a, this.get(b) ^ 65535 % 32768);
        },
        // rmem: 15 a b - read memory at address <b> and write it to <a>
        [OpCode.rmem]: (address, a, b) => {
            if (this.debug) {
                console.log(`>> rmem [${a}:${this.get(a)}] = [${b}:${this.memory[this.get(b)]}]`);
            }
            this.set(a, this.memory[this.get(b)]);
        },
        // wmem: 16 a b - write the value from <b> into memory at address <a>
        [OpCode.wmem]: (address, a, b) => {
            if (this.debug) {
                console.log(`>> wmem [${a}:${this.get(a)}] = [${b}:${this.get(b)}]`);
            }
            this.memory[this.get(a)] = this.get(b);
        },
        // call: 17 a - write the address of the next instruction to the stack and jump to <a>
        [OpCode.call]: (address, a) => {
            if (this.memory[address + 2] !== OpCode.ret) {
                this.stack.push(address + 2);
            }
            if (this.debug) {
                console.log(`>> call [${a}:${this.get(a)}]`);
            }
            return this.get(a);
        },
        // ret: 18 - remove the top element from the stack and jump to it; empty stack = halt
        [OpCode.ret]: () => {
            const value = this.stack.pop();
            let address = this.memory.length;
            if (value !== undefined) {
                address = value;
            }
            if (this.debug) {
                console.log(`>> ret ${value}`);
            }
            return address;
        },
        // out: 19 a - write the character represented by ascii code <a> to the terminal
        [OpCode.out]: (address, a) => {
            if (this.debug) {
                process.stdout.write(`>> out [${a}:${this.get(a)}] - (`);
            }
            process.stdout.write(String.fromCharCode(this.get(a)));
            if (this.debug) {
                process.stdout.write(')\n');
            }
        },
        // in: 20 a - read a character from the terminal and write its ascii code to <a>; it can be assumed that once
        // input starts, it will continue until a newline is encountered; this means that you can safely read whole
        // lines from the keyboard and trust that they will be fully read
        [OpCode.in]: (address, a) => {
            if (this.debug) {
                process.stdout.write(`>> in [${a}:${this.get(a)}]`);
            }
            const input = readlineSync.prompt();
            this.set(a, input.charCodeAt(0));
        },
        // noop: 21 - no operation
        [OpCode.noop]: () => {
            if (this.debug) {
                process.stdout.write(`>> noop`);
            }
        },
    };
}

function to15Bits(value: number) {
    if (value >= 32768) {
        return value % 32768;
    } else {
        return value;
    }
}

enum OpCode {
    halt,
    set,
    push,
    pop,
    eq,
    gt,
    jmp,
    jt,
    jf,
    add,
    mult,
    mod,
    and,
    or,
    not,
    rmem,
    wmem,
    call,
    ret,
    out,
    in,
    noop,
}

type OpFunc = (...[address, ...args]: number[]) => number | void;

main();
