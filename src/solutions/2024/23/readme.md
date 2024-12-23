# Day 23

### Part 1

This puzzle might look scarier than it really is. 

1. Construct a map of all the connections between computers.
2. Use a triple-for loop to get all combinations of connections.
   1. Only continue if at least one of the computers starts with a `t`.
   2. Increment count if all 3 computers are connected.

### Part 2

This part wasn't much more difficult. I was worried we'd have to reach for some network algorithm like in years past.

1. Initialize a network for each computer.
2. Find all other computers that are connected and belong to that network.
3. Keep track of the network with the most computers.
4. Sort the largest network and join with `,` to get the password. We're in!
