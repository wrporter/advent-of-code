# Day 22

### Part 1

Wow, you really gotta read those instructions. Once fully understood though, the solution is fairly straightforward.

### Part 2

As I was going through the description, I couldn't help but think of the song, [Go Bananas!](https://www.youtube.com/watch?v=o6gHL1LJ-HQ). My family has loved that song lately. And I felt like I was going a little bananas when trying to understand this monkey logic...

So here goes the algo:

1. Initialize a map of change sequences to total bananas.
2. For each buyer:
   1. Calculate the next secret number.
   2. Store the change in price from the last number.
   3. Once we have 4 changes.
      1. Track sequences of the last 4 changes.
      2. If we haven't seen this sequence yet, meaning it's the first time we've seen it for this buyer, add the current price  (number of bananas) for that sequence to our map.
3. Find the sequence with the best price.
