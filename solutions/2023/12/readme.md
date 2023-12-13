# Day 12

## Thoughts

Our first day that requires memoization/dynamic programming! This one was rough, with dynamic programming being a weak spot of mine. I've only ever used it mostly on Advent of Code puzzles, not ever in real life and barely in school.

### Part 1

This part **can** be done with a brute force approach, but will be horribly unoptimized. My first thought was for dynamic programming, but I wasn't sure of the inputs for a while.

### Part 2

We are forced to resort to dynamic programming for the approach to be fast enough. Brute force approaches will not get a solution without a ton of processing.

1. Parse the input and copy the data.
2. For each record, sum the arrangement counts.
   1. For convenience with array checks, add an extra `.` **operational spring** to the front of the record of springs.
   2. Create a dynamic programming cache for each spring and the final result.
   3. Populate the first value in the cache to assume 1 possible arrangement.
   4. Seed the remainder of the cache up to the first `#` **damaged spring** with `1` possible arrangement. We no longer have flexibility due to seeing a `#` and we can potentially end up with no possible arrangements. And we start with the assumption that there might be at least `1` arrangement, but there might not be any, so the last part of the cache remains with `0`s.
   5. Finally, we can skip checking springs that have already met a previous size requirement. We can track this with the first index in the next cache that meets a size requirement. This starts at `1` due to the `.` spring we appended.
   6. Loop through the **size** of each contiguous group of damaged springs.
      1. Save space by creating the next cache.
      2. Start a counter for the group of **damaged springs**.
      3. Start by assuming that the size requirement has not been met.
      4. Loop through each spring.
         1. If we see a `#` or `?`, add to our group counter. We have 1 more potential **damaged spring**. Otherwise, set the counter to `0` because we are no longer in a contiguous group.
         2. If we don't see a `#`, increment the next cache position by the current one. This allows us to percolate the met size requirement's possible arrangement counter until we've seen another `#`. At that point, it's no longer possible to meet the size requirement alone and following groups of `#` must meet a subsequent size requirement.
         3. If our group counter is at least the expected size and the start to the group is not a `#`, meaning it is either a `.` or `?` that can be converted to a `.`, increment the next cache position by the previous cache at the start of that group. This allows us to keep track of all the possible arrangements for the size being evaluated. It also ensures that we don't keep tracking additional arrangements when our size requirement has already been met.
         4. If we've seen our first non-zero counter of possible arrangements, and we haven't already done so, set the next size requirement starting index to skip over where we've processed so far. This is a slight optimization that keeps us from processing springs that have already met a previous size requirement.
      5. Set the cache to the next cache.
   7. Our cache now holds the full arrangement count in the last value, return.

#### Example 1

Dynamic programming is often hard for me to wrap my head around, so let's take a look at a full example.

For example, for `???.### 1,1,3`, we end up with the following caches:

```
???.### 1,1,3

. ? ? ? . # # # ^
1 1 1 1 1 1 0 0 0 <- seed cache
0 0 1 2 3 3 1 0 0 <- size 1
0 0 0 0 1 1 3 0 0 <- size 1
0 0 0 0 0 0 0 0 1 <- size 3
```

We have 3 possible arrangements for 1 **damaged spring** in `???`: `#..`, `.#.`, and `..#`. Our associated cache values will then be `0 0 1 2 3`. The first cache position is the `.` we prefixed our record with. The second cache position ends up getting no value because we move our first known result to the next position in the cache. So if our full record was `? 1`, we'd end up with `1` possibility and final cache of `0 0 1`.

With each cache iteration, we are kicking our cache values over because we've already satisfied a size requirement and don't want to count a previously-accounted-for group with the next size requirement.

Going further with `???.### 1,1,3`, our first full cache for the first `1` is: 

```
. ? ? ? . # # # ^
0 0 1 2 3 3 1 0 0
          ^
```

The `3` pointed out above comes from step `2.v.c.b` where we kick the counter to the end. We no longer meet the group size requirement of `1` (it's `0` due to hitting a `.`). The next `1` in the cache gets populated from step `2.v.c.c` where we see one more possibility from the previous cache. But then our group size is too large. So the last two positions in the cache do not get populated because they do not satisfy our current size requirement. Notice, if the original record is `???.### 1`, the result would be `0` because there are no possible arrangements that could satisfy the full size requirements. It would mean there must be only `1` group of contiguous, **damaged springs**.

Now we are at the second size requirement, also `1`. Let's take a look at the cache:

```
      v
. ? ? ? . # # # ^
0 0 0 0 1 1 3 0 0
```

The first value we consider is the first `?` because the first two `??` were accounted for by the first size requirement. So the position for `.` gets our first arrangement for the second size. We aren't yet at the `#` while we evaluate the `.` so we copy the count of arrangements into the next cache position. When we reach the first `#`, we've met our size requirement and populate the cache with the previous cache value. 

Now we are at the third and final size requirement, `3`. Let's take a look at the cache:

```
          v
. ? ? ? . # # # ^
0 0 0 0 0 0 0 0 1
```

Woah! We're at the end! Let's dive into the details to understand what happened. We know previous positions have met size requirements, so we jump to the first `#`. The loop continues until it sees the third `#` which satisfies the requirement. We then take the value from the previous cache just before the first `#` at:

```
. ? ? ? . # # # ^
0 0 0 0 1 1 3 0 0
        ^
```

This is where we chain all the requirements together and end up with a final result of `1` possible arrangement.

#### Example 2

Let's walk through a more complicated example: `.??..??...?##. 1,1,3`. Here's what the resulting cache values look like.

```
.??..??...?##. 1,1,3

. . ? ? . . ? ? . . . ? # # . ^
1 1 1 1 1 1 1 1 1 1 1 1 1 0 0 0 <- seed cache
0 0 0 1 2 2 2 3 4 4 4 4 5 1 0 0 <- size 1
0 0 0 0 0 0 0 2 4 4 4 4 8 4 0 0 <- size 1
0 0 0 0 0 0 0 0 0 0 0 0 0 0 4 4 <- size 3
```

We seed the cache until the first `#`.

We then look for all possibilities that satisfy the first size requirement of `1`. Each pair of `??` can either be `#.` or `.#`, so they each produce two possibilities each. The final `?` gets us one more possibility, and we end after seeing the first `#`, no longer satisfying the requirement.

The next size requirement, also `1` recognizes the same thing, but the first group of `??` have already been accounted for, so we don't start evaluating until the `.` right after the first `??` group. We end up with `4` total possibilities between the two initial size requirements, but they don't satisfy the entire set of springs, so the result is still `0`.

The final size requirement, `3`, then doesn't find a group that satisfies it until it hits `?##` and records the value from the previous size requirements, which is `4`. This percolates to the next position since the next spring is **operational** and we get our final result of `4` possible arrangements.

#### Example 3

The final example is perhaps the most interesting with so many **unknown springs**.

```
?###???????? 3,2,1

. ? # # # ? ? ? ? ? ? ? ? ^
1 1 1 0 0 0 0 0 0 0 0 0 0 0  <- seed cache
0 0 0 0 1 1 1 1 1 1 1 1 1 1  <- size 3
0 0 0 0 0 0 0 0 1 2 3 4 5 6  <- size 2
0 0 0 0 0 0 0 0 0 0 1 3 6 10 <- size 1
```

We seed the cache until the first `#`. 

Then we scan for groups of size `3`. We see our first possibility at `?##`, then again at `###`. At this point, our size requirement has already been met, and must be satisfied by the three `###`, so we continue to percolate the `1` possible arrangement to the end, now that all the `?` can be `.`.

Now, we take a look at size requirement `2`. The first group that meets the requirement is `??` at indexes `6` and `7`. We take the value from the previous cache, and have `1` possible arrangement so far that satisfies the first two requirements. But after that, with each extra `?`, we get another possible arrangement because we can shift the two possible **damaged** springs all the way until the end. So we end up with `6` total possibilities.

Then we look at the final size requirement of `1`, starting at index `9`. We start by finally, fully satisfying the size requirements, so the first count is `1`. After that, we know that there could have been `2` possibilities with the second requirement, and the following `?` means our final requirement can have an additional possibility, so that leaves us with `3` arrangements. Then we double to `6` due to the same reasons as the last. And finally, we end up with `10` total arrangements.

#### Example 4

Here's what the caches look like for all the other sample inputs.

```
?#?#?#?#?#?#?#? 1,3,1,6

. ? # ? # ? # ? # ? # ? # ? # ? ^
1 1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 <- seed cache
0 0 1 1 1 0 0 0 0 0 0 0 0 0 0 0 0 <- size 1
0 0 0 0 0 0 0 1 1 0 0 0 0 0 0 0 0 <- size 3
0 0 0 0 0 0 0 0 0 1 1 0 0 0 0 0 0 <- size 1
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 <- size 6
```

```
????.#...#... 4,1,1

. ? ? ? ? . # . . . # . . . ^
1 1 1 1 1 1 1 0 0 0 0 0 0 0 0 <- seed cache
0 0 0 0 0 1 1 0 0 0 0 0 0 0 0 <- size 4
0 0 0 0 0 0 0 1 1 1 1 0 0 0 0 <- size 1
0 0 0 0 0 0 0 0 0 0 0 1 1 1 1 <- size 1
```

```
????.######..#####. 1,6,5

. ? ? ? ? . # # # # # # . . # # # # # . ^
1 1 1 1 1 1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 <- seed cache
0 0 1 2 3 4 4 1 0 0 0 0 0 0 0 0 0 0 0 0 0 <- size 1
0 0 0 0 0 0 0 0 0 0 0 0 4 4 4 0 0 0 0 0 0 <- size 6
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 4 4 <- size 5
```
