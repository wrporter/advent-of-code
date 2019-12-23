## Day 22

### Part 1

This was pretty straightforward, except that I didn't follow the last instruction of the `position` of the card `2019`. I thought it was asking for the card `at position 2019`.

It was all a matter of writing code per the exact instructions given to shuffle with each technique.

### Part 2

There was just no way for this one. I knew I didn't know enough math theory and had to resort to the Reddit threads. I based my solution off of [mcpower's explanation](https://www.reddit.com/r/adventofcode/comments/ee0rqi/2019_day_22_solutions/fbnkaju/). This was a huge pain to do in Golang due to the way they implemented their [`math/big`](https://golang.org/pkg/math/big/) package. It was incredibly annoying having to pass in the same value I was working off of into the functions.
