# Day 16

## Thoughts

Wow, this one was definitely hardest so far. I had to come back to this one over and over and finally got it on Day 21. Days 16-19 really wore me out with the need for a lot of advanced searching algorithms.

### Part 1

1. Parse the valves into a graph-map-like structure. Each valve should keep track of its name, rate, and other valves it leads to.
2. (Huge optimization). Precompute all the distances from every valve to every other valve. This is called the Floyd Warshall algorithm.
3. Get all the working valves. Because there is no point in opening valves that do not release any pressure.
4. Generate all the paths to open valves. This is essentially permutations, but we want to do some pruning. There are a few optimizations here that are critical to pruning the search space.
   1. (Tiny optimization). Use a generator, so we don't have to deep copy the arrays.
   2. Keep track of the cost to travel to a valve then open it.
   3. (Huge optimization). Only continue adding paths if the cost is left than the time remaining.
   4. (Huge optimization). Do not add cycles. We don't want to revisit a valve.
5. Keep track of the max pressure flow achievable by the various paths.

### Part 2

1. The first 3 steps are the same as Part 1. Below are optimizations.
2. (Huge optimization). Keep track of the best possible pressure release for each combination (prune permutations to combinations with a set). For example, if `CC -> EE -> JJ` releases `471` pressure, but `JJ -> EE -> CC` releases `575`, we only care about `JJ -> EE -> CC` because the other permutations are not going to be any better. I sort these and store them in a map with key `JJ-EE-CC` and the value of the max pressure release for that combination.
3. (Small optimization). We can sort the pressure release in descending order. It's likely that we've found the maximum earlier, so no need to keep trying other paths.
4. Now we iterate over the best releases we've sorted. The human always takes the higher value, then the elephant takes whatever is left.
   1. (Small optimization). If two times the human's value is less than the current max, don't bother exploring anymore. Nothing is going to be as good as what we've found.

## Time

This is how long it took me to complete each part. Yeah... I lost track of these ones...

- Part 1: 4+ hours
- Part 2: 6+ hours

## Puzzle

# Day 16: Proboscidea Volcanium

[https://adventofcode.com/2022/day/16](https://adventofcode.com/2022/day/16)

## Description

### Part One

The sensors have led you to the origin of the distress signal: yet another handheld device, just like the one the Elves gave you. However, you don't see any Elves around; instead, the device is surrounded by elephants! They must have gotten lost in these tunnels, and one of the elephants apparently figured out how to turn on the distress signal.

The ground rumbles again, much stronger this time. What kind of cave is this, exactly? You scan the cave with your handheld device; it reports mostly igneous rock, some ash, pockets of pressurized gas, magma... this isn't just a cave, it's a volcano!

You need to get the elephants out of here, quickly. Your device estimates that you have _30 minutes_ before the volcano erupts, so you don't have time to go back out the way you came in.

You scan the cave for other options and discover a network of pipes and pressure-release _valves_. You aren't sure how such a system got into a volcano, but you don't have time to complain; your device produces a report (your puzzle input) of each valve's _flow rate_ if it were opened (in pressure per minute) and the tunnels you could use to move between the valves.

There's even a valve in the room you and the elephants are currently standing in labeled `AA`. You estimate it will take you one minute to open a single valve and one minute to follow any tunnel from one valve to another. What is the most pressure you could release?

For example, suppose you had the following scan output:

    Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
    Valve BB has flow rate=13; tunnels lead to valves CC, AA
    Valve CC has flow rate=2; tunnels lead to valves DD, BB
    Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
    Valve EE has flow rate=3; tunnels lead to valves FF, DD
    Valve FF has flow rate=0; tunnels lead to valves EE, GG
    Valve GG has flow rate=0; tunnels lead to valves FF, HH
    Valve HH has flow rate=22; tunnel leads to valve GG
    Valve II has flow rate=0; tunnels lead to valves AA, JJ
    Valve JJ has flow rate=21; tunnel leads to valve II


All of the valves begin _closed_. You start at valve `AA`, but it must be damaged or <span title="Wait, sir! The valve, sir! it appears to be... jammed!">jammed</span> or something: its flow rate is `0`, so there's no point in opening it. However, you could spend one minute moving to valve `BB` and another minute opening it; doing so would release pressure during the remaining _28 minutes_ at a flow rate of `13`, a total eventual pressure release of `28 * 13 = 364`. Then, you could spend your third minute moving to valve `CC` and your fourth minute opening it, providing an additional _26 minutes_ of eventual pressure release at a flow rate of `2`, or _`52`_ total pressure released by valve `CC`.

Making your way through the tunnels like this, you could probably open many or all of the valves by the time 30 minutes have elapsed. However, you need to release as much pressure as possible, so you'll need to be methodical. Instead, consider this approach:

    == Minute 1 ==
    No valves are open.
    You move to valve DD.
    
    == Minute 2 ==
    No valves are open.
    You open valve DD.
    
    == Minute 3 ==
    Valve DD is open, releasing 20 pressure.
    You move to valve CC.
    
    == Minute 4 ==
    Valve DD is open, releasing 20 pressure.
    You move to valve BB.
    
    == Minute 5 ==
    Valve DD is open, releasing 20 pressure.
    You open valve BB.
    
    == Minute 6 ==
    Valves BB and DD are open, releasing 33 pressure.
    You move to valve AA.
    
    == Minute 7 ==
    Valves BB and DD are open, releasing 33 pressure.
    You move to valve II.
    
    == Minute 8 ==
    Valves BB and DD are open, releasing 33 pressure.
    You move to valve JJ.
    
    == Minute 9 ==
    Valves BB and DD are open, releasing 33 pressure.
    You open valve JJ.
    
    == Minute 10 ==
    Valves BB, DD, and JJ are open, releasing 54 pressure.
    You move to valve II.
    
    == Minute 11 ==
    Valves BB, DD, and JJ are open, releasing 54 pressure.
    You move to valve AA.
    
    == Minute 12 ==
    Valves BB, DD, and JJ are open, releasing 54 pressure.
    You move to valve DD.
    
    == Minute 13 ==
    Valves BB, DD, and JJ are open, releasing 54 pressure.
    You move to valve EE.
    
    == Minute 14 ==
    Valves BB, DD, and JJ are open, releasing 54 pressure.
    You move to valve FF.
    
    == Minute 15 ==
    Valves BB, DD, and JJ are open, releasing 54 pressure.
    You move to valve GG.
    
    == Minute 16 ==
    Valves BB, DD, and JJ are open, releasing 54 pressure.
    You move to valve HH.
    
    == Minute 17 ==
    Valves BB, DD, and JJ are open, releasing 54 pressure.
    You open valve HH.
    
    == Minute 18 ==
    Valves BB, DD, HH, and JJ are open, releasing 76 pressure.
    You move to valve GG.
    
    == Minute 19 ==
    Valves BB, DD, HH, and JJ are open, releasing 76 pressure.
    You move to valve FF.
    
    == Minute 20 ==
    Valves BB, DD, HH, and JJ are open, releasing 76 pressure.
    You move to valve EE.
    
    == Minute 21 ==
    Valves BB, DD, HH, and JJ are open, releasing 76 pressure.
    You open valve EE.
    
    == Minute 22 ==
    Valves BB, DD, EE, HH, and JJ are open, releasing 79 pressure.
    You move to valve DD.
    
    == Minute 23 ==
    Valves BB, DD, EE, HH, and JJ are open, releasing 79 pressure.
    You move to valve CC.
    
    == Minute 24 ==
    Valves BB, DD, EE, HH, and JJ are open, releasing 79 pressure.
    You open valve CC.
    
    == Minute 25 ==
    Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.
    
    == Minute 26 ==
    Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.
    
    == Minute 27 ==
    Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.
    
    == Minute 28 ==
    Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.
    
    == Minute 29 ==
    Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.
    
    == Minute 30 ==
    Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.


This approach lets you release the most pressure possible in 30 minutes with this valve layout, _`1651`_.

Work out the steps to release the most pressure in 30 minutes. _What is the most pressure you can release?_

### Part Two

You're worried that even with an optimal approach, the pressure released won't be enough. What if you got one of the elephants to help you?

It would take you 4 minutes to teach an elephant how to open the right valves in the right order, leaving you with only _26 minutes_ to actually execute your plan. Would having two of you working together be better, even if it means having less time? (Assume that you teach the elephant before opening any valves yourself, giving you both the same full 26 minutes.)

In the example above, you could teach the elephant to help you as follows:

    == Minute 1 ==
    No valves are open.
    You move to valve II.
    The elephant moves to valve DD.
    
    == Minute 2 ==
    No valves are open.
    You move to valve JJ.
    The elephant opens valve DD.
    
    == Minute 3 ==
    Valve DD is open, releasing 20 pressure.
    You open valve JJ.
    The elephant moves to valve EE.
    
    == Minute 4 ==
    Valves DD and JJ are open, releasing 41 pressure.
    You move to valve II.
    The elephant moves to valve FF.
    
    == Minute 5 ==
    Valves DD and JJ are open, releasing 41 pressure.
    You move to valve AA.
    The elephant moves to valve GG.
    
    == Minute 6 ==
    Valves DD and JJ are open, releasing 41 pressure.
    You move to valve BB.
    The elephant moves to valve HH.
    
    == Minute 7 ==
    Valves DD and JJ are open, releasing 41 pressure.
    You open valve BB.
    The elephant opens valve HH.
    
    == Minute 8 ==
    Valves BB, DD, HH, and JJ are open, releasing 76 pressure.
    You move to valve CC.
    The elephant moves to valve GG.
    
    == Minute 9 ==
    Valves BB, DD, HH, and JJ are open, releasing 76 pressure.
    You open valve CC.
    The elephant moves to valve FF.
    
    == Minute 10 ==
    Valves BB, CC, DD, HH, and JJ are open, releasing 78 pressure.
    The elephant moves to valve EE.
    
    == Minute 11 ==
    Valves BB, CC, DD, HH, and JJ are open, releasing 78 pressure.
    The elephant opens valve EE.
    
    (At this point, all valves are open.)
    
    == Minute 12 ==
    Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.
    
    ...
    
    == Minute 20 ==
    Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.
    
    ...
    
    == Minute 26 ==
    Valves BB, CC, DD, EE, HH, and JJ are open, releasing 81 pressure.


With the elephant helping, after 26 minutes, the best you could do would release a total of _`1707`_ pressure.

_With you and an elephant working together for 26 minutes, what is the most pressure you could release?_

