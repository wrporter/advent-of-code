## Day 18

This one was a walk in the park compared to day18. Perhaps the biggest part was parsing the portals and the code is pretty hard to read. 

### Part 1

I added open passage positions to a map of points to their respective values. Then I maintain a separate map for the portal positions.

After that it's just a Breadth-First Search across neighboring, including portal, positions. 

### Part 2

This one took me longer than expected. I kept putting my level checking logic in the wrong spots, and looking at the diff, it was very simple. I was able to make the same code work for part 1 by passing down a flag that only adds 1 more if statement.
