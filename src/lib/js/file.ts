import * as fs from 'fs';

export function read(year: number, day: number, filename: string): string {
    return fs.readFileSync(`solutions/${year}/${pad(day)}/${filename}`, 'utf8');
}

export function pad(day: number) {
    return `${day}`.padStart(2, '0');
}