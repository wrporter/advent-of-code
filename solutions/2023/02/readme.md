# Day 2

## Thoughts

I'm pretty sure most of my time was spent parsing the input 😄. The logic for both parts was pretty straightforward and what I'd expect in these early days. I initially solved the parts with the parsing inline. This is slightly more optimal, so we don't have to run through each game and cube set twice, but it makes for code that is harder to read. So I spilt it up after solving the day and my times are still in the micro-second range, so that's totally acceptable.

I didn't do anything special with parsing, just a lot of string splitting. I used to use a lot of fancy regexes for this sort of stuff, but the repeating parts would have been more challenging to get right.

### Part 1

1. Parse the input.
2. For each game, keep track of whether it is possible, assume that it is at the start.
3. For each set and each cube, if it exceeds our max amount in the bag, consider the game impossible.
4. After going through each set in a game, if the game is possible, sum its ID.
5. Return the sum.

### Part 2

1. Parse the input.
2. For each game, keep track of the max amount of cubes used for each color.
3. Once we've determined the max amounts, calculate the power and add them to the overall sum.
4. Return the sum.
