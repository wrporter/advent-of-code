import * as fs from 'fs';

export function read(year: number, day: number, filename: string): string {
    return fs.readFileSync(`${year}/day${day}/${filename}`, 'utf8');
}