# Day 4: Giant Squid

[https://adventofcode.com/2021/day/4](https://adventofcode.com/2021/day/4)

## Thoughts & Solution

The giant squid wrapped me up in its dark tentacles and squeezed!

I ran the mistake of trying to be faster and wrote everything in the same function with a bunch of nested for loops. My solution worked for the sample input, but not the given input. I spent way too much time debugging and finally coming up with a correct solution, but with a mess that was hard to go through.

So then I made up my mind to put together a solution that was readable. I split out each board into its own data type and each of the operations into their own functions, determining a winning board, etc. This helped a lot, and I was done within only a few minutes. Perhaps the most important part is to make sure you remove boards that have already won so they are not re-evaluated for wins on consecutive number picks.

This was a fun puzzle that was still pretty easy to reason through. Definitely one that was easier, from my perspective, to write in an object-oriented way.

### Time

- Part 1: 35:30
- Part 2: 17:25

## Description

### Part One

You're already almost 1.5km (almost a mile) below the surface of the ocean, already so deep that you can't see any sunlight. What you _can_ see, however, is a giant squid that has attached itself to the outside of your submarine.

Maybe it wants to play [bingo](https://en.wikipedia.org/wiki/Bingo_(American_version))?

Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen at random, and the chosen number is _marked_ on all boards on which it appears. (Numbers may not appear on all boards.) If all numbers in any row or any column of a board are marked, that board _wins_. (Diagonals don't count.)

The submarine has a _bingo subsystem_ to help passengers (currently, you and the giant squid) pass the time. It automatically generates a random order in which to draw numbers and a random set of boards (your puzzle input). For example:

    7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1
    
    22 13 17 11  0
     8  2 23  4 24
    21  9 14 16  7
     6 10  3 18  5
     1 12 20 15 19
    
     3 15  0  2 22
     9 18 13 17  5
    19  8  7 25 23
    20 11 10 24  4
    14 21 16 12  6
    
    14 21 17 24  4
    10 16 15  9 19
    18  8 23 26 20
    22 11 13  6  5
     2  0 12  3  7
    

After the first five numbers are drawn (`7`, `4`, `9`, `5`, and `11`), there are no winners, but the boards are marked as follows (shown here adjacent to each other to save space):

    22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
     8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
    21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
     6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
     1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
    

After the next six numbers are drawn (`17`, `23`, `2`, `0`, `14`, and `21`), there are still no winners:

    22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
     8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
    21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
     6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
     1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
    

Finally, `24` is drawn:

    22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
     8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
    21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
     6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
     1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
    

At this point, the third board _wins_ because it has at least one complete row or column of marked numbers (in this case, the entire top row is marked: _`14 21 17 24 4`_).

The _score_ of the winning board can now be calculated. Start by finding the _sum of all unmarked numbers_ on that board; in this case, the sum is `188`. Then, multiply that sum by _the number that was just called_ when the board won, `24`, to get the final score, `188 * 24 = _4512_`.

To guarantee victory against the giant squid, figure out which board will win first. _What will your final score be if you choose that board?_

### Part Two

On the other hand, it might be wise to try a different strategy: <span title="That's 'cuz a submarine don't pull things' antennas out of their sockets when they lose. Giant squid are known to do that.">let the giant squid win</span>.

You aren't sure how many bingo boards a giant squid could play at once, so rather than waste time counting its arms, the safe thing to do is to _figure out which board will win last_ and choose that one. That way, no matter which boards it picks, it will win for sure.

In the above example, the second board is the last to win, which happens after `13` is eventually called and its middle column is completely marked. If you were to keep playing until this point, the second board would have a sum of unmarked numbers equal to `148` for a final score of `148 * 13 = _1924_`.

Figure out which board will win last. _Once it wins, what would its final score be?_
