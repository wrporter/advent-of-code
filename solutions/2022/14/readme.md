# Day 14

## Thoughts

This was a refreshing and fun puzzle today. I love how Part 2 makes the sand look like a Christmas tree 😂. Most of my time was spent making sure I created the right lines of rocks.

### Part 1

1. Parse the scan results.
   1. Start with a set of points.
   2. For each path, create lines of rocks and add them to the set.
   3. Keep track of the highest `y` value so you know where the bottom is.
2. Use a function to determine when a unit of sand has entered the void. This is when its `y` value has reached beyond the `bottom` we calculated earlier.
3. Use a function to determine whether the sand can fall to a give position.
4. Count all the sand that comes to rest.
   1. Loop until we've reached our exit criteria - sand entering the void.
      1. Create a unit of sand at the source.
      2. Loop until it comes to rest or enters the void.
         1. Sand comes to rest if it can no longer fall below, to the left, or the right.
         2. Otherwise, move the sand down, down-left, or down-right in respective order.

### Part 2

Part 2 is just like Part 1. Here are the differences.

1. Add 2 to the bottom.
2. Sand can come to rest at the bottom too now, not just around other sand and rocks.
3. We stop when sand reaches the source rather than going out into the void, which is now impossible.

## Time

This is how long it took me to complete each part.

- Part 1: 41:37
- Part 2: 

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

