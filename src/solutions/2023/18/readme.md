# Day 18

## Thoughts

Oh, wow, this one really required learning some advanced math concepts. I looked up hints on the Reddit. I kept seeing [Green's Theorem](https://en.wikipedia.org/wiki/Green%27s_theorem), the [Shoelace Formula](https://en.wikipedia.org/wiki/Shoelace_formula) and [Pick's Theorem](https://en.wikipedia.org/wiki/Pick's_theorem) for calculating area, all of which are related. I really wanted to understand how these formulas worked. So I'll break it down.

The [Shoelace Formula](https://en.wikipedia.org/wiki/Shoelace_formula) uses the points, in a clockwise order, of a polygon and sums the determinants of a 2x2 matrix of each point, relative to the next point. This formula is slightly more complex than **Green's Theorem**, but requires lower-level understanding of math concepts.

[Green's Theorem](https://en.wikipedia.org/wiki/Green%27s_theorem) simplifies the area calculation in our example, as we deal with a rate of change of a curve. This was the next level of calculus that I never took in college, but I get the principle, associated with our simplified problem space of a polygon, while operating with integrals of the area under each line between points. This simplifies our calculations because we can now calculate the area as a series of integrals under each line rather than needing to keep track of each point, as in the **Shoelace Formula**.

[Pick's Theorem](https://en.wikipedia.org/wiki/Pick's_theorem) comes into play at the end of the puzzle because we need to include the perimeter in the overall area calculation. The formula is:

```math
A = i + \frac{b}{2} - 1
```

where $A$ is the inner area of the polygon, $i$ is the number of interior points, and $b$ is the number of boundary points (perimeter). 

We can calculate the area from **Green's Theorem** and the perimeter from the distance we travel with the digger. We then need to solve for the interior points.

```math
i = A - \frac{b}{2} + 1
```

In the puzzle, we are including the perimeter of the trench in our interior points calculation, so the formula becomes:

```math
\begin{align*}
i &= A - \frac{b}{2} + 1 + b \\
  &= A + \frac{b}{2} + 1
\end{align*}
```

### Part 1

1. Initialize our current point to $(0, 0)$, the perimeter to `0` and the area to `0`.
2. Loop through each step in the dig plan.
   1. Parse direction and distance to travel.
   2. Update our current point to the next point.
   3. Update our perimeter by the new distance.
   4. Use **Green's Theorem** to update the area by the rate of change of the line. E.g. $x * dy * distance$.
3. Apply **Pick's Theorem** by returning $A + \frac{b}{2} + 1$.

### Part 2

This is the same as Part 1, but now requires that we use math rather than any kind of brute force. I'm relieved that I spent all that time to understand those theorems 😅.

I refactored my code from Part 1 to pass in a parsing function to get the direction and distance. 
