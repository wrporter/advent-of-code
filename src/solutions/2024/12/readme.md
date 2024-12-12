# Day 12

### Part 1

I used BFS and incremented a perimeter counter every time I saw a neighbor that went out of bounds or was a different type of plant.

### Part 2

This one was tricky! I had a solution in about 15 minutes that worked for all the samples. But... it was off-by-one for a single region on the actual input... Nooooo!! I tried to fix my solution by rotating around the grid to determine the one duplicate in the cycle. But that started to get super complicated.

So, I resorted to looking up hints on the subreddit and found that many people counted corners instead of sides, relying on this thing called **geometry** that states that for all polygons, the number of sides equals the number of corners. I accomplished this by rotating the direction and comparing against the 3 adjacent cells.
