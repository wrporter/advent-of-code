# Day 19

### Part 1

My first reaction to this was, "I can use a Pushdown Automata!" I reused some code from 2020 day 19 and refactored to trim whole prefixes rather than single letters. This worked for Part 1, but was overkill because we don't need to support complex grammar configurations due to how towels can be arranged in any configuration.

### Part 2

I could no longer use my PDA solution 😭

So I rewrote everything to use a simple recursive function to track whether we have a matching prefix and push the remaining word onto the stack. Part 2 required a slight adjustment to add a cache while continuing to find all matches.
