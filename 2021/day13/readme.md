# Day 13: Transparent Origami

[https://adventofcode.com/2021/day/13](https://adventofcode.com/2021/day/13)

## Thoughts & Solution

As usual, it took me about 10 minutes to read the puzzle description, especially as I was still waking up 😄.

Part 1: It took me a while to figure out the right math for folding points, I think because my brain was still half asleep. I wrote down the differences for each point that gets folded and realized that there's a pattern of a multiple of 2 and the difference between how far you are from the fold. I just maintain a set of points, delete any that are on a fold line, and modify the points that are getting folded.

Part 2: I reused the folding code from Part 1. This one just made us have to iterate all the folds now and then put together a 2D grid of the remaining points. I was able to copy/paste code from a previous year's puzzle and modify it to do this as it's been super useful to just maintain a set of points rather than resizing 2D arrays. My output came out upside down though. I'm not sure if that was intended or if my logic is backwards somewhere. Maybe I'll come back later to figure it out.

### Time

- Part 1: 20:19
- Part 2: 5:50

## Description

### Part One

You reach another volcanically active part of the cave. It would be nice if you could do some kind of thermal imaging so you could tell ahead of time which caves are too hot to safely enter.

Fortunately, the submarine seems to be equipped with a thermal camera! When you activate it, you are greeted with:

    Congratulations on your purchase! To activate this infrared thermal imaging
    camera system, please enter the code found on page 1 of the manual.
    

Apparently, the Elves have never used this feature. To your surprise, you manage to find the manual; as you go to open it, page 1 falls out. It's a large sheet of [transparent paper](https://en.wikipedia.org/wiki/Transparency_(projection))! The transparent paper is marked with random dots and includes instructions on how to fold it up (your puzzle input). For example:

    6,10
    0,14
    9,10
    0,3
    10,4
    4,11
    6,0
    6,12
    4,1
    0,13
    10,12
    3,4
    3,0
    8,4
    1,10
    2,14
    8,10
    9,0
    
    fold along y=7
    fold along x=5
    

The first section is a list of dots on the transparent paper. `0,0` represents the top-left coordinate. The first value, `x`, increases to the right. The second value, `y`, increases downward. So, the coordinate `3,0` is to the right of `0,0`, and the coordinate `0,7` is below `0,0`. The coordinates in this example form the following pattern, where `#` is a dot on the paper and `.` is an empty, unmarked position:

    ...#..#..#.
    ....#......
    ...........
    #..........
    ...#....#.#
    ...........
    ...........
    ...........
    ...........
    ...........
    .#....#.##.
    ....#......
    ......#...#
    #..........
    #.#........
    

Then, there is a list of _fold instructions_. Each instruction indicates a line on the transparent paper and wants you to fold the paper _up_ (for horizontal `y=...` lines) or _left_ (for vertical `x=...` lines). In this example, the first fold instruction is `fold along y=7`, which designates the line formed by all of the positions where `y` is `7` (marked here with `-`):

    ...#..#..#.
    ....#......
    ...........
    #..........
    ...#....#.#
    ...........
    ...........
    -----------
    ...........
    ...........
    .#....#.##.
    ....#......
    ......#...#
    #..........
    #.#........
    

Because this is a horizontal line, fold the bottom half _up_. Some of the dots might end up overlapping after the fold is complete, but dots will never appear exactly on a fold line. The result of doing this fold looks like this:

    #.##..#..#.
    #...#......
    ......#...#
    #...#......
    .#.#..#.###
    ...........
    ...........
    

Now, only `17` dots are visible.

Notice, for example, the two dots in the bottom left corner before the transparent paper is folded; after the fold is complete, those dots appear in the top left corner (at `0,0` and `0,1`). Because the paper is transparent, the dot just below them in the result (at `0,3`) remains visible, as it can be seen through the transparent paper.

Also notice that some dots can end up _overlapping_; in this case, the dots merge together and become a single dot.

The second fold instruction is `fold along x=5`, which indicates this line:

    #.##.|#..#.
    #...#|.....
    .....|#...#
    #...#|.....
    .#.#.|#.###
    .....|.....
    .....|.....
    

Because this is a vertical line, fold _left_:

    #####
    #...#
    #...#
    #...#
    #####
    .....
    .....
    

The instructions made a square!

The transparent paper is pretty big, so for now, focus on just completing the first fold. After the first fold in the example above, _`17`_ dots are visible - dots that end up overlapping after the fold is completed count as a single dot.

_How many dots are visible after completing just the first fold instruction on your transparent paper?_

### Part Two

<span title="How can you fold it that many times? You tell me, I'm not the one folding it.">Finish folding</span> the transparent paper according to the instructions. The manual says the code is always _eight capital letters_.

_What code do you use to activate the infrared thermal imaging camera system?_
