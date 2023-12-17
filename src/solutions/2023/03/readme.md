# Day 3

## Thoughts

I was surprised how challenging this one was for a Day 3 puzzle. I got up at 4am to massage my 3-year-old's feet for about an hour because he was having growing pains. After that, I turned on a movie and snuggled him on the couch to get his mind off the pain. Maybe my groggy mind was part of it.

### Part 1

I think this was legitimately a little challenging due to off-by-one errors. I went through 5 attempts before my answer was finally accepted. I ended up adding a debug statement to capture every number that was not part of the schematic and the surrounding area. I also added several small test cases. Here's what I found wrong:

1. I was cutting numbers off at the end.
2. I was excluding numbers that ran up against the right edge.

Here's my process:

1. Loop through each row and column.
2. If a character is a digit, start capturing the full number.
3. Once we hit a non-digit, capture the full number.
4. Check surrounding indices to see if there is an adjacent symbol.
5. If there is an adjacent symbol, we know the number is part of the schematic, so add it to the sum.

### Part 2

I opted for, perhaps a longer, but simpler approach. I decided to take a pass through everything and gather number and star positions. Then loop through each star and see if there were exactly two adjacent numbers. A very naive approach, but I think I was faster at it than if I did something else.
