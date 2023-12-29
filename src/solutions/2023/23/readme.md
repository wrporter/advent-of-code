# Day 23

## Thoughts

Another fun pathfinding problem!

### Part 1

This one can be solved with a fairly simple depth-first search.

1. Use a recursive function to perform the depth-first search.
2. Initialize the longest path to 0.
3. If the current node is the goal, return the number of steps so far.
4. Loop through each neighboring position.
   1. Skip if we've already seen this node.
   2. Add the node to the seen set.
   3. Recurse only if the next position is a `.` or the slope is going downward (e.g. `<` and we are traveling left).
   4. Update our longest path if the recursed path is longer.
   5. Remove the node from the seen set.
5. Return the longest path that was found.

### Part 2

The naive solution for Part 1 may still be used, but with terrible timing. Mine took 28 minutes to get the result. Non-compiled languages would likely take hours. I let it run while I worked on optimizing by collapsing corridors in the grid via a connected graph. I didn't initially make all the optimizations other than this first one. A few days after Christmas I came back and went through a series of optimizations that are described in detail below, with their associated code diffs. Overall, here's my compiled process:

1. Collapse all corridors in the grid and keep only junctions as nodes.
   1. Initialize a slice for the graph. Nodes will be assigned an index ID.
   2. Initialize a map to keep track of points to their associated index IDs.
   3. Perform a breadth-first search from the starting node to explore all nodes. Keep track of which ones are visited so we don't revisit any nodes.
      1. Loop through each edge. If the edge only has 1 other neighbor than the previous node, increase the distance by 1 and continue looping.
      2. If the next node does not yet exist in the graph, set its ID to the size of the graph then append it to the graph.
      3. Add an edge from the current node to the next node and vice-versa.
      4. Add the next node to the point-ID map.
      5. Add the previous position to the visited set so nodes don't attempt to travel back where they came from.
2. Trim the start and end nodes from the graph.
   1. Keep track of the distance from the start and end nodes to their only edge.
   2. Advance the start and end nodes to their edge node and remove the edge back to them.
3. Perform an iterative depth-first search to find the longest path.
   1. Initialize the total remaining distance from the start to the end based on all possible routes (edge distances).
   2. Initialize a stack, keeping track of the current node ID, distance so far, visited bitset, and the remaining distance.
   3. Loop through each node in the stack.
      1. Pop the next value off the stack.
      2. If we've reached the end, update the longest path with the current distance.
      3. Initialize the remaining distance by subtracting all edges that have not been visited.
      4. Loop through each edge.
         1. Only add the next node to the stack if it hasn't been visited and the remaining possible distance is longer than the currently known longest path.

#### Optimization progression

1. `28m` - Brute force with original grid (Part 1 solution).
2. `4.19s` - Pre-construct a collapsed graph of junction nodes and associated edge distance. Positions with only 2 neighbors only go forward or backward to an already visited node. We can collapse these until we see a node with more neighbors where a decision is important. The sample input is reduced to 9 total nodes and my real input to 36 nodes. Here's what the sample input looks like where the start node is 15 steps away from the next node.
 ```mermaid
 graph LR
   0["`(1, 0)`"]-- 15 ---1["`(3, 5)`"]
   1["`(3, 5)`"]-- 22 ---2["`(11, 3)`"]
   1["`(3, 5)`"]-- 22 ---3["`(5, 13)`"]
   2["`(11, 3)`"]-- 30 ---4["`(21, 11)`"]
   2["`(11, 3)`"]-- 24 ---5["`(13, 13)`"]
   3["`(5, 13)`"]-- 12 ---5["`(13, 13)`"]
   3["`(5, 13)`"]-- 38 ---6["`(13, 19)`"]
   4["`(21, 11)`"]-- 10 ---7["`(19, 19)`"]
   4["`(21, 11)`"]-- 18 ---5["`(13, 13)`"]
   5["`(13, 13)`"]-- 10 ---6["`(13, 19)`"]
   6["`(13, 19)`"]-- 10 ---7["`(19, 19)`"]
   7["`(19, 19)`"]-- 5 ---8["`(21, 22)`"]
 ```
and here's my real input:
 ```mermaid
 graph LR
   0["`(1, 0)`"]-- 101 ---1["`(13, 9)`"]
   1["`(13, 9)`"]-- 100 ---2["`(29, 5)`"]
   1["`(13, 9)`"]-- 280 ---3["`(15, 43)`"]
   2["`(29, 5)`"]-- 194 ---4["`(57, 7)`"]
   2["`(29, 5)`"]-- 174 ---5["`(31, 41)`"]
   3["`(15, 43)`"]-- 62 ---5["`(31, 41)`"]
   3["`(15, 43)`"]-- 184 ---6["`(9, 61)`"]
   4["`(57, 7)`"]-- 172 ---7["`(83, 13)`"]
   4["`(57, 7)`"]-- 152 ---8["`(67, 33)`"]
   5["`(31, 41)`"]-- 232 ---8["`(67, 33)`"]
   5["`(31, 41)`"]-- 172 ---9["`(35, 65)`"]
   6["`(9, 61)`"]-- 114 ---9["`(35, 65)`"]
   6["`(9, 61)`"]-- 230 ---10["`(19, 89)`"]
   7["`(83, 13)`"]-- 102 ---11["`(99, 11)`"]
   7["`(83, 13)`"]-- 108 ---12["`(79, 33)`"]
   8["`(67, 33)`"]-- 48 ---12["`(79, 33)`"]
   8["`(67, 33)`"]-- 180 ---13["`(61, 59)`"]
   9["`(35, 65)`"]-- 156 ---13["`(61, 59)`"]
   9["`(35, 65)`"]-- 110 ---14["`(35, 87)`"]
   10["`(19, 89)`"]-- 86 ---14["`(35, 87)`"]
   10["`(19, 89)`"]-- 194 ---15["`(5, 101)`"]
   11["`(99, 11)`"]-- 528 ---16["`(133, 37)`"]
   11["`(99, 11)`"]-- 164 ---17["`(105, 37)`"]
   12["`(79, 33)`"]-- 166 ---17["`(105, 37)`"]
   12["`(79, 33)`"]-- 154 ---18["`(89, 57)`"]
   13["`(61, 59)`"]-- 138 ---18["`(89, 57)`"]
   13["`(61, 59)`"]-- 168 ---19["`(61, 83)`"]
   14["`(35, 87)`"]-- 130 ---19["`(61, 83)`"]
   14["`(35, 87)`"]-- 94 ---20["`(43, 105)`"]
   15["`(5, 101)`"]-- 182 ---20["`(43, 105)`"]
   15["`(5, 101)`"]-- 402 ---21["`(33, 127)`"]
   16["`(133, 37)`"]-- 240 ---22["`(123, 63)`"]
   16["`(133, 37)`"]-- 112 ---17["`(105, 37)`"]
   17["`(105, 37)`"]-- 112 ---23["`(105, 57)`"]
   18["`(89, 57)`"]-- 104 ---23["`(105, 57)`"]
   18["`(89, 57)`"]-- 194 ---24["`(75, 89)`"]
   19["`(61, 83)`"]-- 76 ---24["`(75, 89)`"]
   19["`(61, 83)`"]-- 84 ---25["`(65, 103)`"]
   20["`(43, 105)`"]-- 88 ---25["`(65, 103)`"]
   20["`(43, 105)`"]-- 204 ---21["`(33, 127)`"]
   21["`(33, 127)`"]-- 266 ---26["`(61, 129)`"]
   22["`(123, 63)`"]-- 206 ---27["`(129, 83)`"]
   22["`(123, 63)`"]-- 88 ---23["`(105, 57)`"]
   23["`(105, 57)`"]-- 78 ---28["`(113, 79)`"]
   24["`(75, 89)`"]-- 320 ---28["`(113, 79)`"]
   24["`(75, 89)`"]-- 76 ---29["`(85, 103)`"]
   25["`(65, 103)`"]-- 80 ---29["`(85, 103)`"]
   25["`(65, 103)`"]-- 166 ---26["`(61, 129)`"]
   26["`(61, 129)`"]-- 72 ---30["`(79, 123)`"]
   27["`(129, 83)`"]-- 116 ---31["`(137, 103)`"]
   27["`(129, 83)`"]-- 60 ---28["`(113, 79)`"]
   28["`(113, 79)`"]-- 140 ---32["`(109, 107)`"]
   29["`(85, 103)`"]-- 136 ---32["`(109, 107)`"]
   29["`(85, 103)`"]-- 130 ---30["`(79, 123)`"]
   30["`(79, 123)`"]-- 298 ---33["`(107, 129)`"]
   31["`(137, 103)`"]-- 104 ---34["`(131, 125)`"]
   31["`(137, 103)`"]-- 164 ---32["`(109, 107)`"]
   32["`(109, 107)`"]-- 104 ---33["`(107, 129)`"]
   33["`(107, 129)`"]-- 188 ---34["`(131, 125)`"]
   34["`(131, 125)`"]-- 147 ---35["`(139, 140)`"]
 ```
3. [35d10284](https://github.com/wrporter/advent-of-code/commit/35d10284e7777aff89051d7297da20caf86f7575) - `2.28s` - Switch edges from a map to a slice. Maps have a constant time $O(1)$ lookup, but the hashing function is far more expensive than indexing into an array/slice.
4. [c5153e5c](https://github.com/wrporter/advent-of-code/commit/c5153e5cc5816679d9fe1a5cf05fb7936b5fa99c) - `1.27s` - Switch the graph from a map of points to a slice. Assign each node an index ID. This introduced some minor code complexity by keeping track of and assigning an index ID to each node when constructing the graph.
5. [3b0e71ff](https://github.com/wrporter/advent-of-code/commit/3b0e71ffb582e135d619d59ff9695717bbd42aa5) - `211ms` - Switch the seen map to a slice. This was made possible by switching the graph to a list so we could now index directly into a slice of a pre-defined size.
6. [698e1cf9](https://github.com/wrporter/advent-of-code/commit/698e1cf9c49f22f0803ed98b5c0dd09a7d110723) - `112ms` - Trim the start and end nodes from the graph and advance the start and end pointers to the first junctions. We can do this because these nodes only have a single edge connection. Trimming them from the graph greatly reduces the backtracking required in the depth-first search.
7. [8fb1b382](https://github.com/wrporter/advent-of-code/commit/8fb1b382d0dcbc886cad83c95311dd589601568c) - `105ms` - Use a bitset for tracking visited nodes. We can do this because we have less than 64 nodes, and we specified an ID from 0 to 36 (less for the sample input).
8. [b1de2e0e](https://github.com/wrporter/advent-of-code/commit/b1de2e0e8249f29ccb995f83dba4c015c8bb0e6c)- `97ms` - Use an iterative DFS over recursive. We no longer need backtracking and have no need for a recursive solution. However, we need to make sure that when we reach the goal node we need to stop execution down that path, otherwise the runtime is double.
9. [48d11425](https://github.com/wrporter/advent-of-code/commit/48d114259d039fcd59baecdc4bf0e70e2628d52f)- `5.9ms` - Prune the search space by keeping track of how much distance remains from the current node. If that total distance is below the longest distance we've seen so far, we can exit that branch early. 
