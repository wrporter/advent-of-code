# Day 1

## Thoughts

I have been awaiting Advent of Code quite anxiously. So excited to get up early this morning to enjoy the first puzzles!

### Part 1

This one was a cakewalk, as day 1 normally is. Fun and easy to think through.

1. Loop through each line.
2. Use a for loop to find the first digit.
3. Reverse a for loop to find the last digit.
4. Concatenate the digits as strings and convert to a number, then add to the total sum.

### Part 2

Now this one was a doozy 😅! Maybe there are better approaches, but I went for the dumb approach and created a map of all the numbers to their associated values. I did a quick manual check of the input to make sure there wasn't anything over 20 (as words).

1. Loop through each line.
2. Add all the found numbers to a map of locations. Make sure to get the first and last index in case there are duplicates.
3. Find the first and last numbers.
4. Make sure to get the first and last digit within those numbers in case there are teens to account for.
5. Concatenate the digits as strings and convert to a number, then add to the total sum.
