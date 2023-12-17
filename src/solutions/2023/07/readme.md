# Day 7

## Thoughts

Surprisingly, I've never played poker, so I might just have an advantage by not having preconceived ideas on how the game is supposed to work 😂. 

### Part 1

As soon as I tried my initial solution, the sample input passed. I added more test cases, but they didn't catch the bugs. So I started debug printing the sorted hands and their hand types and manually inspected whether the hands arrived in the correct order. This was super helpful.

Two things that tripped me up during debugging:

1. I initially just took a random count and it simplified my logic when I kept track of the most occurrences of a card.
2. When sorting, I forgot to account for when the first card is stronger than the second card and was treating them as if they were equal.

My process:

1. Parse each hand and determine the type.
    1. Use a map to count each card type in the hand. Keep track of the most occurrences of a given card.
    2. Now we have a large if-else chain to determine the hand type. Conveniently, we can keep this in order of strength and as they are explained in the instructions.
2. Sort the hands based on their type. If they match, loop through and sort based on the strength of the individual cards. 
3. Total the winnings as instructed!

### Part 2

Where I kept messing up on this one was the nuanced logic to determine the hand type. When I counted the most occurrences of a card, I realized I needed to exclude the wild jokers to simplify my logic. Once I did that, it was far easier to think through.

Once I had a working solution, I refactored the code so I could easily reuse common parts. Here's what I changed from Part 1 to be compatible with Part 2:

1. Replace all `J`s with `W`s.
2. Add `W` to the bottom of my strength map.
3. When determining hand type, separate wilds to pair them with the card that occurs the most.
