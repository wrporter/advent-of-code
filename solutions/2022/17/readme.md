# Day 17

## Thoughts

This was a fun variation on Tetris! I was excited, thinking we were going to simulate playing the game as part of the puzzle. At least it was kinda close 😊.

### Part 1

1. I initialize a chamber at the beginning for visualizations. The chamber is a map of coordinates with their respective characters. Other than this initialization, I only use it to store locations of rock chunks. There is no real parsing to do here because our jets are just a single string read from the file.
2. Keep track of the following.
   1. The rock pattern we are currently on. We will cycle through these.
   2. The jet we are on. We will also cycle through these.
   3. How tall the rock structure is, starts at `0`.
3. Simulate rocks falling until we get to our max rock count of `2022`.
   1. Create a rock structure and drop is until it comes to rest.
   2. Detect if a jet pushes the rock.
      1. Nothing happens if the rock hits the left or right wall.
      2. Nothing happens if the rock collides with another rock. Check all rock chunks on the current rock and see if it's going to hit any existing chunks in our map.
      3. Push the rock in the correct direction otherwise.
   3. Cycle to the next jet.
   4. Detect if the rock comes to rest.
      1. The rock comes to rest if it hits the bottom of the chamber or collides with a rock underneath it.
      2. We set the height of the chamber to the tallest point on the rock. In my case, I decided to track rocks by their top-left coordinate position.
      3. Otherwise, the rock continue to move down.
   5. Cycle to the next rock pattern.
4. Return the height of the rock structure.

### Part 2

I immediately knew we needed to detect a cycle, but I wasn't sure how at first. And it was time to go throughout my day. This and the puzzles for the surrounding days were mentally exhausting for me, so I felt like I needed a break from it for a while. Coming back days later, it was clearer to me.

I was able to use my same solution for both parts while adding cycle detection and jumping forward.

1. Before simulating rocks, add a map of states we've seen. We key this off of the following:
   1. The top few rows (I chose 7) of the rock structure. I just make this a single string of `.` (air) and `#` (rock).
   2. The current rock pattern ID.
   3. The current jet index.
2. If we encounter a state we've seen, then we've hit a cycle!
3. Calculate how many repetitions there are in the cycle to get to our desired maximum rock count. 
   1. In my example, the start of the cycle was at rock number `75` and repeated at rock number `1775`. So a cycle is `1775 - 75 = 1700` rocks after the `75th` rock.
   2. Now I need to jump forward as far as we are able just below our maximum rock count. We take `(maxRocks - rocksAlreadyPlaced) / (cycleSize)` or `(1000000000000 - 1775) / 1700 = 588235293`. This is how many times we repeat the cycle right before we get to our huge max number.
   3. Then we can jump our rock counter forward by that much `1700 * 588235293 = 999999999875`.
   4. Now we can figure how high we can jump without having to wait for more rocks to drop. This is `cycleSize * (height + cycleStartHeight)` or `588235293 * (2768 + 126) = 1554117644106`.
4. Then we go through the remaining rocks to reach `1000000000000`, so `125` more and add that to our tracked height.
5. Then we'll exit the cycle and return our height plus our jump height. For me this is `2964 + 1554117644106 = 1554117647070`.

## Time

This is how long it took me to complete each part.

- Part 1: 1:27:48
- Part 2: 2+ hours

## Puzzle

# Day 17: Pyroclastic Flow

[https://adventofcode.com/2022/day/17](https://adventofcode.com/2022/day/17)

## Description

### Part One

Your handheld device has located an alternative exit from the cave for you and the elephants. The ground is rumbling almost continuously now, but the strange valves bought you some time. It's definitely getting warmer in here, though.

The tunnels eventually open into a very tall, narrow chamber. Large, oddly-shaped rocks are falling into the chamber from above, presumably due to all the rumbling. If you can't work out where the rocks will fall next, you might be <span title="I am the man who arranges the blocks / that descend upon me from up above!">crushed</span>!

The five types of rocks have the following peculiar shapes, where `#` is rock and `.` is empty space:

    ####
    
    .#.
    ###
    .#.
    
    ..#
    ..#
    ###
    
    #
    #
    #
    #
    
    ##
    ##


The rocks fall in the order shown above: first the `-` shape, then the `+` shape, and so on. Once the end of the list is reached, the same order repeats: the `-` shape falls first, sixth, 11th, 16th, etc.

The rocks don't spin, but they do get pushed around by jets of hot gas coming out of the walls themselves. A quick scan reveals the effect the jets of hot gas will have on the rocks as they fall (your puzzle input).

For example, suppose this was the jet pattern in your cave:

    >>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>


In jet patterns, `<` means a push to the left, while `>` means a push to the right. The pattern above means that the jets will push a falling rock right, then right, then right, then left, then left, then right, and so on. If the end of the list is reached, it repeats.

The tall, vertical chamber is exactly _seven units wide_. Each rock appears so that its left edge is two units away from the left wall and its bottom edge is three units above the highest rock in the room (or the floor, if there isn't one).

After a rock appears, it alternates between _being pushed by a jet of hot gas_ one unit (in the direction indicated by the next symbol in the jet pattern) and then _falling one unit down_. If any movement would cause any part of the rock to move into the walls, floor, or a stopped rock, the movement instead does not occur. If a _downward_ movement would have caused a falling rock to move into the floor or an already-fallen rock, the falling rock stops where it is (having landed on something) and a new rock immediately begins falling.

Drawing falling rocks with `@` and stopped rocks with `#`, the jet pattern in the example above manifests as follows:

    The first rock begins falling:
    |..@@@@.|
    |.......|
    |.......|
    |.......|
    +-------+
    
    Jet of gas pushes rock right:
    |...@@@@|
    |.......|
    |.......|
    |.......|
    +-------+
    
    Rock falls 1 unit:
    |...@@@@|
    |.......|
    |.......|
    +-------+
    
    Jet of gas pushes rock right, but nothing happens:
    |...@@@@|
    |.......|
    |.......|
    +-------+
    
    Rock falls 1 unit:
    |...@@@@|
    |.......|
    +-------+
    
    Jet of gas pushes rock right, but nothing happens:
    |...@@@@|
    |.......|
    +-------+
    
    Rock falls 1 unit:
    |...@@@@|
    +-------+
    
    Jet of gas pushes rock left:
    |..@@@@.|
    +-------+
    
    Rock falls 1 unit, causing it to come to rest:
    |..####.|
    +-------+
    
    A new rock begins falling:
    |...@...|
    |..@@@..|
    |...@...|
    |.......|
    |.......|
    |.......|
    |..####.|
    +-------+
    
    Jet of gas pushes rock left:
    |..@....|
    |.@@@...|
    |..@....|
    |.......|
    |.......|
    |.......|
    |..####.|
    +-------+
    
    Rock falls 1 unit:
    |..@....|
    |.@@@...|
    |..@....|
    |.......|
    |.......|
    |..####.|
    +-------+
    
    Jet of gas pushes rock right:
    |...@...|
    |..@@@..|
    |...@...|
    |.......|
    |.......|
    |..####.|
    +-------+
    
    Rock falls 1 unit:
    |...@...|
    |..@@@..|
    |...@...|
    |.......|
    |..####.|
    +-------+
    
    Jet of gas pushes rock left:
    |..@....|
    |.@@@...|
    |..@....|
    |.......|
    |..####.|
    +-------+
    
    Rock falls 1 unit:
    |..@....|
    |.@@@...|
    |..@....|
    |..####.|
    +-------+
    
    Jet of gas pushes rock right:
    |...@...|
    |..@@@..|
    |...@...|
    |..####.|
    +-------+
    
    Rock falls 1 unit, causing it to come to rest:
    |...#...|
    |..###..|
    |...#...|
    |..####.|
    +-------+
    
    A new rock begins falling:
    |....@..|
    |....@..|
    |..@@@..|
    |.......|
    |.......|
    |.......|
    |...#...|
    |..###..|
    |...#...|
    |..####.|
    +-------+


The moment each of the next few rocks begins falling, you would see this:

    |..@....|
    |..@....|
    |..@....|
    |..@....|
    |.......|
    |.......|
    |.......|
    |..#....|
    |..#....|
    |####...|
    |..###..|
    |...#...|
    |..####.|
    +-------+
    
    |..@@...|
    |..@@...|
    |.......|
    |.......|
    |.......|
    |....#..|
    |..#.#..|
    |..#.#..|
    |#####..|
    |..###..|
    |...#...|
    |..####.|
    +-------+
    
    |..@@@@.|
    |.......|
    |.......|
    |.......|
    |....##.|
    |....##.|
    |....#..|
    |..#.#..|
    |..#.#..|
    |#####..|
    |..###..|
    |...#...|
    |..####.|
    +-------+
    
    |...@...|
    |..@@@..|
    |...@...|
    |.......|
    |.......|
    |.......|
    |.####..|
    |....##.|
    |....##.|
    |....#..|
    |..#.#..|
    |..#.#..|
    |#####..|
    |..###..|
    |...#...|
    |..####.|
    +-------+
    
    |....@..|
    |....@..|
    |..@@@..|
    |.......|
    |.......|
    |.......|
    |..#....|
    |.###...|
    |..#....|
    |.####..|
    |....##.|
    |....##.|
    |....#..|
    |..#.#..|
    |..#.#..|
    |#####..|
    |..###..|
    |...#...|
    |..####.|
    +-------+
    
    |..@....|
    |..@....|
    |..@....|
    |..@....|
    |.......|
    |.......|
    |.......|
    |.....#.|
    |.....#.|
    |..####.|
    |.###...|
    |..#....|
    |.####..|
    |....##.|
    |....##.|
    |....#..|
    |..#.#..|
    |..#.#..|
    |#####..|
    |..###..|
    |...#...|
    |..####.|
    +-------+
    
    |..@@...|
    |..@@...|
    |.......|
    |.......|
    |.......|
    |....#..|
    |....#..|
    |....##.|
    |....##.|
    |..####.|
    |.###...|
    |..#....|
    |.####..|
    |....##.|
    |....##.|
    |....#..|
    |..#.#..|
    |..#.#..|
    |#####..|
    |..###..|
    |...#...|
    |..####.|
    +-------+
    
    |..@@@@.|
    |.......|
    |.......|
    |.......|
    |....#..|
    |....#..|
    |....##.|
    |##..##.|
    |######.|
    |.###...|
    |..#....|
    |.####..|
    |....##.|
    |....##.|
    |....#..|
    |..#.#..|
    |..#.#..|
    |#####..|
    |..###..|
    |...#...|
    |..####.|
    +-------+


To prove to the elephants your simulation is accurate, they want to know how tall the tower will get after 2022 rocks have stopped (but before the 2023rd rock begins falling). In this example, the tower of rocks will be _`3068`_ units tall.

_How many units tall will the tower of rocks be after 2022 rocks have stopped falling?_

### Part Two

The elephants are not impressed by your simulation. They demand to know how tall the tower will be after _`1000000000000`_ rocks have stopped! Only then will they feel confident enough to proceed through the cave.

In the example above, the tower would be _`1514285714288`_ units tall!

_How tall will the tower be after `1000000000000` rocks have stopped?_

