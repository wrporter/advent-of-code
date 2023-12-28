# Day 23

## Thoughts

Another fun pathfinding problem!

### Part 1

This one can be solved with a fairly simple depth-first search.

### Part 2

You can still use your same solution from Part 1, mine took 28 minutes to get the result. I let it run while I worked on optimizing by collapsing corridors in the grid via a connected graph.

By pre-constructing a graph of the connected nodes, collapsing the paths between junctions, my runtime was reduced to 4.19 seconds. A huge improvement! I'm sure there are more optimizations I can make. I might come back to this one later to learn more.

#### Optimization progression

1. `28m` - Brute force with original grid (Part 1 solution).
2. `4.19s` - Collapse paths.
3. [35d10284](https://github.com/wrporter/advent-of-code/commit/35d10284e7777aff89051d7297da20caf86f7575) - `2.28s` - Switch edges from a map to a slice. 
4. [c5153e5c](https://github.com/wrporter/advent-of-code/commit/c5153e5cc5816679d9fe1a5cf05fb7936b5fa99c) - `1.27s` - Switch the graph from a map of points to a slice. Assign each node an index ID. This introduced some minor code complexity by keeping track of and assigning an index ID to each node when constructing the graph.
5. [3b0e71ff](https://github.com/wrporter/advent-of-code/commit/3b0e71ffb582e135d619d59ff9695717bbd42aa5) - `211ms` - Switch the seen map to a slice. This was made possible by switching the graph to a list so we could now index directly into a slice of a pre-defined size.
6. [698e1cf9](https://github.com/wrporter/advent-of-code/commit/698e1cf9c49f22f0803ed98b5c0dd09a7d110723) - `112ms` - Trim the start and end nodes from the graph and advance the start and end pointers to the first junctions. We can do this because these nodes only have a single edge connection. Trimming them from the graph greatly reduces the backtracking required in the depth-first search.
7. [8fb1b382](https://github.com/wrporter/advent-of-code/commit/8fb1b382d0dcbc886cad83c95311dd589601568c) - `105ms` - Use a bitset for tracking visited nodes. We can do this because we have less than 64 nodes, and we specified an ID from 0 to 36 (less for the sample input).
8. [b1de2e0e](https://github.com/wrporter/advent-of-code/commit/b1de2e0e8249f29ccb995f83dba4c015c8bb0e6c)- `97ms` - Use an iterative DFS over recursive. We no longer need backtracking and have no need for a recursive solution. However, we need to make sure that when we reach the goal node we need to stop execution down that path, otherwise the runtime is double.
9. [48d11425](https://github.com/wrporter/advent-of-code/commit/48d114259d039fcd59baecdc4bf0e70e2628d52f)- `5.9ms` - Prune the search space by keeping track of how much distance remains from the current node. If that total distance is below the longest distance we've seen so far, we can exit that branch early. 
