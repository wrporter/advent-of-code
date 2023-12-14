# Day 14

## Thoughts

I woke up fairly groggy and kept making silly mistakes that I spent a lot of time on debugging. Overall, very fun problem. Another cycle detection, like we've seen in years past.

### Part 1

This one is fairly straightforward. I kept wondering if there was a more optimal way to just look at rocks rather than traversing the entire grid to move them. But that would require a sorted map. I love how java has its `SortedMap` that makes this kind of thing really nice. This kind of thing would require a lot more code, having to track the coordinates of all the cube-shaped and rounded rocks. Here's what I did:

1. Parse the grid into a 2D array of characters.
2. Loop through each column to roll the rounded rocks.
   1. Keep track of the first empty position. 
   2. Loop through each row.
      1. If we encounter a rounded rock `O`, set its current position to empty `.`. Move the rounded rock to our first the empty position and increase the empty position by `1` (right above the rounded rock we just moved).
      2. If we encounter a cube-shaped rock `#`, update the empty position by the rock's position plus `1` (right above that rock).
3. Calculate the north load by looping through each row and column for each rounded rock `O`. We need to flip the load from our row, so this can be done by subtracting the size of the grid by the current row.

### Part 2

It's obvious from the beginning that this one will need to detect a cycle and jump our processing forward. The word `cycle` is constantly repeated and we see that we need to cycle 1 billion times, too much to brute force.

I kept deliberating about ways to optimize, but all my thoughts about rolling rocks in a specific direction kept ending in disaster, with a lot more code than I'd want. I settled on a slightly slower solution today.

1. Parse the grid as before.
2. Keep track of which grids configurations we've already seen, with their associated cycle.
3. Loop through each cycle.
   1. Loop through each tilt direction.
      1. Roll the rocks as before.
      2. Rotate the grid 90 degrees. I used a function I copied from somewhere on the internets last year, but then found a more optimal solution later because copying the grid every time seemed like such a waste of computing resources.
   2. Encode the grid back to a string so we can use it as a key in our "seen" map.
   3. If we end up with a grid we've already seen, we know we've hit a cycle! Find the cycle length (or period) by subtracting the current cycle by the start of the cycle (defined by what was in our map). Determine how many cycles we have remaining by subtracting the total number of cycles by the current cycle. Then jump to the cycle where we are no longer in a cycle.
4. Calculate the load as before.
