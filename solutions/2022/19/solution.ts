import { AbstractSolution } from '~/solution';

export class Solution extends AbstractSolution {
    year = 2022;
    day = 19;
    filename = 'input.txt';

    part1(input: string, ...args: unknown[]): string | number {
        const blueprints = parseBlueprints(input);
        let qualityLevel = 0;

        for (const blueprint of blueprints) {
            const maxGeodes = findMaxGeodes(blueprint, 24);
            qualityLevel += maxGeodes * blueprint.id;
        }

        return qualityLevel;
    }

    part2(input: string, ...args: unknown[]): string | number {
        const blueprints = parseBlueprints(input).slice(0, 3);
        let product = 1;

        for (const blueprint of blueprints) {
            const maxGeodes = findMaxGeodes(blueprint, 32);
            product *= maxGeodes;
        }

        return product;
    }
}

function findMaxGeodes(blueprint: Blueprint, maxTime: number) {
    const maxCost = getMaxCost(blueprint);
    const { cost } = blueprint;

    let maxGeodes = 0;
    const seen: { [key: string]: boolean } = {};
    const states: Node[] = [{
        resources: { ore: 0, clay: 0, obsidian: 0, geode: 0 },
        robots: { ore: 1, clay: 0, obsidian: 0, geode: 0 },
        timeLeft: maxTime,
    }];

    while (states.length > 0) {
        const node = states.pop() as Node;
        let { resources, robots, timeLeft } = node;

        maxGeodes = Math.max(maxGeodes, resources.geode);
        if (timeLeft === 0) {
            continue;
        }

        const current: Node = {
            resources: {
                ore: getPotentialResource(resources, robots, maxCost.ore, timeLeft, 'ore'),
                clay: getPotentialResource(resources, robots, cost.obsidian.clay, timeLeft, 'clay'),
                obsidian: getPotentialResource(resources, robots, cost.geode.obsidian, timeLeft, 'obsidian'),
                geode: resources.geode,
            },
            robots: {
                ore: Math.min(robots.ore, maxCost.ore),
                clay: Math.min(robots.clay, cost.obsidian.clay),
                obsidian: robots.obsidian,
                geode: Math.min(robots.geode, cost.geode.obsidian),
            },
            timeLeft: timeLeft - 1,
        };

        const key = keyOf(current);
        if (seen[key]) {
            continue;
        }
        seen[key] = true;

        const next = clone(current, {
            ore: current.robots.ore,
            clay: current.robots.clay,
            obsidian: current.robots.obsidian,
            geode: current.robots.geode
        });

        states.push(next);

        if (current.resources.ore >= cost.geode.ore && current.resources.obsidian >= cost.geode.obsidian) {
            states.push(clone(next, {
                ore: -cost.geode.ore,
                obsidian: -cost.geode.obsidian,
            }, {
                geode: 1,
            }));
        } else if (current.resources.ore >= cost.obsidian.ore && current.resources.clay >= cost.obsidian.clay) {
            states.push(clone(next, {
                ore: -cost.obsidian.ore,
                clay: -cost.obsidian.clay,
            }, {
                obsidian: 1,
            }));
        } else {
            if (current.resources.ore >= cost.clay.ore) {
                states.push(clone(next, {
                    ore: -cost.clay.ore,
                }, {
                    clay: 1,
                }));
            }
            if (current.resources.ore >= cost.ore.ore) {
                states.push(clone(next, {
                    ore: -cost.ore.ore,
                }, {
                    ore: 1,
                }));
            }
        }
    }

    return maxGeodes;
}

function getPotentialResource(resources: Minerals, robots: Minerals, cost: number, timeLeft: number, mineral: Mineral) {
    return Math.min(resources[mineral], timeLeft * cost - robots[mineral] * (timeLeft - 1));
}

function keyOf({ resources, robots, timeLeft }: Node) {
    return `${resources.ore},${resources.clay},${resources.obsidian},${resources.geode}-${robots.ore},${robots.clay},${robots.obsidian},${robots.geode}-${timeLeft}`;
}

function clone({ resources, robots, timeLeft }: Node, resourceDiff: Partial<Minerals>, robotDiff?: Partial<Minerals>) {
    const copy = {
        resources: { ...resources },
        robots: { ...robots },
        timeLeft: timeLeft,
    };

    for (const mineral of MINERALS) {
        copy.resources[mineral] += resourceDiff[mineral] ?? 0;
        copy.robots[mineral] += robotDiff?.[mineral] ?? 0;
    }

    return copy;
}

type Mineral = 'ore' | 'clay' | 'obsidian' | 'geode';

const MINERALS: Mineral[] = [
    'ore',
    'clay',
    'obsidian',
    'geode',
];

function getMaxCost(blueprint: Blueprint): Minerals {
    const maxRobots: Minerals = { ore: 0, clay: 0, obsidian: 0, geode: 0 };
    for (const robot of MINERALS) {
        for (const mineral of MINERALS) {
            maxRobots[robot] = Math.max(maxRobots[robot], blueprint.cost[mineral][robot]);
        }
    }
    return maxRobots;
}

interface Node {
    resources: Minerals;
    robots: Minerals;
    timeLeft: number;
}

interface Minerals {
    ore: number;
    clay: number;
    obsidian: number;
    geode: number;
}

function parseBlueprints(input: string) {
    const regex = /Blueprint (\d+): Each ore robot costs (\d+) ore\. Each clay robot costs (\d+) ore\. Each obsidian robot costs (\d+) ore and (\d+) clay\. Each geode robot costs (\d+) ore and (\d+) obsidian\./;
    return input.split('\n').map((line) => {
        const match = line.match(regex) as RegExpMatchArray;
        const blueprint: Blueprint = {
            id: toInt(match[1]),
            cost: {
                ore: { ore: toInt(match[2]), clay: 0, obsidian: 0, geode: 0 },
                clay: { ore: toInt(match[3]), clay: 0, obsidian: 0, geode: 0 },
                obsidian: { ore: toInt(match[4]), clay: toInt(match[5]), obsidian: 0, geode: 0 },
                geode: { ore: toInt(match[6]), clay: 0, obsidian: toInt(match[7]), geode: 0 },
            },
        };
        return blueprint;
    });
}

interface Blueprint {
    id: number;
    cost: {
        ore: Minerals,
        clay: Minerals,
        obsidian: Minerals,
        geode: Minerals,
    };
}

function toInt(value: string): number {
    return Number.parseInt(value, 10);
}
