# Day 23

## Thoughts

I really loved today's puzzle. This was such a fun and relieving day. I enjoy the problems that require a more well-thought-out design and solution according to a larger amount of constraints and requirements rather than clever deep dives into algorithms. This is yet another variation on Conway's Game of Life that Eric likes to throw in each year.

### Part 1

1. Parse the locations of the elves into a hash map. It is a pain to maintain this kind of state with a 2D array. I suggest using a set or hash map.
2. Simulate moving the elves for each round.
   1. Set the first direction to `0` or `North`. Whichever way you might be keeping track of the direction groupings.
   2. First half: Determine where each elf can move.
      1. For each direction make sure there are no elves in that direction. If there are not, then move the elf in that direction. I do this by keeping track of the elves' positions in a hashmap with a list of elves. This way, I can determine later if they tried to move to the same spot. Because of this, we also need to keep track of where the elves came from so that we can leave them where they are, so they don't move if they aren't allowed to.
      2. Otherwise, if the elf could not move, add it to the next set of elves with its current position.
   3. Second half: Now, copy over any elves that did not move. If there is only 1 elf that tried to move to a spot, let it move to that spot and copy it over to the new set.
   4. Assign the new set to the current view of the map.
   5. Update the first direction to be the next direction, wrapping with modulo of the number of direction groups (4).
3. Create a grid out of the set of elves.
   1. Find the bounding points, minimum and maximum x and y coordinates.
   2. From the bounds, determine the height and width of the grid.
   3. Now, for each position, create a 2D array and fill each empty spot with a `.` and each spot with an elf with `#`.
4. Sum all the empty spaces `.` and return the result!

### Part 2

Part 2 was a very simple modification to Part 1.

1. While simulating movement of elves, keep track of whether any elf has moved during the round. If no elf has moved, stop simulating.
2. Return the current round number.

## Time

This is how long it took me to complete each part.

- Part 1: 1:23:26
- Part 2: 3:35

## Animation!

I animated this one with Go and Ebitengine. Super fun. Looks so much like Game of Life.

![Elves Go Marching](elves.mp4)

## Puzzle

# Day 23: Unstable Diffusion

[https://adventofcode.com/2022/day/23](https://adventofcode.com/2022/day/23)

## Description

### Part One

You enter a large crater of gray dirt where the grove is supposed to be. All around you, plants you imagine were expected to be full of fruit are instead withered and broken. A large group of Elves has formed in the middle of the grove.

"...but this volcano has been dormant for months. Without ash, the fruit can't grow!"

You look up to see a massive, snow-capped mountain towering above you.

"It's not like there are other active volcanoes here; we've looked everywhere."

"But our scanners show active magma flows; clearly it's going _somewhere_."

They finally notice you at the edge of the grove, your pack almost overflowing from the random _star_ fruit you've been collecting. Behind you, elephants and monkeys explore the grove, looking concerned. Then, the Elves recognize the ash cloud slowly spreading above your recent detour.

"Why do you--" "How is--" "Did you just--"

Before any of them can form a complete question, another Elf speaks up: "Okay, new plan. We have almost enough fruit already, and ash from the plume should spread here eventually. If we quickly plant new seedlings now, we can still make it to the extraction point. Spread out!"

The Elves each reach into their pack and pull out a tiny plant. The plants rely on important nutrients from the ash, so they can't be planted too close together.

There isn't enough time to let the Elves figure out where to plant the seedlings themselves; you quickly scan the grove (your puzzle input) and note their positions.

For example:

    ....#..
    ..###.#
    #...#.#
    .#...##
    #.###..
    ##.#.##
    .#..#..


The scan shows Elves `#` and empty ground `.`; outside your scan, more empty ground extends a long way in every direction. The scan is oriented so that _north is up_; orthogonal directions are written N (north), S (south), W (west), and E (east), while diagonal directions are written NE, NW, SE, SW.

The Elves follow a time-consuming process to figure out where they should each go; you can speed up this process considerably. The process consists of some number of _rounds_ during which Elves alternate between considering where to move and actually moving.

During the _first half_ of each round, each Elf considers the eight positions adjacent to themself. If no other Elves are in one of those eight positions, the Elf _does not do anything_ during this round. Otherwise, the Elf looks in each of four directions in the following order and _proposes_ moving one step in the _first valid direction_:

*   If there is no Elf in the N, NE, or NW adjacent positions, the Elf proposes moving _north_ one step.
*   If there is no Elf in the S, SE, or SW adjacent positions, the Elf proposes moving _south_ one step.
*   If there is no Elf in the W, NW, or SW adjacent positions, the Elf proposes moving _west_ one step.
*   If there is no Elf in the E, NE, or SE adjacent positions, the Elf proposes moving _east_ one step.

After each Elf has had a chance to propose a move, the _second half_ of the round can begin. Simultaneously, each Elf moves to their proposed destination tile if they were the _only_ Elf to propose moving to that position. If two or more Elves propose moving to the same position, _none_ of those Elves move.

Finally, at the end of the round, the _first direction_ the Elves considered is moved to the end of the list of directions. For example, during the second round, the Elves would try proposing a move to the south first, then west, then east, then north. On the third round, the Elves would first consider west, then east, then north, then south.

As a smaller example, consider just these five Elves:

    .....
    ..##.
    ..#..
    .....
    ..##.
    .....


The northernmost two Elves and southernmost two Elves all propose moving north, while the middle Elf cannot move north and proposes moving south. The middle Elf proposes the same destination as the southwest Elf, so neither of them move, but the other three do:

    ..##.
    .....
    ..#..
    ...#.
    ..#..
    .....


Next, the northernmost two Elves and the southernmost Elf all propose moving south. Of the remaining middle two Elves, the west one cannot move south and proposes moving west, while the east one cannot move south _or_ west and proposes moving east. All five Elves succeed in moving to their proposed positions:

    .....
    ..##.
    .#...
    ....#
    .....
    ..#..


Finally, the southernmost two Elves choose not to move at all. Of the remaining three Elves, the west one proposes moving west, the east one proposes moving east, and the middle one proposes moving north; all three succeed in moving:

    ..#..
    ....#
    #....
    ....#
    .....
    ..#..


At this point, no Elves need to move, and so the process ends.

The larger example above proceeds as follows:

    == Initial State ==
    ..............
    ..............
    .......#......
    .....###.#....
    ...#...#.#....
    ....#...##....
    ...#.###......
    ...##.#.##....
    ....#..#......
    ..............
    ..............
    ..............
    
    == End of Round 1 ==
    ..............
    .......#......
    .....#...#....
    ...#..#.#.....
    .......#..#...
    ....#.#.##....
    ..#..#.#......
    ..#.#.#.##....
    ..............
    ....#..#......
    ..............
    ..............
    
    == End of Round 2 ==
    ..............
    .......#......
    ....#.....#...
    ...#..#.#.....
    .......#...#..
    ...#..#.#.....
    .#...#.#.#....
    ..............
    ..#.#.#.##....
    ....#..#......
    ..............
    ..............
    
    == End of Round 3 ==
    ..............
    .......#......
    .....#....#...
    ..#..#...#....
    .......#...#..
    ...#..#.#.....
    .#..#.....#...
    .......##.....
    ..##.#....#...
    ...#..........
    .......#......
    ..............
    
    == End of Round 4 ==
    ..............
    .......#......
    ......#....#..
    ..#...##......
    ...#.....#.#..
    .........#....
    .#...###..#...
    ..#......#....
    ....##....#...
    ....#.........
    .......#......
    ..............
    
    == End of Round 5 ==
    .......#......
    ..............
    ..#..#.....#..
    .........#....
    ......##...#..
    .#.#.####.....
    ...........#..
    ....##..#.....
    ..#...........
    ..........#...
    ....#..#......
    ..............


After a few more rounds...

    == End of Round 10 ==
    .......#......
    ...........#..
    ..#.#..#......
    ......#.......
    ...#.....#..#.
    .#......##....
    .....##.......
    ..#........#..
    ....#.#..#....
    ..............
    ....#..#..#...
    ..............


To make sure they're on the right track, the Elves like to check after round 10 that they're making good progress toward covering enough ground. To do this, count the number of empty ground tiles contained by the smallest rectangle that contains every Elf. (The edges of the rectangle should be aligned to the N/S/E/W directions; the Elves do not have the patience to calculate <span title="Arbitrary Rectangles is my Piet Mondrian cover band.">arbitrary rectangles</span>.) In the above example, that rectangle is:

    ......#.....
    ..........#.
    .#.#..#.....
    .....#......
    ..#.....#..#
    #......##...
    ....##......
    .#........#.
    ...#.#..#...
    ............
    ...#..#..#..


In this region, the number of empty ground tiles is _`110`_.

Simulate the Elves' process and find the smallest rectangle that contains the Elves after 10 rounds. _How many empty ground tiles does that rectangle contain?_

### Part Two

It seems you're on the right track. Finish simulating the process and figure out where the Elves need to go. How many rounds did you save them?

In the example above, the _first round where no Elf moved_ was round _`20`_:

    .......#......
    ....#......#..
    ..#.....#.....
    ......#.......
    ...#....#.#..#
    #.............
    ....#.....#...
    ..#.....#.....
    ....#.#....#..
    .........#....
    ....#......#..
    .......#......


Figure out where the Elves need to go. _What is the number of the first round where no Elf moves?_

