# Day 1

## Thoughts

Yet another exciting year 😊. 

### Part 1

Some optimizations I made at the start:

1. Provide the capacity for the lists so they do not have resize.
2. Insert each number into the list in sorted order, so we don't have to sort after parsing the lines.

### Part 2

To simplify, I used a map of counts for each number's appearance in the second list.
