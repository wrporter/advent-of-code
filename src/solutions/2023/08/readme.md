# Day 8

## Thoughts

This was a fun graph traversal problem. I feel like we don't get to do those very often. The problem even talks about a "network" and the input is listed like "nodes". I had the sneaking suspicion that we'd have to optimize for Part 2. My initial thoughts were, "Oh, are we going to need depth-first search with pruning? I sure hope not, those are the toughest..."

### Part 1

I spent several minutes trying to remember how to parse a regex in Go 😅. We just don't do that enough in our day job, apparently.

1. Parse the input.
   1. We can just keep track of the instructions from the first line in a string.
   2. Store the network (graph) in a map of nodes with a left, right pair.
2. Set our starting location to `AAA` and steps to `0`.
3. Keep taking steps until we arrive at `ZZZ`.
   1. Get the next instruction based on the current step, modded by the length of the instructions so we keep repeating instructions.
   2. Increment the number of steps we've taken.
   3. Follow the node in our graph based on the left/right instruction.
4. We eventually make it to `ZZZ` and can return the number of steps.

### Part 2

This is the first time this thought came naturally to me, even before I started coding. Advent of Code has gotten me so used to spotting patterns, cycles, and pruning. I figured that stepping through would probably take longer than I'd want on my computer cycles. But first, I tried the naive approach.

1. Parse the input as before.
2. Set all of our starting locations to those that end in `A`.
3. Keep track of whether we've reached the end, denoted by all current locations ending in `Z`.
4. Keep taking steps until we arrive at the end.
   1. Get the next instruction, as before.
   2. Increment the number of steps, as before.
   3. For each of our current locations, take a step left or right, as before.

I'm not sure how long this would take to complete, but as soon as I ran it, my suspicions were confirmed. We are going to have to find a cycle. My thoughts immediately turned to the Chinese Remainder Theorem. So I immediately added logic to find the first step number that each location arrives at a location ending in `Z`. When I printed out the numbers; however, they were not prime.

I immediately jumped to using the least common multiple of the values. I already have functions in my mini-AoC library for that, so I plugged that in and got a result super quick. Plugged that in and it got accepted!

My solution takes about `11ms` so I feel like there might be other optimizations I'm not thinking of, but I will let it rest there.

**Edit:** The performance on this bothered me, so I thought about it a little more and realized two optimizations I could make:

1. Once we've reached a location that ends in `Z` we can stop looping over that location. By extracting the similar code from Part 1, it made this simpler to reorder the logic because I had a lot more going on in Part 2 that I didn't need, such as maps keeping track of the chain of locations. This got my solution down to `3ms` from `11ms`. Then I went to my next optimization.
2. We can process all the paths in parallel. Either way, we'd be using the same number of CPU cycles, so we might as well put computing power to work! So I added a wait group and mutex to thread-safely update the list of step counts. This gets us down to about `1.3ms`. Now I am truly satisfied 😄.
