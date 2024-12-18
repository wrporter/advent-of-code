# Day 9: Smoke Basin

[https://adventofcode.com/2021/day/9](https://adventofcode.com/2021/day/9)

## Thoughts & Solution

Part 1: At first, I thought we were supposed to get the lowest spots in each row. Gotta make sure to thoroughly read those instructions! But then I realized we just care about points that are lower than all their neighbors, so it was still simple, just traverse the whole grid and compare to find the lowest points.

Part 2: This is similar to an interview question some engineers at Qualtrics used to ask, but I've never solved it -- find the size of the largest island. We already found our starting points in Part 1, so we just need to expand out from there with a BFS or DFS search to get all the neighbors. I chose to use my "visited" set to track points in the basin and exclude anything that was 9 or higher. Then I sorted the sizes in descending order and multiplied the top 3. Bingo! 

### Time

- Part 1: 8:13
- Part 2: 14:38

## Description

### Part One

These caves seem to be [lava tubes](https://en.wikipedia.org/wiki/Lava_tube). Parts are even still volcanically active; small hydrothermal vents release smoke into the caves that slowly <span title="This was originally going to be a puzzle about watersheds, but we're already under water.">settles like rain</span>.

If you can model how the smoke flows through the caves, you might be able to avoid it and be that much safer. The submarine generates a heightmap of the floor of the nearby caves for you (your puzzle input).

Smoke flows to the lowest point of the area it's in. For example, consider the following heightmap:

    2199943210
    3987894921
    9856789892
    8767896789
    9899965678
    

Each number corresponds to the height of a particular location, where `9` is the highest and `0` is the lowest a location can be.

Your first goal is to find the _low points_ - the locations that are lower than any of its adjacent locations. Most locations have four adjacent locations (up, down, left, and right); locations on the edge or corner of the map have three or two adjacent locations, respectively. (Diagonal locations do not count as adjacent.)

In the above example, there are _four_ low points, all highlighted: two are in the first row (a `1` and a `0`), one is in the third row (a `5`), and one is in the bottom row (also a `5`). All other locations on the heightmap have some lower adjacent location, and so are not low points.

The _risk level_ of a low point is _1 plus its height_. In the above example, the risk levels of the low points are `2`, `1`, `6`, and `6`. The sum of the risk levels of all low points in the heightmap is therefore _`15`_.

Find all of the low points on your heightmap. _What is the sum of the risk levels of all low points on your heightmap?_

### Part Two

Next, you need to find the largest basins so you know what areas are most important to avoid.

A _basin_ is all locations that eventually flow downward to a single low point. Therefore, every low point has a basin, although some basins are very small. Locations of height `9` do not count as being in any basin, and all other locations will always be part of exactly one basin.

The _size_ of a basin is the number of locations within the basin, including the low point. The example above has four basins.

The top-left basin, size `3`:

    2199943210
    3987894921
    9856789892
    8767896789
    9899965678
    

The top-right basin, size `9`:

    2199943210
    3987894921
    9856789892
    8767896789
    9899965678
    

The middle basin, size `14`:

    2199943210
    3987894921
    9856789892
    8767896789
    9899965678
    

The bottom-right basin, size `9`:

    2199943210
    3987894921
    9856789892
    8767896789
    9899965678
    

Find the three largest basins and multiply their sizes together. In the above example, this is `9 * 14 * 9 = 1134`.

_What do you get if you multiply together the sizes of the three largest basins?_
