# Day 20

### Part 1

I tracked all positions along the path that are a distance of 2 apart.

### Part 2

This required some refactoring. The nature of the input makes for an easier solution. There is a single path along the entire grid. 

1. Flood fill to get the distance for every step along the path. 
2. For every position on the path, calculate the manhattan distance and the distance we save from that cheat.
3. Make sure the cheat meets our criteria and add it to our count.
