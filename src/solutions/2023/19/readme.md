# Day 19

## Thoughts

It was fun making a workflow state machine for this one.

### Part 1

1. Parse each of the workflows and part ratings.
2. For each rating, run through the workflows to determine if each is accepted or rejected.
3. Add the total value of the ratings when accepted.

### Part 2

1. Parse the values as before.
2. With a depth-first search, recurse until a range of values is accepted. If the workflow is rejected, return 0. If it is accepted, return the number of possibilities by multiplying the accepted rating ranges.
3. For each rule in the current workflow, construct the next range and recurse.
