# Day 24

### Part 1

I felt like this took me a longer than it should have. In the end, it was a matter of constructing a tree and evaluating it in post-order.

### Part 2

I spent about 15 minutes thinking about how to possibly solve this one by swapping all combinations of 4 pairs of output gates. A solution like that would not only be complex, but might also not be performant enough. So I scrapped this and figured it would be easiest to compute the `x`, `y`, and incorrect `z` values and inspect which bits are wrong and traverse the trees by hand to find the wrong gates.

Later, looking at the subreddit, the top comment has a clever solution by recognizing a ripple carry adder.
