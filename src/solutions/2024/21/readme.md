# Day 21

### Part 1

Oh man, so many hazardous areas! My initial reaction: "This looks like a pain!" Let's break down the algorithm.

1. Start from the robot at the first directional keypad. We need to know what the shortest sequence of keys is for the numeric keypad, then we can work backwards to know the third robot's movements,, then our own on the directional keypads.
2. Cache the current code and area we are in so we can quickly evaluate duplicates we've seen. This is particularly important because most of the interactions we are doing are on the directional keypads and we can speed up most of the work here.
3. Get the shortest sequence of key presses. 
   1. Get all possible sequences with a Depth-First Search by moving toward the next key.
   2. Get the sequence with the least amount of changes in direction. This is important because going `^^<<` is less expensive for downstream robots than `^<^<`. 

### Part 2

Luckily, if we followed the algorithm above for Part 1, there is no need to modify the code for Part 2 other than introduce a new parameter for how many areas there are!
