# Day 15

## Thoughts

This was another nice breather from the hard days. I didn't start until pretty late because I've gotten sick, just like all my kids, several of which had a dentist appointment this morning, so we had to get them to that.

### Part 1

If you can make it through all the description text, this one's a breeze! It's on the level of a Day 1 puzzle.

1. Parse the input by simply splitting the string on commas `,`.
2. Initialize a total sum of `0`.
3. For each step in the sequence, calculate the hash and add it to the sum.
   - Calculate the hash just as the instructions indicate. Start with a value of `0`, then loop over each character.
     1. Add the ASCII code of the character to the value. In Go, this was very simple because indexing into a string is automatically converted into a rune/byte so we already have the ASCII code. This is the same in many other languages, except maybe JavaScript where you need to use [charCodeAt](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/charCodeAt).
     2. Multiply the value by `17`.
     3. Mod the value by `256`.
4. Return the final sum.

### Part 2

Again, description overload! Looking at the examples was much more helpful for me to understand what was going on.

1. Parse the sequence as before.
2. Set up a map of boxes to a list of lenses.
3. Perform the operations for each step.
   1. If the operation is an `=`, create a new lens, splitting the step by `=` for the label and the focal length. Get the box number by taking the hash of the label. If the label is already in the box, update its focal length to the new one. Otherwise, add it to the list of lenses in that box.
   2. Otherwise, operate with a `-`. Determine the box the label belongs to based on its hash. If the lens is in that box, remove it.
4. Calculate the total focusing power by looping through each box, then each lens in that box. Make sure to add `1` to the box and slot numbers because our indexes start at `0`.
5. Return the total power.
