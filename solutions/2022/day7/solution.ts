import { AbstractSolution } from '~/solution';

const TOTAL_SPACE = 70_000_000;
const UPDATE_SPACE = 30_000_000;

export class Solution extends AbstractSolution {
    year = 2022;
    day = 7;
    filename = 'input.txt';

    part1(input: string): string | number {
        const root = this.parseInput(input);

        const sum = (node: Node) => {
            let size = 0;
            if (node.size <= 100_000) {
                size += node.size;
            }
            Object.values(node.children).forEach((child) => {
                if (child.type === NodeType.Directory) {
                    size += sum(child);
                }
            });
            return size;
        };

        return sum(root);
    }

    part2(input: string): string | number {
        const root = this.parseInput(input);
        const spaceOnDisk = TOTAL_SPACE - root.size;
        const requiredSpace = UPDATE_SPACE - spaceOnDisk;

        const getSmallest = (node: Node) => {
            let smallest = node.size;

            Object.values(node.children)
                .filter((child) => child.type === NodeType.Directory)
                .forEach((child) => {
                    const space = getSmallest(child);
                    if (space >= requiredSpace) {
                        smallest = Math.min(smallest, space);
                    }
                });

            return smallest;
        };

        return getSmallest(root);
    }

    private parseInput(input: string) {
        const lines = input.split('\n').slice(1);
        const root = new Node('/');
        let cwd = root;

        for (let i = 0; i < lines.length;) {
            const [_, command, ...args] = lines[i].split(' ');

            if (command === 'ls') {
                let fileIndex = i + 1;
                for (; fileIndex < lines.length && !lines[fileIndex].startsWith('$'); fileIndex++) {
                    const sections = lines[fileIndex].split(' ');
                    if (sections[0] === 'dir') {
                        cwd.add(new Node(sections[1], cwd));
                    } else {
                        const size = Number.parseInt(sections[0], 10);
                        cwd.add(new Node(sections[1], cwd, NodeType.File, size));
                    }
                }
                i = fileIndex;
            } else if (command === 'cd') {
                const [dir] = args;
                if (dir == '..') {
                    cwd = cwd.getParent();
                } else {
                    cwd = cwd.get(dir);
                }
                i++;
            }
        }

        root.populateDirectorySizes();
        return root;
    }
}

enum NodeType {
    Directory,
    File,
}

class Node {
    public readonly parent: Node;

    constructor(
        public name: string,
        parent: Node | null = null,
        public type: NodeType = NodeType.Directory,
        public size: number = 0,
        public readonly children: { [name: string]: Node } = {},
    ) {
        this.parent = this;
        if (parent) {
            this.parent = parent;
        }
        if (type === NodeType.Directory) {
            this.children = {};
        }
    }

    add(node: Node) {
        this.children[node.name] = node;
    }

    get(child: string) {
        return this.children[child];
    }

    getParent() {
        return this.parent;
    }

    populateDirectorySizes() {
        Object.values(this.children).forEach((child) => {
            this.size += child.populateDirectorySizes();
        });
        return this.size;
    }
}
