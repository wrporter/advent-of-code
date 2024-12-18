# Day 18

### Part 1

What a relief from previous days! This one can be knocked out with a simple Breadth-First Search.

### Part 2

The only addition to this that really has to be done is to try our BFS solution from Part 1 after every byte falls. However, it takes 1.7 seconds to run on my machine, so I'll come back and optimize this.

#### Optimizations

1. `1.7s` - Use maps with expensive lookup to track byte and visited positions.
2. [f0dacdbd](https://github.com/wrporter/advent-of-code/commit/f0dacdbdff3a30a687954e2ba92b6cfd7aabbe0d) - `150ms` - Use 2D grids instead of a map to track byte and visited positions. Maps have a constant time $O(1)$ lookup, but the hashing function is far more expensive than indexing into an array/slice.
3. [619dd184](https://github.com/wrporter/advent-of-code/commit/619dd18472301a8da05e709085da7c7a3ac533eb) - `135ms` - Use 1D arrays instead of a grid to track byte and visited positions. This barely reduces setup and lookup time.
