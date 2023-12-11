# Day 11

## Thoughts

This was a nice breather from the previous day. I'm still sick and my 5-year-old son was up a lot through the night with fever and pink eye. So after enough wake-ups, I took an early-morning look at the puzzle.

### Part 1

My initial thought was that a naive approach would be to do a BFS between the points to get the distances, but I immediately jumped to being able to calculate the Manhattan distance between all the pairs of galaxies to drastically optimize.

My next thoughts on performance were to start by pre-calculating the y-gap and x-gap spaces. Then, factor in the gaps while creating the galaxy coordinates. 

For a final optimization, I wondered if I could just compute the gaps while creating the coordinates. Unfortunately, we can only do this with one or the other because we only traverse a single plane at a time while gathering galaxies. Here's what that would look like, but I feel like it sacrifices some readability by combining concerns:

```go
yGap := 0
for y, line := range image {
    yIsEmpty := true

    for x, char := range line {
        if char == '#' {
            yIsEmpty = false
            galaxies = append(galaxies, geometry.NewPoint(x+xGaps[x], y+yGap))
        }
    }

    if yIsEmpty {
        yGap += addGap
    }
}
```

1. Parse the input.
   1. Split the input by newlines.
   2. Calculate the y-gaps and x-gaps for each index.
      1. Set the initial gap size to 0.
      2. Every time we see a row or column that is empty, add to the total gap size. Keep track of every gap at the associated y and x coordinate.
   3. Finally, collect the galaxy coordinates. Add the y-gap and x-gap at each coordinate.
2. For loop over pairs by using a nested for loop of the current position + 1. Get the Manhattan distance between points and add to the total sum.

### Part 2

This part required a little refactoring to account for the "multiplied" rather than "added" space. So our Part 1, really multiplies space by 2 when adding a single gap. The process is then the same as Part 1, but with a gap multiplier of 1,000,000.
