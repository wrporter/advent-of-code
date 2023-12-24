# Day 23

## Thoughts

Another fun pathfinding problem!

### Part 1

This one can be solved with a fairly simple depth-first search.

### Part 2

You can still use your same solution from Part 1, mine took 28 minutes to get the result. I let it run while I worked on optimizing by collapsing corridors in the grid via a connected graph.

By pre-constructing a graph of the connected nodes, collapsing the paths between junctions, my runtime was reduced to 4.1 seconds. A huge improvement! I'm sure there are more optimizations I can make. I might come back to this one later to learn more.
