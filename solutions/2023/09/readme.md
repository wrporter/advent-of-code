# Day 9

## Thoughts

This one totally looked like it was going to require some polynomial math, which I don't remember or know much about anymore. But luckily, we could compute it just by following instructions and didn't need a math trick for performance. I was able to use most of my Part 1 code in Part 2 which was really nice yet again.

Recursion was very helpful here because it allowed us to start calculating the extrapolated value at the bottom of the history and work our way back up. If we did this with for loops, we'd have to collect the full history first, then start at the bottom and add or subtract the overall value.

### Part 1

1. Parse the input by creating a table of sequences.
2. Sum the extrapolated values.
3. To extrapolate the values, I wrote a recursive function. The base case is when all the values are `0` in the sequence. Just return the sequence and the value `0`. Otherwise, get the next sequence by getting the delta of all the values. Recursively call the function and return the next sequence (to be used by the recursion) and the computed value.
4. To compute the value, we add the extrapolated value to the last number in the sequence. 

### Part 2

Do the same as Part 1, but modify how to compute the values. Take the first value in the sequence and subtract it by the extrapolated value.
