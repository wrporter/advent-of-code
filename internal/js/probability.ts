export function combinations(values: string[], startSize: number, endSize: number) {
    const result: string[][] = [];

    function combo(values: string[], size: number) {
        const s = new Array(size);
        const last = size - 1;

        const recurse = function (start: number, next: number) {
            for (let current = 0; current < values.length; current++) {
                const value = values[current];
                s[start] = value;

                if (start === last) {
                    result.push([...s]);
                } else {
                    recurse(start + 1, current + 1);
                }
            }
        };

        recurse(0, 0);
    }

    for (let i = startSize; i <= endSize; i++) {
        combo(values, i);
    }

    return result;
}

/**
 * Permutations without repeating elements. Permutations are where order matters. So 'ab' is not equivalent to 'ba'.
 * @param list
 * @param min
 * @param max
 */
export function permutations<T>(list: T[], min: number = 1, max: number = list.length) {
    let result: T[][] = [];
    for (let size = min; size <= max; size++) {
        result = result.concat(permute(list, size));
    }
    return result;
}

/**
 * Take from https://stackoverflow.com/questions/59369727/permutation-with-max-length-parameter
 * @param list
 * @param size
 */
function permute<T>(list: T[], size: number = list.length) {
    const result: T[][] = [];

    if (size === 0) {
        return [[]];
    }

    for (let i = 0; i < list.length; i++) {
        const copy = [...list];
        const head = copy.splice(i, 1);
        const rest = permute(copy, size - 1);

        for (let j = 0; j < rest.length; j++) {
            const next = head.concat(rest[j]);
            result.push(next);
        }
    }

    return result;
}
