## Day 23

### Part 1

Some more fun with the Intcode computer! I've learned a lot about concurrency in Golang with these puzzles. For this, we can use a `select` statement to unblock the channel operations. If there is no output from the current computer, then take in an input of `-1`.

### Part 2

Again, this one was fairly simple, just store the last NAT and previous NAT packets. When we get to 50 idle computers, send the last NAT packet to address `0`. 

However, there was a little trick I struggled to figure out. We should assume that the network starts out idle and therefore start by passing each computer `-1` to indicate that it's packet queue starts out empty.
