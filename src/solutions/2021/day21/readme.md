# Day 21: Dirac Dice

[https://adventofcode.com/2021/day/21](https://adventofcode.com/2021/day/21)

## Thoughts & Solution

Part 1: This was pretty fun, simulating the dice game. I liked maintaining an object for the dice and updating its current roll value. Very straightforward.

Part 2: The last few days have been exhausting! A ton of very advanced concepts. Although, I appreciated how this one wasn't nearly as complicated as something like Day 19. Still, I didn't feel like spending all the time to figure this out while on vacation, so I took ideas from the solution Reddit.

### Time

- Part 1: 19:37
- Part 2: 1:12:11

## Description

### Part One

There's not much to do as you slowly descend to the bottom of the ocean. The submarine computer <span title="A STRANGE GAME.">challenges you to a nice game</span> of _Dirac Dice_.

This game consists of a single [die](https://en.wikipedia.org/wiki/Dice), two [pawns](https://en.wikipedia.org/wiki/Glossary_of_board_games#piece), and a game board with a circular track containing ten spaces marked `1` through `10` clockwise. Each player's _starting space_ is chosen randomly (your puzzle input). Player 1 goes first.

Players take turns moving. On each player's turn, the player rolls the die _three times_ and adds up the results. Then, the player moves their pawn that many times _forward_ around the track (that is, moving clockwise on spaces in order of increasing value, wrapping back around to `1` after `10`). So, if a player is on space `7` and they roll `2`, `2`, and `1`, they would move forward 5 times, to spaces `8`, `9`, `10`, `1`, and finally stopping on `2`.

After each player moves, they increase their _score_ by the value of the space their pawn stopped on. Players' scores start at `0`. So, if the first player starts on space `7` and rolls a total of `5`, they would stop on space `2` and add `2` to their score (for a total score of `2`). The game immediately ends as a win for any player whose score reaches _at least `1000`_.

Since the first game is a practice game, the submarine opens a compartment labeled _deterministic dice_ and a 100-sided die falls out. This die always rolls `1` first, then `2`, then `3`, and so on up to `100`, after which it starts over at `1` again. Play using this die.

For example, given these starting positions:

    Player 1 starting position: 4
    Player 2 starting position: 8
    

This is how the game would go:

*   Player 1 rolls `1`+`2`+`3` and moves to space `10` for a total score of `10`.
*   Player 2 rolls `4`+`5`+`6` and moves to space `3` for a total score of `3`.
*   Player 1 rolls `7`+`8`+`9` and moves to space `4` for a total score of `14`.
*   Player 2 rolls `10`+`11`+`12` and moves to space `6` for a total score of `9`.
*   Player 1 rolls `13`+`14`+`15` and moves to space `6` for a total score of `20`.
*   Player 2 rolls `16`+`17`+`18` and moves to space `7` for a total score of `16`.
*   Player 1 rolls `19`+`20`+`21` and moves to space `6` for a total score of `26`.
*   Player 2 rolls `22`+`23`+`24` and moves to space `6` for a total score of `22`.

...after many turns...

*   Player 2 rolls `82`+`83`+`84` and moves to space `6` for a total score of `742`.
*   Player 1 rolls `85`+`86`+`87` and moves to space `4` for a total score of `990`.
*   Player 2 rolls `88`+`89`+`90` and moves to space `3` for a total score of `745`.
*   Player 1 rolls `91`+`92`+`93` and moves to space `10` for a final score, `1000`.

Since player 1 has at least `1000` points, player 1 wins and the game ends. At this point, the losing player had `745` points and the die had been rolled a total of `993` times; `745 * 993 = 739785`.

Play a practice game using the deterministic 100-sided die. The moment either player wins, _what do you get if you multiply the score of the losing player by the number of times the die was rolled during the game?_

### Part Two

Now that you're warmed up, it's time to play the real game.

A second compartment opens, this time labeled _Dirac dice_. Out of it falls a single three-sided die.

As you experiment with the die, you feel a little strange. An informational brochure in the compartment explains that this is a _quantum die_: when you roll it, the universe _splits into multiple copies_, one copy for each possible outcome of the die. In this case, rolling the die always splits the universe into _three copies_: one where the outcome of the roll was `1`, one where it was `2`, and one where it was `3`.

The game is played the same as before, although to prevent things from getting too far out of hand, the game now ends when either player's score reaches at least _`21`_.

Using the same starting positions as in the example above, player 1 wins in _`444356092776315`_ universes, while player 2 merely wins in `341960390180808` universes.

Using your given starting positions, determine every possible outcome. _Find the player that wins in more universes; in how many universes does that player win?_
