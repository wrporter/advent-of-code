# Day 17

## Thoughts

I was worried because I don't have the best tools in Go for a path traversal problem and this was obviously one. But searching through my past solutions, I found 2021, day 15 had a priority queue implementation I could use! That was the crux of this whole thing.

### Part 1

We can't simply do a BFS or DFS solution. Those will be way too slow due to the number of possibilities. So we need to reach for Dijkstra or A*. I don't think I've ever actually implemented A*, but I'm pretty familiar with Dijkstra from past AoC events. It's a modified DFS, but you include the cost of each step (weighted graph traversal). You need a priority queue for this.

1. Create the queue with the starting position and possible directions (right and down).
2. Create a set of vectors we've seen.
3. Create a map of the cost of each vector. This way, we don't have to recalculate costs we've already seen.
4. Loop until the queue is empty (meaning we didn't find a solution).
   1. Pop off the next node to evaluate.
   2. If it's at our goal, we've seen our lowest value so far! Let's return its cost (total heat loss).
   3. If we've seen this vector before, let's skip it.
   4. Loop for each next direction (turn right or left).
      1. Set our current cost increase of moving to the next position to `0`.
      2. Loop through each possible distance (1-3).
         1. If the next coordinate is outside the grid, break out of the loop.
         2. Increase the cost of the next step.
         3. If there is a cost we've already seen from our current position (and direction) that is better than our current cost, skip this possibility.
         4. Add the cost of the next step to our cost map.
         5. Add the next node to our priority queue.

### Part 2

We just need to make a slight modification to our algorithm. Prior to adding the next node to the queue, only do so if our current distance is greater than the minimum distance the crucible can travel.
