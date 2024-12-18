# Day 18

### Part 1

What a relief from previous days! This one can be knocked out with a simple Breadth-First Search.

### Part 2

The only addition to this that really has to be done is to try our BFS solution from Part 1 after every byte falls. However, it takes 1.7 seconds to run on my machine, so I'll come back and optimize this.

#### Optimizations

1. `1.7s` - Use maps with `O(log(n))` lookup to track byte and visited positions.
2. `150ms` - Use 2D grids instead of a map to track byte and visited positions. Grid indexing is `O(1)`, whereas map lookup is `O(log(n))`.
