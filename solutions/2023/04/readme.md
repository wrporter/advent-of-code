# Day 4

## Thoughts

The stories about the elves simply make me chuckle, making me think of my own children. 

I originally didn't have the parsing code separated for speed of coding and performance. But I split it out after I solved the puzzle because it made for far cleaner code and was very helpful for Part 2 since we needed to perform some additional processing.

### Part 1

For a minute or two, I got tripped up with my getting my powers of 2 backwards. Overall, pretty straightforward. Most of my time was spent in the initial parsing of the input.

1. Iterate over each card.
2. Iterate over each number that card has.
3. Check to see if each number is in the list of winning numbers (use a set for performance here).
4. Count the amount of matching numbers.
5. Sum 2 the power of the matching numbers minus 1. This is to account for no matches. If there aren't any matches we get `2^-1 = 0` (truncated as 0 for being an integer).

### Part 2

I messed up by initially converting my cards to a map and realized through debugging that iterating over a map does not guarantee order. So I flipped that to an array and it worked like a charm. The ultimate solution for this part is much like Part 1.

1. Now we add the number of copies to each card and keep track of them.
2. Iterate over each card.
3. Iterate over each number that card has.
4. Check to see if each number is in the list of winning numbers (use a set for performance here).
5. Count the amount of matching numbers.
6. For each matching number, add to the amount of copies each consecutive card has, as long as we don't hit the end of the list of cards.
7. Sum up the total number of copies of the current card.
