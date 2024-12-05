# Day 5

## Thoughts

At first glance, I didn't feel like trying to solve this puzzle this morning. It looked long and complicated. After taking my time to read through it all, I started to formulate a decent plan and then the code just splatted into the editor. I was surprised by how much my design I could use for Part 2. Time to complete Part 2 after Part 1 was only 3m34s!

### Part 1

1. For each number, keep track of a set of numbers that come after it.
2. Loop through each update in reverse order. For each page, check against the rules to make sure all the numbers after it are supposed to be after it.
3. If any update violates a rule, add to the middle page to the sum.

### Part 2

Flip the final piece in Part 1 and sort incorrect updates based on the rules. Then proceed to add the middle page to the sum.
