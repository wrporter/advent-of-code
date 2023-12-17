## Day 17

### Part 1 

Ah, the classic Subset Sum problem. I used to ask this question in software engineer interviews a lot at the beginning of my career, but soon realized it didn't give me good signal. Permutations and combinations can sometimes be challenging to implement in 30 minutes, especially since we don't do it everyday.

Of course the puzzle also leads us astray by specifically using the word "combinations" when it really means "permutations". Ugh!

But yeah, just gotta go through all the permutations from 1 item to all of them that meet the target value.

### Part 2

I kept track of the counts of all the sizes and the smallest amount of containers to optimize at the end without having to loop over the whole thing again.
