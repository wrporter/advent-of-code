import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 16;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        const valves = parseValves(input);
        const distances = floydWarshall(valves);
        return findMostPressure(valves, distances);
    }

    part2(input: string, ...args: unknown[]): string | number {
        const valves = parseValves(input);
        const distances = floydWarshall(valves);
        const nonZeroValves = Object.values(valves).filter(({ rate }) => rate > 0);
        const paths = generateOpenPaths(valves, distances, nonZeroValves, { name: 'AA', open: [], timeLeft: 26 });

        const bestPressureFlow: { [key: string]: number } = {};

        let pathCount = 0;
        for (const path of paths) {
            pathCount++;
            const sorted = path.sort((a, b) => a.localeCompare(b));
            const pressureFlow = getPressureFlow(valves, distances, path, 26);
            const key = JSON.stringify(sorted);
            bestPressureFlow[key] = Math.max(bestPressureFlow[key] ?? 0, pressureFlow);
        }

        const bestScores = Object.entries(bestPressureFlow).map(([path, score]) => {
            return { open: new Set(JSON.parse(path)), score };
        });

        console.log(bestScores.length);
        console.log(pathCount);

        let max = 0;
        // let { path: humanPath, score: humanScore } = {path: '', score: 0};
        // let { path: elephantPath, score: elephantScore } = {path: '', score: 0};
        //
        // for (let human = 0; human < bestScores.length; human++) {
        //     humanPath = bestScores[human].path;
        //     humanScore = bestScores[human].score;
        //     for (let elephant = human + 1; elephant < bestScores.length; elephant++) {
        //         elephantPath = bestScores[elephant].path;
        //         elephantScore = bestScores[elephant].score;
        //     }
        // }
        // max = Math.max(max, humanScore + elephantScore);

        for (let human = 0; human < bestScores.length; human++) {
            const { open: humanOpen, score: humanScore } = bestScores[human];
            // if (humanScore * 2 < max) {
            //     break;
            // }
            for (let elephant = 0; elephant < bestScores.length; elephant++) {
                const { open: elephantOpen, score: elephantScore } = bestScores[elephant];
                const intersect = new Set([...humanOpen].filter(x => elephantOpen.has(x)));
                if (intersect.size === 0) {
                    const score = humanScore + elephantScore;
                    max = Math.max(max, score);
                }
            }
        }

        return max;
    }
}

type FloydWarshallDistances = { [key: string]: { [key: string]: number } };

function floydWarshall(valves: ValveGraph) {
    const distances: FloydWarshallDistances = {};
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
                const dist = Math.min(distances[b][c], distances[b][a] + distances[a][c]);
                distances[b][c] = dist;
            });
        });
    });

    return distances;
}

function findMostPressure(valves: ValveGraph, distances: FloydWarshallDistances) {
    const nonZeroValves = Object.values(valves).filter(({ rate }) => rate > 0);
    const paths = generateOpenPaths(valves, distances, nonZeroValves, { name: 'AA', open: [], timeLeft: 30 });

    let max = 0;
    for (const path of paths) {
        max = Math.max(max, getPressureFlow(valves, distances, path, 30));
    }
    return max;
}

function getPressureFlow(valves: ValveGraph, distances: FloydWarshallDistances, openValvePath: string[], time: number) {
    let current = 'AA';
    let flow = 0;
    let timeLeft = time;

    for (const next of openValvePath) {
        timeLeft -= distances[current][next] + 1;
        flow += valves[next].rate * timeLeft;
        current = next;
    }

    return flow;
}

function* generateOpenPaths(valves: ValveGraph, distances: FloydWarshallDistances, nonZeroValves: Valve[], {
    name,
    open,
    timeLeft
}: Node): Generator<string[]> {
    for (const next of nonZeroValves) {
        if (!open.includes(next.name) && distances[name][next.name] <= timeLeft) {
            open.push(next.name);
            yield* generateOpenPaths(valves, distances, nonZeroValves, {
                name: next.name,
                open: open,
                timeLeft: timeLeft - distances[name][next.name] - 1
            });
            open.pop();
        }
    }

    yield [...open];
}

function parseValves(input: string) {
    const regex = /Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z]{2}(, [A-Z]{2})*)/;
    return input.split('\n').reduce((result, line) => {
        const m = regex.exec(line);
        if (!m) {
            return result;
        }

        const name = m[1];
        const rate = Number.parseInt(m[2], 10);
        const leads = m[3].split(', ');
        result[name] = { name, rate, leads };

        return result;
    }, {} as ValveGraph);
}

interface Node {
    name: string;
    open: string[];
    timeLeft: number;
}

type ValveGraph = { [valve: string]: Valve }

interface Valve {
    name: string;
    rate: number;
    leads: string[];
}
