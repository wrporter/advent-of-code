## Day 15

### Part 1

This was the first puzzle where the algorithm was not immediately apparent to me for 2015. After some thinking, I figured it would be easier if I could compute all possibilities that add up to our teaspoon amount (`100`) given the number of slots (`4` ingredients). This simplified my thinking quite a bit as that was the most challenging part of the problem. The rest was evaluating against every combination of the ingredient amounts to see which one gave the highest total score.

## Part 2

Again, just a slight change in the program. I added a flag for the calorie amount so it could still be used for Part 1. All we had to do was add an extra if statement to check if we are at the desired calorie amount.
