import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 16;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        const time = 30;
        const valves = parseValves(input);
        const distances = getFloydWarshallDistances(valves);
        const workingValves = getWorkingValves(valves);

        let max = 0;
        for (const path of generateOpenPaths(distances, workingValves, time)) {
            max = Math.max(max, getFlow(valves, distances, path, time));
        }
        return max;
    }

    part2(input: string, ...args: unknown[]): string | number {
        const time = 26;
        const valves = parseValves(input);
        const distances = getFloydWarshallDistances(valves);
        const workingValves = getWorkingValves(valves);

        const bestFlows: { [key: string]: number } = {};
        for (const path of generateOpenPaths(distances, workingValves, time)) {
            const flow = getFlow(valves, distances, path, time);
            const sorted = path.sort((a, b) => a.localeCompare(b));
            const key = sorted.join('-');
            bestFlows[key] = Math.max(bestFlows[key] ?? 0, flow);
        }

        const bestFlowsSorted = Object.entries(bestFlows)
            .map(([path, flow]) => ({
                open: new Set(path.split('-')),
                flow,
            }))
            .sort((a, b) => b.flow - a.flow);

        let max = 0;

        for (let human = 0; human < bestFlowsSorted.length; human++) {
            const { open: humanOpen, flow: humanFlow } = bestFlowsSorted[human];
            if (humanFlow * 2 < max) {
                break;
            }

            for (let elephant = human + 1; elephant < bestFlowsSorted.length; elephant++) {
                const { open: elephantOpen, flow: elephantFlow } = bestFlowsSorted[elephant];
                const intersect = new Set([...humanOpen].filter(x => elephantOpen.has(x)));
                const score = humanFlow + elephantFlow;

                if (intersect.size === 0 && score > max) {
                    max = score;
                }
            }
        }

        return max;
    }
}

function getWorkingValves(valves: ValveGraph) {
    return Object.values(valves)
        .filter(({ rate }) => rate > 0)
        .map((valve) => valve.name);
}

function getFlow(valves: ValveGraph, distances: Dictionary<Dictionary<number>>, path: string[], time: number) {
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
    workingValves: string[],
    timeLeft: number,
    from: string = 'AA',
    open: string[] = []
): Generator<string[]> {
    for (const to of workingValves) {
        const cost = distances[from][to] + 1;
        if (!open.includes(to) && cost < timeLeft) {
            open.push(to);
            yield* generateOpenPaths(distances, workingValves, timeLeft - cost, to, open);
            open.pop();
        }
    }
    yield [...open];
}

type Dictionary<T> = { [key: string]: T }

function getFloydWarshallDistances(valves: ValveGraph) {
    const distances: Dictionary<Dictionary<number>> = {};
    const names = Object.keys(valves);

    names.forEach((a) => {
        distances[a] = {};
        distances[a][a] = 0;

        names.forEach((b) => {
            distances[a][b] = Infinity;
        });

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
