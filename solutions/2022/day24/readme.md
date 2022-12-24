# Day 24

## Thoughts

I was initially worried that this was going to be another deep dive into search algorithms. Fortunately, it was a simple breadth-first search!

### Part 1

1. Parse the valley.
   1. Store the walls as a set of points.
   2. Store the blizzards as a set of vectors.
   3. Store the height and width of the valley.
   4. Store the start and end points.
2. Find the minimum number of minutes required to pass through the valley, using a Breadth-First Search with our starting point as the initial queue state and until the queue is empty.
   1. Update the time.
   2. Update the blizzards to move in their respective directions, respecting wrapping around the valley.
      - For this to account for the walls, we need to mod by `size - 2` and always add `1` to the result to offset the first wall and subtract before mod to account for the last wall.
      - The modulo function needs to account for negative values and wrap. I wrote my own function for this since JavaScript does not support it.
      - The resulting operation looks like this: `1 + mod(x + time * dx - 1, width - 2)`;
   3. Determine the next positions to explore in the search.
      - For each direction (up, down, left, right, and stay).
        - Only add if it is in bounds and does not collide with a wall or blizzard.
   4. Check if we've reached the goal.

### Part 2

We can easily adapt Part 1 to satisfy Part 2.

1. Rather than a single goal, start with 3 (the end, start, and end).
2. If we've reached the current goal, erase the queue and add the next goal.

## Time

This is how long it took me to complete each part.

- Part 1: 45:28
- Part 2: 2:56

## Puzzle

# Day 24: Blizzard Basin

[https://adventofcode.com/2022/day/24](https://adventofcode.com/2022/day/24)

## Description

### Part One

With everything replanted for next year (and with elephants and monkeys to tend the grove), you and the Elves leave for the extraction point.

Partway up the mountain that shields the grove is a flat, open area that serves as the extraction point. It's a bit of a climb, but nothing the expedition can't handle.

At least, that would normally be true; now that the mountain is covered in snow, things have become more difficult than the Elves are used to.

As the expedition reaches a valley that must be traversed to reach the extraction site, you find that strong, turbulent winds are pushing small _blizzards_ of snow and sharp ice around the valley. It's a good thing everyone packed warm clothes! To make it across safely, you'll need to find a way to avoid them.

Fortunately, it's easy to see all of this from the entrance to the valley, so you make a map of the valley and the blizzards (your puzzle input). For example:

    #.#####
    #.....#
    #>....#
    #.....#
    #...v.#
    #.....#
    #####.#


The walls of the valley are drawn as `#`; everything else is ground. Clear ground - where there is currently no blizzard - is drawn as `.`. Otherwise, blizzards are drawn with an arrow indicating their direction of motion: up (`^`), down (`v`), left (`<`), or right (`>`).

The above map includes two blizzards, one moving right (`>`) and one moving down (`v`). In one minute, each blizzard moves one position in the direction it is pointing:

    #.#####
    #.....#
    #.>...#
    #.....#
    #.....#
    #...v.#
    #####.#


Due to <span title="I think, anyway. Do I look like a theoretical blizzacist?">conservation of blizzard energy</span>, as a blizzard reaches the wall of the valley, a new blizzard forms on the opposite side of the valley moving in the same direction. After another minute, the bottom downward-moving blizzard has been replaced with a new downward-moving blizzard at the top of the valley instead:

    #.#####
    #...v.#
    #..>..#
    #.....#
    #.....#
    #.....#
    #####.#


Because blizzards are made of tiny snowflakes, they pass right through each other. After another minute, both blizzards temporarily occupy the same position, marked `2`:

    #.#####
    #.....#
    #...2.#
    #.....#
    #.....#
    #.....#
    #####.#


After another minute, the situation resolves itself, giving each blizzard back its personal space:

    #.#####
    #.....#
    #....>#
    #...v.#
    #.....#
    #.....#
    #####.#


Finally, after yet another minute, the rightward-facing blizzard on the right is replaced with a new one on the left facing the same direction:

    #.#####
    #.....#
    #>....#
    #.....#
    #...v.#
    #.....#
    #####.#


This process repeats at least as long as you are observing it, but probably forever.

Here is a more complex example:

    #.######
    #>>.<^<#
    #.<..<<#
    #>v.><>#
    #<^v^^>#
    ######.#


Your expedition begins in the only non-wall position in the top row and needs to reach the only non-wall position in the bottom row. On each minute, you can _move_ up, down, left, or right, or you can _wait_ in place. You and the blizzards act _simultaneously_, and you cannot share a position with a blizzard.

In the above example, the fastest way to reach your goal requires _`18`_ steps. Drawing the position of the expedition as `E`, one way to achieve this is:

    Initial state:
    #E######
    #>>.<^<#
    #.<..<<#
    #>v.><>#
    #<^v^^>#
    ######.#
    
    Minute 1, move down:
    #.######
    #E>3.<.#
    #<..<<.#
    #>2.22.#
    #>v..^<#
    ######.#
    
    Minute 2, move down:
    #.######
    #.2>2..#
    #E^22^<#
    #.>2.^>#
    #.>..<.#
    ######.#
    
    Minute 3, wait:
    #.######
    #<^<22.#
    #E2<.2.#
    #><2>..#
    #..><..#
    ######.#
    
    Minute 4, move up:
    #.######
    #E<..22#
    #<<.<..#
    #<2.>>.#
    #.^22^.#
    ######.#
    
    Minute 5, move right:
    #.######
    #2Ev.<>#
    #<.<..<#
    #.^>^22#
    #.2..2.#
    ######.#
    
    Minute 6, move right:
    #.######
    #>2E<.<#
    #.2v^2<#
    #>..>2>#
    #<....>#
    ######.#
    
    Minute 7, move down:
    #.######
    #.22^2.#
    #<vE<2.#
    #>>v<>.#
    #>....<#
    ######.#
    
    Minute 8, move left:
    #.######
    #.<>2^.#
    #.E<<.<#
    #.22..>#
    #.2v^2.#
    ######.#
    
    Minute 9, move up:
    #.######
    #<E2>>.#
    #.<<.<.#
    #>2>2^.#
    #.v><^.#
    ######.#
    
    Minute 10, move right:
    #.######
    #.2E.>2#
    #<2v2^.#
    #<>.>2.#
    #..<>..#
    ######.#
    
    Minute 11, wait:
    #.######
    #2^E^2>#
    #<v<.^<#
    #..2.>2#
    #.<..>.#
    ######.#
    
    Minute 12, move down:
    #.######
    #>>.<^<#
    #.<E.<<#
    #>v.><>#
    #<^v^^>#
    ######.#
    
    Minute 13, move down:
    #.######
    #.>3.<.#
    #<..<<.#
    #>2E22.#
    #>v..^<#
    ######.#
    
    Minute 14, move right:
    #.######
    #.2>2..#
    #.^22^<#
    #.>2E^>#
    #.>..<.#
    ######.#
    
    Minute 15, move right:
    #.######
    #<^<22.#
    #.2<.2.#
    #><2>E.#
    #..><..#
    ######.#
    
    Minute 16, move right:
    #.######
    #.<..22#
    #<<.<..#
    #<2.>>E#
    #.^22^.#
    ######.#
    
    Minute 17, move down:
    #.######
    #2.v.<>#
    #<.<..<#
    #.^>^22#
    #.2..2E#
    ######.#
    
    Minute 18, move down:
    #.######
    #>2.<.<#
    #.2v^2<#
    #>..>2>#
    #<....>#
    ######E#


_What is the fewest number of minutes required to avoid the blizzards and reach the goal?_

### Part Two

As the expedition reaches the far side of the valley, one of the Elves looks especially dismayed:

He _forgot his snacks_ at the entrance to the valley!

Since you're so good at dodging blizzards, the Elves humbly request that you go back for his snacks. From the same initial conditions, how quickly can you make it from the start to the goal, then back to the start, then back to the goal?

In the above example, the first trip to the goal takes `18` minutes, the trip back to the start takes `23` minutes, and the trip back to the goal again takes `13` minutes, for a total time of _`54`_ minutes.

_What is the fewest number of minutes required to reach the goal, go back to the start, then reach the goal again?_

