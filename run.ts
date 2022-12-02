import * as fs from 'fs';
import chalk from 'chalk';

function main() {
    const year = process.argv[2] ?? new Date().getFullYear();

    console.log(chalk.green(`
===============================
= Saving Santa for year: ${year} =
===============================
`));

    fs.readdirSync(`${year}`).forEach(async (dayDir) => {
        const { Solution } = await import(`./${year}/${dayDir}/solution.ts`);
        new Solution().run();
        console.log();
    });
}

main();
