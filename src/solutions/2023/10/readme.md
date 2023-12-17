# Day 10

## Thoughts

Alright, the first very challenging problem this year. I woke up sick today, so didn't come back to Part 2 for a few hours while I napped. 

### Part 1

I felt like this part was pretty straightforward, but required a lot of logic to determine which direction we should be moving, based on how we enter each pipe. I started with a massive if-else chain and later converted that same logic into a map to clean it up. My map takes in the pipe character and the previous direction where we came from. The map outputs a pipe with the next and previous directions to travel in.

1. Parse the input. I initially parsed this as a single array of strings, but debugging made it hard to view runes/bytes as the actual characters, so I switched to a 2D grid of strings. Because we travel through the whole grid, we might as well take note of the start position.
2. Determine the valid pipe for the start. This assumes other "junk" pipes do not intend to connect to the start. I don't think any of the puzzle inputs would do that because you would need to do quite a bit more work to travel through other potential paths, determine whether it loops, then pick the right path.
   1. Loop through each pipe configuration.
   2. Verify that the next and previous positions are valid.
   3. Get the next available pipes and make sure they are valid.
   4. Finally, if both the pipes can travel back to the start in the direction they came from, then we know we found the start pipe!
3. Because of our pipe-direction map, we can now easily loop through and move the current point in the pipe's next direction and count the number of steps until we reach the start again.
4. Because we are dealing with a loop, we can return half the step count to get to the furthest tile.

### Part 2

This part was far more tricky, requiring some understanding of floodfill algorithms. It would be far simpler if we didn't have to deal with "squeezing" through pipes, allowing us to do a naive recursive algorithm to find all the tiles inside the loop. The algorithm would require us to keep track of whether we hit an outside edge, if we do, we know that we are not inside the loop. Due to "squeezing", I considered the following ideas:

1. Modify the floodfill algorithm to check if we've hit a "corridor" where we can "squeeze" through pipes. For example, if we hit a `7F` next to each other, then we could follow the pipes until we go outside by hitting an outside edge or stay inside by hitting `LJ`. However, there is more to this approach, because squeezing can still take turns. For example:
    ```
    S----7
    |....|
    |....|
    L-7F-J
    ..||..
    ..|L-7
    ..L--J
    ```
    I figured this approach would be terribly complicated to code for, with many cases. So, I didn't even try to code it after planning it out on paper.
2. Expand our grid by twice the size and insert "ground" between squeezed pipes. Then perform a floodfill algorithm like normal and subtract all the "net new" ground tiles. I didn't try coding this approach because I figured it would be difficult to properly subtract the new tiles and add the correct pipe tiles in the new spaces.
3. The final approach I considered was to Google for "algorithm to find point inside shape". I found this [StackOverflow post](https://stackoverflow.com/questions/217578/how-can-i-determine-whether-a-2d-point-is-within-a-polygon) that explained the **ray casting** algorithm by drawing a line through the grid and counting intersections we pass through over the loop. If we have passed through an odd number of intersections, we know the tile is inside the loop, otherwise it's outside.
   1. I attempted to perform this algorithm going left-to-right, then top-to-bottom, but they don't tell the full story of being able to squeeze through pipes.
   2. I then attempted to use the logic in some of the answers, but they did not work and there wasn't enough explanation about the equations to debug properly.
   3. I then wondered if I could perform the ray casting in a diagonal since we are dealing with perfect rectangles and not a jagged polygon. A trick to this was determining which pipes to include as intersections. I excluded `L` and `7` while drawing rays from top-left to bottom-right because they would cause us to skip over the inside of the loop and back to the outside. This approach worked!

My final algorithm is as follows.

1. Parse the input as before.
2. Get the start pipe as before. Set the correct pipe in the grid at the start. We use this during the ray-casting algorithm to properly detect intersections.
3. Follow the loop as before, and keep track of the points in a map for quick lookups.
4. For every point on the top and left edges of the grid, draw a ray (line) to the bottom-right.
   1. Every time we cross into the loop, increment our intersection counter.
   2. If we are at a spot that's not on the loop and our intersection counter is odd, we know we're inside. This works for squeezing between pipes because squeezing requires the pipes to loop back on each other in another location so we'll get to an even number of intersections in those cases.
