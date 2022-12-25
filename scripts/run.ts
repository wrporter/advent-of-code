import * as fs from 'fs';
import chalk from 'chalk';
import { sleep } from '~/';

function main() {
    const year = process.argv[2] ?? new Date().getFullYear();

    console.log(chalk.green(`
===============================
= Advent of Code: ${year}        =
===============================
`));
    sleep(1000);

    for (let day = 1; day <= 25; day++) {
        const file = `solutions/${year}/day${day}/main.ts`;
        if (fs.existsSync(file)) {
            require(file);
            console.log();
            sleep(1000);
        }
    }
}

main();
