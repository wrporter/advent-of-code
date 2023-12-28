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
3. 35d10284e7777aff89051d7297da20caf86f7575 - `2.28s` - Switch edges from a map to a slice. 
4. c5153e5cc5816679d9fe1a5cf05fb7936b5fa99c - `1.27s` - Switch the graph from a map of points to a slice. Assign each node an index ID. This introduced some minor code complexity by keeping track of and assigning an index ID to each node when constructing the graph.
5. 3b0e71ffb582e135d619d59ff9695717bbd42aa5 - `211ms` - Switch the seen map to a slice. This was made possible by switching the graph to a list so we could now index directly into a slice of a pre-defined size.
6. 698e1cf9c49f22f0803ed98b5c0dd09a7d110723 - `112ms` - Trim the start and end nodes from the graph and advance the start and end pointers to the first junctions. We can do this because these nodes only have a single edge connection. Trimming them from the graph greatly reduces the backtracking required in the depth-first search.
7. 8fb1b382d0dcbc886cad83c95311dd589601568c - `105ms` - Use a bitset for tracking visited nodes. We can do this because we have less than 64 nodes, and we specified an ID from 0 to 36 (less for the sample input).

Changes that did not work:

1. I figured that an iterative DFS would be more performant than a recursive one. However, this is not the case. I wonder if the Go compiler is optimizing the function call, particularly since all of the arguments are straight integers now that we use a bitset to track visited nodes. Whereas, the iterative solution is growing and contracting a slice. The iterative solution is `170ms` compared to the `105ms` recursive solution.
