# Day 16

## Thoughts

I really enjoyed this one. After reading the description, I already had an idea of how to solve the problem, but I kept making so many little mistakes along the way! I knew from the beginning that we'd have to determine when to stop the arrows if there were infinite loops between sets of mirrors and would need to include the direction (using a vector rather than just a point). I kept debugging and once I got to printing out the grid, I already had my beams going in all the right directions, but wasn't counting based on uniquely energized tiles.

### Part 1

1. Parse the input into a 2D grid of characters.
2. Create a list of beams with an initial one starting at `(0, 0)` with its direction heading to the **right**.
3. Create a set of vectors for the beam positions we've already seen.
4. Loop while there are still beams. And loop for each beam.
   1. Remove the first beam from the list.
   2. If we've already visited this beam or it has gone out of bounds, continue.
   3. Add the beam to the seen set.
   4. Update the beam direction if it hits a mirror.
   5. Split the beam if it hits a splitter.
   6. If the beam was not split, move it and addit to the list of beams.
5. Count the energized tiles. This can be done by creating a new set of points (rather than vectors) and adding all the seen vector's points to it, then simply returning the size of that new set.

### Part 2

I felt so relieved that this one could be brute-forced. It can probably be optimized by memoizing the energized tile counts for each cell and direction (vector). I might come back to optimize later.

I refactored my Part 1 code into a separate function and looked for the max value for each side of the grid.
