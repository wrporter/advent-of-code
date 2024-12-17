# Day 17

### Part 1

Oh, what fun it is to ride on an opcode computer, hey!

I loved the intcode problems in 2019, my first year playing AoC. So it was fun to revisit that. Part 1 was a breeze by just following the long instructions.

### Part 2

Now this one was ugly. I was up with sick kids all night and didn't feel like trying to deconstruct the program to figure out what it was doing. So I looked up hints on the subreddit. The code loops through 3 bits at a time until it finds an increasing match on the end of the program. This cuts branches of the search space considerably.
