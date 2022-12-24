# Day 22

## Thoughts

Wow, this was an infuriating yet incredibly fun problem, especially when I had to cut out paper to visualize and calculate the wrapping onto various sides.

![Cube](cube.jpg)

I am exhausted with these long days and don't feel motivated right now to clean up the code to share commonalities between Part 1 and Part 2. Nor do I want to clean up the logic. This will have to be it for now...

### Part 1

1. Parse the current state of the map.
   1. Detect the current location on the map in the top left open spot.
   2. Add all characters to a 2D array of strings.
   3. Parse the instructions into a list of rotations and moves.
2. Execute each instruction.
   1. If it's a rotation instruction, rotate direction left or right.
   2. Move the appropriate amount.
      1. Stop when hitting a wall.
      2. If we go out of bounds or hit an empty space ` `, then we know we need to wrap. My current solution to this is to clone my current position, then walk the clone the other direction until it falls off the end, then re-add my position. If it's at a wall, then don't move, otherwise step there and change direction 180 degrees.
      3. If the next spot is an open space `.` then move to it.
3. Calculate the password `1000 * (y + 1) + 4 * (x + 1) + facing`.

### Part 2

Everything is the same except the wrapping logic. We now need to fold our input into a cube. I have not yet thought of a generic way to do this for different configurations (such as the sample versus real input). However, I did solve for a generic way to wrap to other faces by tracking the relative position within a face and determining the right direction when wrapping onto the next face. I just didn't finish this part for all possible face configurations.

## Time

This is how long it took me to complete each part.

- Part 1: 1:17:25
- Part 2: 2:36:14

## Puzzle

# Day 22: Monkey Map

[https://adventofcode.com/2022/day/22](https://adventofcode.com/2022/day/22)

## Description

### Part One

The monkeys take you on a surprisingly easy trail through the jungle. They're even going in roughly the right direction according to your handheld device's Grove Positioning System.

As you walk, the monkeys explain that the grove is protected by a _force field_. To pass through the force field, you have to enter a password; doing so involves tracing a specific _path_ on a strangely-shaped board.

At least, you're pretty sure that's what you have to do; the elephants aren't exactly fluent in monkey.

The monkeys give you notes that they took when they last saw the password entered (your puzzle input).

For example:

            ...#
            .#..
            #...
            ....
    ...#.......#
    ........#...
    ..#....#....
    ..........#.
            ...#....
            .....#..
            .#......
            ......#.
    
    10R5L5R10L4R5L5


The first half of the monkeys' notes is a _map of the board_. It is comprised of a set of _open tiles_ (on which you can move, drawn `.`) and _solid walls_ (tiles which you cannot enter, drawn `#`).

The second half is a description of _the path you must follow_. It consists of alternating numbers and letters:

*   A _number_ indicates the _number of tiles to move_ in the direction you are facing. If you run into a wall, you stop moving forward and continue with the next instruction.
*   A _letter_ indicates whether to turn 90 degrees _clockwise_ (`R`) or _<span title="Or &quot;anticlockwise&quot;, if you're anti-counterclockwise.">counterclockwise</span>_ (`L`). Turning happens in-place; it does not change your current tile.

So, a path like `10R5` means "go forward 10 tiles, then turn clockwise 90 degrees, then go forward 5 tiles".

You begin the path in the leftmost open tile of the top row of tiles. Initially, you are facing _to the right_ (from the perspective of how the map is drawn).

If a movement instruction would take you off of the map, you _wrap around_ to the other side of the board. In other words, if your next tile is off of the board, you should instead look in the direction opposite of your current facing as far as you can until you find the opposite edge of the board, then reappear there.

For example, if you are at `A` and facing to the right, the tile in front of you is marked `B`; if you are at `C` and facing down, the tile in front of you is marked `D`:

            ...#
            .#..
            #...
            ....
    ...#.D.....#
    ........#...
    B.#....#...A
    .....C....#.
            ...#....
            .....#..
            .#......
            ......#.


It is possible for the next tile (after wrapping around) to be a _wall_; this still counts as there being a wall in front of you, and so movement stops before you actually wrap to the other side of the board.

By drawing the _last facing you had_ with an arrow on each tile you visit, the full path taken by the above example looks like this:

            >>v#    
            .#v.    
            #.v.    
            ..v.    
    ...#...v..v#    
    >>>v...>#.>>    
    ..#v...#....    
    ...>>>>v..#.    
            ...#....
            .....#..
            .#......
            ......#.


To finish providing the password to this strange input device, you need to determine numbers for your final _row_, _column_, and _facing_ as your final position appears from the perspective of the original map. Rows start from `1` at the top and count downward; columns start from `1` at the left and count rightward. (In the above example, row 1, column 1 refers to the empty space with no tile on it in the top-left corner.) Facing is `0` for right (`>`), `1` for down (`v`), `2` for left (`<`), and `3` for up (`^`). The _final password_ is the sum of 1000 times the row, 4 times the column, and the facing.

In the above example, the final row is `6`, the final column is `8`, and the final facing is `0`. So, the final password is 1000 \* 6 + 4 \* 8 + 0: _`6032`_.

Follow the path given in the monkeys' notes. _What is the final password?_

### Part Two

As you reach the force field, you think you hear some Elves in the distance. Perhaps they've already arrived?

You approach the strange _input device_, but it isn't quite what the monkeys drew in their notes. Instead, you are met with a large _cube_; each of its six faces is a square of 50x50 tiles.

To be fair, the monkeys' map _does_ have six 50x50 regions on it. If you were to _carefully fold the map_, you should be able to shape it into a cube!

In the example above, the six (smaller, 4x4) faces of the cube are:

            1111
            1111
            1111
            1111
    222233334444
    222233334444
    222233334444
    222233334444
            55556666
            55556666
            55556666
            55556666


You still start in the same position and with the same facing as before, but the _wrapping_ rules are different. Now, if you would walk off the board, you instead _proceed around the cube_. From the perspective of the map, this can look a little strange. In the above example, if you are at A and move to the right, you would arrive at B facing down; if you are at C and move down, you would arrive at D facing up:

            ...#
            .#..
            #...
            ....
    ...#.......#
    ........#..A
    ..#....#....
    .D........#.
            ...#..B.
            .....#..
            .#......
            ..C...#.


Walls still block your path, even if they are on a different face of the cube. If you are at E facing up, your movement is blocked by the wall marked by the arrow:

            ...#
            .#..
         -->#...
            ....
    ...#..E....#
    ........#...
    ..#....#....
    ..........#.
            ...#....
            .....#..
            .#......
            ......#.


Using the same method of drawing the _last facing you had_ with an arrow on each tile you visit, the full path taken by the above example now looks like this:

            >>v#    
            .#v.    
            #.v.    
            ..v.    
    ...#..^...v#    
    .>>>>>^.#.>>    
    .^#....#....    
    .^........#.    
            ...#..v.
            .....#v.
            .#v<<<<.
            ..v...#.


The final password is still calculated from your final position and facing from the perspective of the map. In this example, the final row is `5`, the final column is `7`, and the final facing is `3`, so the final password is 1000 \* 5 + 4 \* 7 + 3 = _`5031`_.

Fold the map into a cube, _then_ follow the path given in the monkeys' notes. _What is the final password?_

