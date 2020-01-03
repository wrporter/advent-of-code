## Day 8

### Part 1

This one just asks us to parse string literals and figure the difference of code characters versus what is stored in memory.

I just looked up `golang parse string literals` and found [`strconv.Unquote`](https://golang.org/pkg/strconv/#Unquote) at the top of the results. Worked out well! I figure most programming languages will have something similar or maybe I just got lucky using Go.

### Part 2

Ah man... I finished this one so fast I would have been in the top 50 on the leaderboard! Again, not sure if I just got lucky with Golang having the tools. This was just the opposite and using [`strconv.Quote`](https://golang.org/pkg/strconv/#Quote).
