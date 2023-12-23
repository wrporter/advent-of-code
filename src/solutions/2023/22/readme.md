# Day 22

## Thoughts

3D Tetris! I also got some Jenga vibes. This day was very similar to Day 14, but with a third dimension - `z`. 

### Part 1

1. Parse the brick coordinate ranges. Sort the bricks by their starting `z` value so we can easily drop the bricks one at a time. Much like when we rolled the round rocks on Day 14.
2. Drop the bricks that are still floating in the air.
   1. Keep track of the peaks, or tallest `z` positions.
   2. Loop through each brick.
      1. Keep track of the `x, y` area the brick occupies and the tallest peak it will rest on.
      2. For every `x, y` spot, determine the highest peak the brick will fall on.
      3. Loop through the `x, y` area of the brick and update the peaks with the height of the current brick.
      4. Calculate how much the brick must fall by its starting `z` minus the highest peak it will rest on.
      5. Update the brick's `z` coordinates.
3. For each brick, remove it and drop the bricks. If none have fallen, we know that brick is stable. This is a very inefficient way to do this and we can instead create a graph of connected bricks, but this is a much simpler approach.

### Part 2

The same as Part 1, but we now sum all the bricks that will fall for each brick that we might disintegrate.
