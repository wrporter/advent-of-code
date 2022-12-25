# Day 18

## Thoughts

This was a fun geometry problem. I was exhausted with Days 16-19 and took such a long time thinking about the more complex puzzles that it took me a while to come back to this one and finish out Part 2, but it wasn't too bad.

### Part 1

1. Parse the locations of the lava cubes and store them in a map where the key is a string of their coordinates.
2. For every cube, get all of its sides. For each side that is not touching another cube, add 1 unit to the surface area.

### Part 2

This one was similar, but we had to do a flood fill algorithm. How to go about this was not apparent to me at first. I kept thinking I needed my starting points for the flood fill to be every cube and expand out to see if they touch the external air. This was far more complicated. Later, I visualized it as if we sprayed water at the lava chunk or submerged it in water. This was by far a simpler solution.

1. Parse the lava as in Part 1. However, add a min and max of the coordinates with a buffer of 1 unit. We want to essentially wrap water around the lava.
2. Our starting point for the Depth-First Search flood fill can be the bottom-left-back corner, associated with our min values. Think of this like pouring water out over the lava from this corner and it's going to expand to all other parts, just barely surrounding the lava.
3. Continue until our queue is empty.
   1. Pop off a 3D point.
   2. For each side, ensure it is within our min and max bounds, is is only air, and we haven't visited it yet.
   3. Push all valid positions onto our queue.
   4. Add the point to our visited set.
4. Calculate the surface area very similar to before, only slightly different. We want to find all cube sides that the external air touches.

## Time

This is how long it took me to complete each part.

- Part 1: 21:37
- Part 2: 49:14

## Puzzle

# Day 18: Boiling Boulders

[https://adventofcode.com/2022/day/18](https://adventofcode.com/2022/day/18)

## Description

### Part One

You and the elephants finally reach fresh air. You've emerged near the base of a large volcano that seems to be actively erupting! Fortunately, the lava seems to be flowing away from you and toward the ocean.

Bits of lava are still being ejected toward you, so you're sheltering in the cavern exit a little longer. Outside the cave, you can see the lava landing in a pond and hear it loudly hissing as it solidifies.

Depending on the specific compounds in the lava and speed at which it cools, it might be forming [obsidian](https://en.wikipedia.org/wiki/Obsidian)! The cooling rate should be based on the surface area of the lava droplets, so you take a quick scan of a droplet as it flies past you (your puzzle input).

Because of how quickly the lava is moving, the scan isn't very good; its resolution is quite low and, as a result, it approximates the shape of the lava droplet with _1x1x1 <span title="Unfortunately, you forgot your flint and steel in another dimension.">cubes</span> on a 3D grid_, each given as its `x,y,z` position.

To approximate the surface area, count the number of sides of each cube that are not immediately connected to another cube. So, if your scan were only two adjacent cubes like `1,1,1` and `2,1,1`, each cube would have a single side covered and five sides exposed, a total surface area of _`10`_ sides.

Here's a larger example:

    2,2,2
    1,2,2
    3,2,2
    2,1,2
    2,3,2
    2,2,1
    2,2,3
    2,2,4
    2,2,6
    1,2,5
    3,2,5
    2,1,5
    2,3,5


In the above example, after counting up all the sides that aren't connected to another cube, the total surface area is _`64`_.

_What is the surface area of your scanned lava droplet?_

### Part Two

Something seems off about your calculation. The cooling rate depends on exterior surface area, but your calculation also included the surface area of air pockets trapped in the lava droplet.

Instead, consider only cube sides that could be reached by the water and steam as the lava droplet tumbles into the pond. The steam will expand to reach as much as possible, completely displacing any air on the outside of the lava droplet but never expanding diagonally.

In the larger example above, exactly one cube of air is trapped within the lava droplet (at `2,2,5`), so the exterior surface area of the lava droplet is _`58`_.

_What is the exterior surface area of your scanned lava droplet?_

