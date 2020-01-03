## Day 7

### Part 1

This one was awesome! We're given wires with instructions for logical gates and we have to compute the signal for a given wire. I felt like this was the first one that didn't have an apparent algorithm. I feel like my initial attempt is pretty ugly -- I've been trying to just crank through these without making them perfect on the first round.

I started off with regexes to parse the expressions, but soon realized that there were more variations to the expressions, creating large if-else chains.

I then wrote a recursive function for evaluating the signal for a given wire.

I got into what appeared to either be infinite recursion or just a ton of unnecessary processing and the program just hung. I then added a type of cache for signals we've already computed and we pull from there instead if the value exists already.

I'm sure there is a lot of cleanup that can happen in the expression parsing and evaluating. Maybe I'll clean it up someday or in Part 2.
 