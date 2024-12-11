# Day 10

### Part 1

I ran this with a BFS to find the 9s. The only nuance is we don't want duplicate 9s to count, so I store the locations in a set.

**EDIT [6fe92f0](https://github.com/wrporter/advent-of-code/commit/6fe92f05992eee0a236cb690135bfd5ad930ca16):** I spent 2 minutes combining the solutions for both parts because DFS is faster than BFS (450µs vs 800µs). I **think** the only reasons for the slowness in the BFS solution is due to (1) changing size of the queue and (2) O(log(n)) map comparisons to ensure we don't duplicate scores.

### Part 2

This one requires DFS so we can backtrack. I realized as soon as I finished that I could have used DFS for both parts and just excluded backtracking for Part 1.
