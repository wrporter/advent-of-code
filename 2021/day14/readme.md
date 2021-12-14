# Day 14: Extended Polymerization

[https://adventofcode.com/2021/day/14](https://adventofcode.com/2021/day/14)

## Thoughts & Solution

I was up later than usual and decided to hop on for the puzzle. I finished Part 1, then realized I would need to do a lot more optimization via memoization for Part 2, so I just went to bed and finished the next morning.

Part 1: I started off with a basic approach of keeping track of the polymer template and making it grow over time, by applying the rules and inserting new elements. I kept track of counts along the way as an optimization. Part 1 ran in 7ms. Not too shabby.

Part 2: Yeah... my approach to Part 1 wasn't going to fly. We are on a whole new scale. I printed out the min and max differences for each step to see if there was a common increasing pattern, and although there wasn't, I realized that our template was growing at an exponential rate. It was no longer feasible to keep track of the actual template as it grew. So instead, I realized I could keep track of all the counts of pairs and elements along the way. Then for each existing pair and associated count, perform the following.

1. Get the element to insert from the rules.
2. Decrease the counter for the pair because we are going to split that pair and it will no longer exist after this specific insertion.
3. Increase the counter for the first character of the pair, followed by the element to insert.
4. Increase the counter for the element to insert, followed by the second character of the pair.
5. Increase the counter for the element that just got inserted.

This was a pretty fun puzzle and a good lesson in memoization. It seems there are always a few puzzles each year of this flavor with a really high scale on Part 2.

### Time

- Part 1: 42:13
- Part 2: 41:55

## Description

### Part One

The incredible pressures at this depth are starting to put a strain on your submarine. The submarine has [polymerization](https://en.wikipedia.org/wiki/Polymerization) equipment that would produce suitable materials to reinforce the submarine, and the nearby volcanically-active caves should even have the necessary input elements in sufficient quantities.

The submarine manual contains <span title="HO

HO -> OH">instructions</span> for finding the optimal polymer formula; specifically, it offers a _polymer template_ and a list of _pair insertion_ rules (your puzzle input). You just need to work out what polymer would result after repeating the pair insertion process a few times.

For example:

    NNCB
    
    CH -> B
    HH -> N
    CB -> H
    NH -> C
    HB -> C
    HC -> B
    HN -> C
    NN -> C
    BH -> H
    NC -> B
    NB -> B
    BN -> B
    BB -> N
    BC -> B
    CC -> N
    CN -> C
    

The first line is the _polymer template_ - this is the starting point of the process.

The following section defines the _pair insertion_ rules. A rule like `AB -> C` means that when elements `A` and `B` are immediately adjacent, element `C` should be inserted between them. These insertions all happen simultaneously.

So, starting with the polymer template `NNCB`, the first step simultaneously considers all three pairs:

*   The first pair (`NN`) matches the rule `NN -> C`, so element _`C`_ is inserted between the first `N` and the second `N`.
*   The second pair (`NC`) matches the rule `NC -> B`, so element _`B`_ is inserted between the `N` and the `C`.
*   The third pair (`CB`) matches the rule `CB -> H`, so element _`H`_ is inserted between the `C` and the `B`.

Note that these pairs overlap: the second element of one pair is the first element of the next pair. Also, because all pairs are considered simultaneously, inserted elements are not considered to be part of a pair until the next step.

After the first step of this process, the polymer becomes `NCNBCHB`.

Here are the results of a few steps using the above rules:

    Template:     NNCB
    After step 1: NCNBCHB
    After step 2: NBCCNBBBCBHCB
    After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB
    After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB
    

This polymer grows quickly. After step 5, it has length 97; After step 10, it has length 3073. After step 10, `B` occurs 1749 times, `C` occurs 298 times, `H` occurs 161 times, and `N` occurs 865 times; taking the quantity of the most common element (`B`, 1749) and subtracting the quantity of the least common element (`H`, 161) produces `1749 - 161 = 1588`.

Apply 10 steps of pair insertion to the polymer template and find the most and least common elements in the result. _What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?_

### Part Two

The resulting polymer isn't nearly strong enough to reinforce the submarine. You'll need to run more steps of the pair insertion process; a total of _40 steps_ should do it.

In the above example, the most common element is `B` (occurring `2192039569602` times) and the least common element is `H` (occurring `3849876073` times); subtracting these produces _`2188189693529`_.

Apply _40_ steps of pair insertion to the polymer template and find the most and least common elements in the result. _What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?_
