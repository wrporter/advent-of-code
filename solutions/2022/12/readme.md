# Day 12

## Thoughts

Oh man, I was very tired when I first tackled this one. I started off with a depth-first search when this is very obviously a breadth-first search problem. We just want to find out how many steps it takes for the shortest path. Depth-first search is bad here because it will spend far more time exploring deeply nested paths and backtracking before it might ever arrive at the solution.

My mind then thought I had to use a more advanced algorithm such as Dijkstra or A*. Ohhhh, was I wrong. Just breadth-first search!

For all I kept trying to implement A*, I couldn't get it to work right. I think my Priority Queue implementation was broken. Then it finally dawned on me two days later... Breadth-first search. Easy-peasy. Same solution I already wrote up, just take from the front of the queue instead of the back! 🤦

### Part 1

1. Loop over the input.
   1. Make a grid of the characters.
   2. Mark where the start and end are. Write their elevations as `a` and `z` respectively.
2. Breadth-first search to find the shortest number of steps.
   1. Start with a queue populated with just the starting position. Keep track of the number of steps already taken with each node.
   2. Keep track of which positions we've already visited.
   3. Try the first node in the queue until we've reached the end from the start.
      1. If it's been visited, skip it. Otherwise, add it to our visited set.
      2. Add all valid neighbors to the end of the queue. In bounds and only up to 1 elevation higher than the current position.

### Part 2

The only difference for Part 2 is to add all the `a` elevations as starting points. Then start the breadth-first search with all of them in the queue. And... done!

## Time

This is how long it took me to complete each part.

- Part 1: 4:27:52
- Part 2: 10:16

## Puzzle

# Day 14: Regolith Reservoir

[https://adventofcode.com/2022/day/14](https://adventofcode.com/2022/day/14)

## Description

### Part One

The distress signal leads you to a giant waterfall! Actually, hang on - the signal seems like it's coming from the waterfall itself, and that doesn't make any sense. However, you do notice a little path that leads _behind_ the waterfall.

Correction: the distress signal leads you behind a giant waterfall! There seems to be a large cave system here, and the signal definitely leads further inside.

As you begin to make your way deeper underground, you feel the ground rumble for a moment. Sand begins pouring into the cave! If you don't quickly figure out where the sand is going, you could quickly become trapped!

Fortunately, your [familiarity](https://adventofcode.com/2018/day/17) with analyzing the path of falling material will come in handy here. You scan a two-dimensional vertical slice of the cave above you (your puzzle input) and discover that it is mostly _air_ with structures made of _rock_.

Your scan traces the path of each solid rock structure and reports the `x,y` coordinates that form the shape of the path, where `x` represents distance to the right and `y` represents distance down. Each path appears as a single line of text in your scan. After the first point of each path, each point indicates the end of a straight horizontal or vertical line to be drawn from the previous point. For example:

    498,4 -> 498,6 -> 496,6
    503,4 -> 502,4 -> 502,9 -> 494,9


This scan means that there are two paths of rock; the first path consists of two straight lines, and the second path consists of three straight lines. (Specifically, the first path consists of a line of rock from `498,4` through `498,6` and another line of rock from `498,6` through `496,6`.)

The sand is pouring into the cave from point `500,0`.

Drawing rock as `#`, air as `.`, and the source of the sand as `+`, this becomes:


      4     5  5
      9     0  0
      4     0  3
    0 ......+...
    1 ..........
    2 ..........
    3 ..........
    4 ....#...##
    5 ....#...#.
    6 ..###...#.
    7 ........#.
    8 ........#.
    9 #########.


Sand is produced _one unit at a time_, and the next unit of sand is not produced until the previous unit of sand _comes to rest_. A unit of sand is large enough to fill one tile of air in your scan.

A unit of sand always falls _down one step_ if possible. If the tile immediately below is blocked (by rock or sand), the unit of sand attempts to instead move diagonally _one step down and to the left_. If that tile is blocked, the unit of sand attempts to instead move diagonally _one step down and to the right_. Sand keeps moving as long as it is able to do so, at each step trying to move down, then down-left, then down-right. If all three possible destinations are blocked, the unit of sand _comes to rest_ and no longer moves, at which point the next unit of sand is created back at the source.

So, drawing sand that has come to rest as `o`, the first unit of sand simply falls straight down and then stops:

    ......+...
    ..........
    ..........
    ..........
    ....#...##
    ....#...#.
    ..###...#.
    ........#.
    ......o.#.
    #########.


The second unit of sand then falls straight down, lands on the first one, and then comes to rest to its left:

    ......+...
    ..........
    ..........
    ..........
    ....#...##
    ....#...#.
    ..###...#.
    ........#.
    .....oo.#.
    #########.


After a total of five units of sand have come to rest, they form this pattern:

    ......+...
    ..........
    ..........
    ..........
    ....#...##
    ....#...#.
    ..###...#.
    ......o.#.
    ....oooo#.
    #########.


After a total of 22 units of sand:

    ......+...
    ..........
    ......o...
    .....ooo..
    ....#ooo##
    ....#ooo#.
    ..###ooo#.
    ....oooo#.
    ...ooooo#.
    #########.


Finally, only two more units of sand can possibly come to rest:

    ......+...
    ..........
    ......o...
    .....ooo..
    ....#ooo##
    ...o#ooo#.
    ..###ooo#.
    ....oooo#.
    .o.ooooo#.
    #########.


Once all _`24`_ units of sand shown above have come to rest, all further sand flows out the bottom, falling into the endless void. Just for fun, the path any new sand takes before falling forever is shown here with `~`:

    .......+...
    .......~...
    ......~o...
    .....~ooo..
    ....~#ooo##
    ...~o#ooo#.
    ..~###ooo#.
    ..~..oooo#.
    .~o.ooooo#.
    ~#########.
    ~..........
    ~..........
    ~..........


Using your scan, simulate the falling sand. _How many units of sand come to rest before sand starts flowing into the abyss below?_

### Part Two

You realize you misread the scan. There isn't an <span title="Endless Void is my C cover band.">endless void</span> at the bottom of the scan - there's floor, and you're standing on it!

You don't have time to scan the floor, so assume the floor is an infinite horizontal line with a `y` coordinate equal to _two plus the highest `y` coordinate_ of any point in your scan.

In the example above, the highest `y` coordinate of any point is `9`, and so the floor is at `y=11`. (This is as if your scan contained one extra rock path like `-infinity,11 -> infinity,11`.) With the added floor, the example above now looks like this:

            ...........+........
            ....................
            ....................
            ....................
            .........#...##.....
            .........#...#......
            .......###...#......
            .............#......
            .............#......
            .....#########......
            ....................
    <-- etc #################### etc -->


To find somewhere safe to stand, you'll need to simulate falling sand until a unit of sand comes to rest at `500,0`, blocking the source entirely and stopping the flow of sand into the cave. In the example above, the situation finally looks like this after _`93`_ units of sand come to rest:

    ............o............
    ...........ooo...........
    ..........ooooo..........
    .........ooooooo.........
    ........oo#ooo##o........
    .......ooo#ooo#ooo.......
    ......oo###ooo#oooo......
    .....oooo.oooo#ooooo.....
    ....oooooooooo#oooooo....
    ...ooo#########ooooooo...
    ..ooooo.......ooooooooo..
    #########################


Using your scan, simulate the falling sand until the source of the sand becomes blocked. _How many units of sand come to rest?_

