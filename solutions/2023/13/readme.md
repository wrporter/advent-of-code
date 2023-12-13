# Day 13

## Thoughts

I've seen similar problems in Advent of Code in the past, so I thought of bitmasks to speed things up. Not sure if it's faster or not though.

### Part 1

I thought it'd be fun and a slight optimization to use a generator for the parsing. That way, we don't have to loop through the grids twice.

1. Parse the input.
   1. Create separate arrays for the rows and columns. If we did this with strings rather than bitmasks, we'd use the initial grid for row checking, then transpose it for column checking.
   2. Construct the rows and columns. Each time we see a `#`, count it as a bit that we set. Using bitwise operators, we can `or` the current value with the respective row or column. If we want to mirror the grid, we can flip sides by `or`-ing with `width - col - 1`. I made a very slight optimization by flipping the grid. This works because we are only evaluating the opposite axis.
   3. Use a function parameter to process the values as a generator. This concept differs slightly per language.
2. Summarize the patterns by taking the sum of `100 * row_mirror_index + col_mirror_index`.
3. Loop through each value in the row or column, excluding the last one, in case we have an odd number of values and cannot compare an end value.
   1. Start by assuming that the pattern is reflecting.
   2. Set the previous index to the current position.
   3. Set the next index to the next position.
   4. Loop until our previous and next pointers reach the end of the pattern and we still see a reflection.
      1. Update whether we are still reflecting if the previous and next values match.
      2. Update the previous and next values.
   5. If we see a full reflection, then return the current `index + 1`, since the row/column values are off-by-one in the puzzle description.
   6. If no reflection is seen, return `0`.

### Part 2

This part was very similar to Part 1, but it took me a while to figure out how to determine a smudged position with bitwise logic. I ended up Googling "binary check if only 1 bit is set" and found this great [StackOverflow post](https://stackoverflow.com/a/51094793) that explains that we can simply check if the value is a power of 2. First I `XOR` the two values then check if the result is a power of 2, meaning only 1 bit differs.

The process is then nearly the same, but we keep track of whether a value in the pattern is "smudged" by being a single bit off. Only one value can be smudged while checking the reflection, as stated in the puzzle description.
