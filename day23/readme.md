## Day 23

### Part 1

Some more fun with the Intcode computer! I've learned a lot about concurrency in Golang with these puzzles. For this, we can use a `select` statement to unblock the channel operations. If there is no output from the current computer, then take in an input of `-1`.
