## Day 20

### Part 1

I did not know an efficient algorithm for getting divisors of a number without looping through all numbers from 1 to our target value. So I looked that one up, but was on an okay track, I felt like. I just loop through all houses from 1 on up and calculate the number of presents delivered to each house. Once we reach a house that has at least as much as the puzzle input, we return that house number.

### Part 2

It's too bad that there's no smaller, sane example to test with. I bet there's some kind of math trick to calculate this straight out, but I don't remember all my math theory from college over 10 years ago. I'm writing loops and getting around 3.6 seconds for each answer. Definitely not performant, but good enough for me on these puzzles without the special knowledge.
