import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 16;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        const valves = parseValves(input);
        const distances = floydWarshall(valves);
        return getMaxFlow(valves, distances);
    }

    part2(input: string, ...args: unknown[]): string | number {
        const valves = parseValves(input);
        const distances = floydWarshall(valves);
        const nonZeroValves = getNonZeroValves(valves);
        const paths = generateOpenPaths(distances, nonZeroValves, 26);
        const bestPressureFlow: { [key: string]: number } = {};

        for (const path of paths) {
            const pressureFlow = getPressureFlow(valves, distances, path, 26);
            const sorted = path.sort((a, b) => a.localeCompare(b));
            const key = sorted.join('-');
            bestPressureFlow[key] = Math.max(bestPressureFlow[key] ?? 0, pressureFlow);
        }

        const bestScores = Object.entries(bestPressureFlow)
            .map(([path, score]) => ({
                open: new Set(path.split('-')),
                score,
            }))
            .sort((a, b) => b.score - a.score);

        let max = 0;

        for (let human = 0; human < bestScores.length; human++) {
            const { open: humanOpen, score: humanScore } = bestScores[human];
            if (humanScore * 2 < max) {
                continue;
            }
            for (let elephant = 0; elephant < bestScores.length; elephant++) {
                const { open: elephantOpen, score: elephantScore } = bestScores[elephant];
                const intersect = new Set([...humanOpen].filter(x => elephantOpen.has(x)));
                const score = humanScore + elephantScore;
                if (intersect.size === 0 && score > max) {
                    max = score;
                }
            }
        }

        return max;
    }
}

type Dictionary<T> = { [key: string]: T }

function floydWarshall(valves: ValveGraph) {
    const distances: Dictionary<Dictionary<number>> = {};
    const names = Object.keys(valves);

    names.forEach((a) => {
        distances[a] = {};
        names.forEach((b) => {
            distances[a][b] = Number.MAX_SAFE_INTEGER;
        });
    });

    names.forEach((a) => {
        distances[a][a] = 0;
        valves[a].leads.forEach((b) => {
            distances[a][b] = 1;
        });
    });

    names.forEach((a) => {
        names.forEach((b) => {
            names.forEach((c) => {
                distances[b][c] = Math.min(distances[b][c], distances[b][a] + distances[a][c]);
            });
        });
    });

    return distances;
}

function getMaxFlow(valves: ValveGraph, distances: Dictionary<Dictionary<number>>) {
    const nonZeroValves = getNonZeroValves(valves);
    const paths = generateOpenPaths(distances, nonZeroValves, 30);

    let max = 0;
    for (const path of paths) {
        max = Math.max(max, getPressureFlow(valves, distances, path, 30));
    }
    return max;
}

function getNonZeroValves(valves: ValveGraph) {
    return Object.values(valves)
        .filter(({ rate }) => rate > 0)
        .map((valve) => valve.name);
}

function getPressureFlow(valves: ValveGraph, distances: Dictionary<Dictionary<number>>, path: string[], time: number) {
    let from = 'AA';
    let flow = 0;
    let timeLeft = time;

    for (const to of path) {
        const cost = distances[from][to] + 1;
        timeLeft -= cost;
        flow += valves[to].rate * timeLeft;
        from = to;
    }

    return flow;
}

function* generateOpenPaths(
    distances: Dictionary<Dictionary<number>>,
    nonZeroValves: string[],
    timeLeft: number,
    name: string = 'AA',
    open: string[] = []
): Generator<string[]> {
    for (const next of nonZeroValves) {
        const cost = distances[name][next] + 1;
        if (!open.includes(next) && cost < timeLeft) {
            open.push(next);
            yield* generateOpenPaths(distances, nonZeroValves, timeLeft - cost, next, open);
            open.pop();
        }
    }

    yield [...open];
}

function parseValves(input: string) {
    const regex = /Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.+)/;
    return input.split('\n').reduce((result, line) => {
        const match = line.match(regex);
        if (!match) {
            return result;
        }

        const name = match[1];
        const rate = Number.parseInt(match[2], 10);
        const leads = match[3].split(', ');
        result[name] = { name, rate, leads };

        return result;
    }, {} as ValveGraph);
}

type ValveGraph = { [valve: string]: Valve }

interface Valve {
    name: string;
    rate: number;
    leads: string[];
}
