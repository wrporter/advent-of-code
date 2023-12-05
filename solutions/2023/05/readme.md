# Day 5

## Thoughts

These problems seem to be harder this early than previous years. I was actually awake when the puzzle came out at 10pm due to an alert (I am on-call this week), but after I skimmed the problem description, I said, "Nope, not staying up any later for this. Way too many instructions to parse through without sleep."

I thought about the problem a bit just before jumping into bed and had a rough idea at a solution. I knew, though, that parsing the input was going to take most of my time.

### Part 1

And sure enough, parsing did take most of my time, but wasn't too bad.

1. Loop through each category map.
2. Loop through each seed.
3. Loop through each source/destination conversion and update the seed values. As noted in the instructions, if a seed does not match a range, its value remains the same, so we don't have to worry about modifying those.
4. My input followed the same pattern with `location` being the final category, so I didn't bother keeping track of categories. I then take the minimum seed value and return that as my result.

### Part 2

I knew I was going to have to do some range math. A brute force solution would mean keeping track of 2,207,992,798 seeds! I let the code run while I started working on a better algorithm, and it finished in 5 minutes 46 seconds! 😅

Here's my brute force method, which took only a minute to write, then another few minutes to debug because I didn't read the instructions properly that the second values are the length and not the end of a range.

1. Create seeds for every value in the ranges.
2. Follow the steps from Part 1.

Overall, I was very lucky with my delta time between solving parts 1 and 2. Advent of Code shows it was 21m55s.

I'll commit the brute force solution now, with plans to return and clean things up.
