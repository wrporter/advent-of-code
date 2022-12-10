export function gridToString(grid: string[][]) {
    let delimiter = '\n';
    return grid.reduce((result, row, index) => {
        if (index === grid.length - 1) {
            delimiter = '';
        }
        return result + row.join('') + delimiter;
    }, '');
}