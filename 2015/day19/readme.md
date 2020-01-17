## Day 19

### Part 1

Interesting chemical-looking problem. I started off by storing the replacement values in a map of strings to a list of strings. Then I loop over every replacement string and add all the distinct molecules to a map, counting how many times each appears, then return the size (how many keys there are) of the map.

### Part 2

Well bummer, this was the first one I had to look up. I couldn't think of a solution that wouldn't take too long or would take forever. Turns out that knowing organic chemistry gives you an advantage on this one. I do not know O-Chem. It also looks like all of the inputs were built with this in mind and can even be done by hand (although after a while due to the length of the molecule). We just traverse across all the outputs, replacing them with their respective inputs until we can break the molecule down to `e`.

What's unfortunate is that the examples provided (e.g. `HOHOHO`) do not follow the same paradigm and will not work in the algorithm.
