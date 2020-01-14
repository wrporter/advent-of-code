## Day 19

### Part 1

Interesting chemical-looking problem. I started off by storing the replacement values in a map of strings to a list of strings. Then I loop over every replacement string and add all the distinct molecules to a map, counting how many times each appears, then return the size (how many keys there are) of the map.
