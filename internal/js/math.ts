// mod allows for modular arithmetic with negative numbers wrapping backwards.
export function mod(value: number, modulo: number) {
    return ((value % modulo) + modulo) % modulo;
}

export function isNum(str: string | number) {
    if (typeof str === 'number') {
        return true;
    }
    return !isNaN(parseInt(str));
}

export function toInt(value: string): number {
    return Number.parseInt(value, 10);
}
