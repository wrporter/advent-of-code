# Day 6

## Thoughts

Classic cycle detection in AoC.

### Part 1

The problem looked large at first, but was pretty straightforward in implementation.

### Part 2

We can optimize by only placing obstacles along the guard's original path. Other locations won't make a difference because the guard will never travel there.

I had a bug in my cycle detection logic. I would take a step forward at the same time as rotating. We need to add to our seen locations just after rotating as well to properly detect every moment of a cycle.
