## Day 24

### Part 1

My initial reaction was, "Ooo! Conway's Game of Life!!". Then I thought, "There's got to be some kind of trick, if not in part 1, undoubtedly in part 2." 

So I figured to start with a 2D grid of `byte`s then store each iteration in a flattened string form to conserve space and so I could index in a `map` for checking past iterations.

### Part 2

Of course Eris is an old Plutonian settlement and has to have recursively-folded space! My initial impression is that this does complicate things, but not by much. We just need to include a depth to our layouts.

I couldn't think of a better way to check recursive neighboring positions, so it's a good amount of logic. And I don't really need to add a new set of levels until every 2 minutes, so maybe I'll come back and fix it someday.
