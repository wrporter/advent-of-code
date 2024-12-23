# Day 21

## Thoughts

I think all I can do is shake my head. But these kinds of problems come up about once per year. Part 1 you actually solve the problem. Part 2, it's not really expected that you solve the problem with code.

### Part 1

1. Parse all the monkeys and their respective jobs into a map.
2. Recursively solve the mathematical expression for `root`.
   1. If the monkey just calls out a number, return it.
   2. Otherwise, recurse down the left and right sides and perform the operation on the results of the two.

### Part 2

1. Parse all the monkeys as before.
2. Replace the `root` equation with an `=`.
3. Expand the equation with recursive logic similar to above and add it all together in a large string.
4. Use an algebra solver online, a library, or by hand. It's not practical to solve it automatically for the puzzle. And I believe that's expected. Just for fun so you know what I mean, here's my equation where you need to solve for `humn`. Pretty long!

```
((3) * ((((5) + (2)) * (((((((4) * (((3) * ((2) + (5))) / (3))) + ((3) * (5))) * (((2) * ((((7) * (((((3) + ((5) * (2))) * (3)) - (1)) + ((12) + (17)))) + ((((((8) + ((2) * (11))) + (1)) * ((7) * (2))) + (((9) + (2)) * (3))) * (3))) + ((((((((((9) * (2)) * ((9) + (15))) / (8)) + ((2) * ((6) + (1)))) / (4)) * ((((((4) * (2)) * (((1) + (8)) + ((2) * ((2) + (9))))) + (((2) * (19)) + (((2) * (20)) + ((((3) * ((5) + (2))) + (8)) - (((1) + (8)) - (1)))))) + ((6) * ((4) + (10)))) * (2))) + (((9) * (10)) * ((17) + ((((((11) + ((18) + (5))) + (13)) * (4)) / (4)) * (2))))) + ((((((((((5) + (14)) - (2)) * (2)) - (11)) + ((2) * (((4) * (2)) + ((3) + (4))))) * (3)) + ((((((5) + (2)) * (3)) + (10)) * (3)) * (3))) + ((2) * (((14) * (5)) - (19)))) + (((((2) * (((3) * ((7) + ((1) + (9)))) + ((18) + (4)))) / (2)) * ((((15) + ((3) * (3))) + (((((2) * (3)) * (2)) * (3)) + (((3) + (((2) + (5)) + (((5) * (11)) - (18)))) * (2)))) + (3))) - ((((10) * (2)) / (2)) * (((15) * (4)) + (1)))))) + ((((((2) * (5)) * (((11) + (6)) + ((2) * (13)))) + ((2) * (((((2) * (4)) * (2)) + (1)) + (((4) * (2)) * ((3) * (4)))))) + (((((2) + (5)) * (4)) + (9)) + (((5) + ((4) * (3))) * (5)))) * (((4) + (4)) + ((((((((3) * (3)) * ((((2) * (5)) - (3)) + (4))) - ((16) * (2))) + (((17) * (2)) + ((5) + ((8) + ((5) + (((4) * (3)) + ((3) * (3)))))))) + (4)) / (2)) / ((2) * (4)))))))) / (2))) + ((((((((((4) * (2)) + ((((5) * (3)) + (16)) * (3))) * (4)) + ((3) * ((((18) + ((4) * (2))) + ((11) * (5))) + ((14) * (2))))) + ((((5) * ((3) * (5))) + (((10) + (1)) * (2))) * (4))) * ((((3) * (16)) / (2)) * ((3) * (9)))) / (2)) + ((((2) * ((2) * (3))) * ((((((((4) * ((4) + ((((2) * ((11) + ((3) + ((10) + (13))))) * (2)) / (4)))) / (2)) * ((2) * (4))) + ((11) * ((13) * (7)))) + ((((13) * (6)) - ((11) - (1))) + ((((2) * (((1) + ((4) * (3))) * (3))) + ((((5) * (3)) + (7)) / (2))) - (13)))) * (2)) + (((4) * (((3) * (3)) + (2))) + ((3) + (((14) * (2)) - (8)))))) + ((((9) * (3)) + ((((6) * (4)) + (5)) * ((2) * ((6) + (1))))) * (11)))) * (12))) * (((((12) + ((4) + (((2) * (4)) + (((3) * (4)) + (1))))) * ((((((5) + ((5) * (2))) * (3)) + (20)) + ((2) * ((((((((3) * (4)) * (3)) * (20)) / (3)) + (1)) * (2)) / (2)))) * (2))) + (((((2) * (15)) * (12)) + (((2) * (11)) + ((5) * (3)))) * ((((8) + (5)) * (17)) - (((2) * (4)) * (3))))) * (4))) + ((((1) + (((2) * (4)) * (5))) * (((2) * (((2) * (((2) * (((((((4) * (2)) * (5)) - (9)) * ((((9) + ((6) * (2))) * (((((((((3) + (4)) + (4)) * (2)) + (1)) * ((5) + (2))) + ((5) * (5))) / (3)) / (2))) / ((3) + (4)))) * ((((17) * (19)) + (((4) + (3)) * (((5) + ((4) * (2))) + ((11) * (3))))) + (((3) * (((3) * (2)) + (5))) + ((((9) * ((2) * (17))) + (((((5) + ((((2) * ((8) + (((3) * (((2) * ((10) + ((5) + (2)))) + ((11) + ((2) * (((9) * (2)) - (5)))))) / (3)))) / (2)) * ((3) * (2)))) * (2)) - ((6) + (((5) * ((4) * (3))) + (13)))) / (3))) * (3))))) - (((7) * (((15) * (((4) * (4)) + (((5) + (5)) - (3)))) - ((4) * ((4) * (4))))) * (((((((5) * (6)) + (1)) * (2)) + (((((((3) * (((((7) + ((6) * (6))) + ((4) * ((((13) * (3)) - ((9) + (1))) + (2)))) - (12)) / (5))) * (2)) / (2)) * (2)) / (6)) * (2))) + ((2) + (((2) * (((((15) + (16)) * (2)) / (2)) * (2))) / (4)))) * (2))))) / (2))) / (2))) / (2))) * (((((7) * ((16) + (15))) * ((2) * (((3) * (2)) + (1)))) + ((((5) * (((17) * (3)) - (8))) * (3)) + (((((((((((5) * (3)) + (4)) * (2)) / (2)) + ((3) + (4))) + (3)) * (3)) * ((2) * (3))) + ((19) * (5))) * (11)))) - ((((7) * (7)) - (6)) * ((((2) * (4)) * ((4) * (2))) + (3))))))) - ((((((((2) * (17)) + (3)) + (((4) * (2)) * (8))) * (7)) - ((20) * (9))) + ((((((((((3) * ((((2) * (((3) * (9)) + ((((20) + (7)) + ((7) + (3))) * (2)))) + (((11) * (13)) - (9))) + ((((((((4) * (((11) * ((15) + (4))) - ((2) + (10)))) / (2)) * (2)) + (((5) * ((11) * (2))) + (((((13) * ((5) * ((3) * (3)))) + (((((((((((((((((((((5) * (10)) + ((((13) * (5)) + ((2) * ((3) * ((3) + (8))))) * (2))) - ((((2) * (((2) * ((((2) * (15)) + (16)) + (13))) + ((3) * ((6) + (5))))) / (2)) - ((((((2) + (((5) + (8)) * (2))) * (2)) + (3)) - ((4) * (4))) + (7)))) * (2)) + ((2) * ((((((((9) + (17)) - (7)) + (4)) * (2)) + (((((7) * (3)) + (((2) * (14)) + (4))) * (3)) * (2))) / (2)) + ((((((2) * (((16) + ((3) + (10))) + ((4) * (3)))) / (2)) * (2)) / (2)) + ((((2) * (13)) / (2)) * (2)))))) + (((((4) * (((((2) * (((((((4) * (5)) * (((((5) * (2)) + (3)) * (((14) + (3)) - (4))) + (((((2) * (4)) * (((((8) + (9)) * (11)) + (((17) * (11)) + (((3) * (4)) + (((((3) * (((((((((2) * (3)) * (4)) - (1)) + ((10) + ((2) * (3)))) / (3)) + (6)) * (2)) / (2))) - ((3) * ((5) + (1)))) + (7)) / (2))))) + ((((((2) * ((((((((2) * ((((3) + ((2) * (19))) * (3)) + ((((((6) * ((2) + (4))) + ((2) * (3))) - (7)) + (2)) + (((5) + (8)) * (3))))) * (2)) + (((13) * (4)) * (((((5) * (5)) * (2)) + ((3) * (((2) * (5)) + ((3) * (3))))) + (((humn) - ((((7) + (((3) * (6)) * (3))) + ((12) / (2))) + (((12) * ((2) * (4))) - ((5) * (5))))) / (2))))) / (2)) - ((((2) * (16)) + (2)) * ((5) + (2)))) * (2)) + ((10) * (6)))) - (((((10) + (((4) + (3)) * (3))) * (13)) - ((5) * (((9) * (2)) + (4)))) * (2))) / (10)) - ((((2) * (5)) * ((2) * ((12) + ((10) + (1))))) + (((2) * (5)) + ((3) * (3))))) / (4)))) - (((2) * ((((2) * (((3) * (((3) * (2)) + (5))) / (3))) * ((5) * (5))) + ((1) + (20)))) - ((11) * (((3) * (3)) * (2))))) / (8)))) + (((4) * (4)) + ((5) * (5)))) + (((((2) * (((5) + (2)) + ((1) + (5)))) + (((17) + ((3) * (2))) + ((14) - (4)))) * (2)) + (((7) * (3)) * ((13) + (2))))) / (2)) - ((((((17) + ((8) * (3))) + ((9) + (10))) + (((2) + (6)) + (((4) + ((7) + (2))) * (3)))) * (3)) * (3)))) + ((((3) * (((2) * ((4) * (3))) / (2))) + (((4) * ((((1) + (6)) + (15)) + (6))) + ((2) + (9)))) * (3))) / (5)) - ((17) + ((2) * (((12) + ((2) + ((10) + (1)))) + ((3) * (18))))))) + (((5) + ((((3) * (4)) + (1)) * (2))) * ((2) + ((6) * (5))))) * (2)) - (((3) * (7)) * (((5) + (3)) + (15))))) / (3)) + ((3) * ((((3) * ((3) * (((4) * (4)) + (1)))) + (15)) + ((3) * (((2) * (((2) * (3)) * (2))) + (17)))))) * (2)) - (((((13) * ((20) + ((4) * (7)))) - ((((2) * (((((3) * (13)) + (2)) * (2)) / (2))) / (2)) * (3))) + ((((2) + (11)) * (13)) * (2))) + (((1) + ((8) + (((5) + (2)) * (3)))) + (13)))) / (6)) + (((3) * ((((5) + (18)) * (5)) + ((2) * (((19) + ((3) * (2))) + (4))))) - (((2) * (((2) * ((12) + (((8) - (1)) * ((10) - (3))))) - (1))) / (2)))) * (3)) - ((2) * (((3) * ((5) * (5))) + (((3) + ((((2) * ((((4) + (9)) * (3)) - ((3) * (2)))) + (4)) + (((20) + (2)) / (2)))) - (10))))) + (4)) / (3)) + ((((((((((5) * ((1) + (18))) - ((5) + (17))) + ((3) * (((1) + (6)) + (4)))) / (2)) + (((5) * (3)) * (4))) * (2)) * (3)) + ((4) * ((((((2) * ((((5) + (14)) * (2)) + (3))) / (2)) + (((((4) * (7)) + ((((11) * (3)) + (8)) - (4))) / (5)) * (2))) + ((15) + (((5) + ((3) * (2))) + (4)))) * (2)))) / (2))) * (2)) - ((6) * (((((5) * (7)) + ((13) * (2))) * (2)) / (2)))) / (3))) * (10)) - (((((18) / (2)) + ((4) * (2))) * ((4) + ((8) - (1)))) + (((3) * (8)) * (14)))))) / (3)) - (((7) * ((2) * (((2) * (13)) + ((((((4) + ((7) * (17))) / (3)) * ((6) * (4))) / (4)) / (6))))) - (((3) * (((16) * (2)) - (3))) * (3)))) / (2)))) + ((11) * (11))) / (2)) - (((2) + (((7) + ((5) * (2))) + (12))) + (((3) * (((2) * (3)) + (((((3) * (2)) * (6)) / (2)) + ((7) + ((((((((2) * (14)) - (9)) + (11)) - (7)) * (2)) / (2)) - (7)))))) + ((5) * (((2) * (3)) + ((((2) * ((6) + (8))) / (2)) / (2))))))) * (3)) - (((2) * (((14) + ((5) + ((2) * ((3) * (2))))) - (8))) * (9))) * (2)) + (((2) * ((5) + (4))) * ((14) * (3)))) / (3))) / (3)))) = ((((((((2) + ((3) * (5))) * (2)) * (2)) * ((((((((((1) + ((5) * (5))) * (18)) + (((((6) + (5)) * ((((14) + (4)) + ((11) + ((16) * (18)))) * (2))) + (((((2) * (4)) * (4)) * (4)) * ((((2) * (8)) + (1)) + (2)))) / (2))) + ((((((9) * (3)) + ((((((7) + ((1) + ((7) * (3)))) * (2)) + (((2) + (((11) * (2)) / (2))) + (((12) * (3)) / (2)))) - (((13) - (2)) + (12))) / (2))) - (19)) - (12)) * (((5) * (5)) - ((3) * (2))))) / (2)) * ((((((3) * (3)) * (2)) + (5)) * (2)) / (2))) + (((2) * ((2) * (((11) * (((7) * (4)) + ((((3) * ((((2) * (4)) + ((7) * (2))) / (2))) - (10)) + (((3) * (4)) - (4))))) + ((2) * (((2) * (((5) * (((3) + (5)) + (3))) + ((2) * ((((((5) + (2)) * (2)) * (3)) + (4)) + ((5) * (((4) + (3)) + (4))))))) / (2)))))) * (17))) + (((((1) + (((((4) * ((3) * (2))) + (5)) + (6)) + ((13) * (2)))) / (2)) * (((4) * ((3) * (3))) + (((2) * ((10) + (((4) * (2)) - (1)))) + ((5) * ((3) + ((2) * (4))))))) * ((((3) + (6)) + ((4) + ((6) + (1)))) * (((6) + ((((((5) * ((2) * (3))) + (((6) + (1)) * (11))) + (2)) + ((19) * (3))) / (2))) * (4))))) * (4))) + (((((5) * (18)) + (11)) * (((((4) + ((2) * (7))) + (20)) * (10)) + ((((2) * (3)) + (17)) * (5)))) * ((((((5) * (((((2) + ((3) * (3))) * (((((9) + (((4) * (4)) + ((((10) - (2)) * (3)) - (5)))) + (12)) + ((5) * (2))) + (((4) + (6)) + ((2) + (5))))) + ((((((13) * (3)) * ((2) * (4))) + ((8) * (3))) + (1)) + (((4) * ((18) + (5))) + (((((3) * (7)) * (2)) + ((((8) + (1)) + (((4) + (6)) - (3))) + (((7) * (5)) - (4)))) * (3))))) + ((((10) * ((1) + (8))) + ((((11) * (4)) / (4)) + ((4) * (((4) + (2)) * (2))))) * (6)))) - ((((((((7) * (2)) * (2)) * (4)) + ((((2) * (5)) + (((11) + (2)) + (6))) * (9))) + ((((17) * (7)) + (20)) * (2))) + (((8) * ((8) * (2))) * (16))) - ((((2) * (((((((2) * ((((11) + (6)) * (2)) / (2))) / (2)) * (3)) + (((5) * (3)) + ((6) + (1)))) + ((2) * (3))) * (2))) + (((2) * (11)) * ((13) + ((3) * (2))))) + ((((9) + (8)) + ((6) * ((4) + (3)))) + (4))))) + ((((((((((4) * ((3) * ((((3) * ((3) * (3))) + ((2) * ((3) + (14)))) + (6)))) + (((2) * (8)) + (1))) * (2)) / (2)) * ((3) * (3))) * (3)) - (((4) * ((2) * (11))) * (((18) * (3)) + (13)))) * (2)) / (2))) + ((((((2) * ((((((3) * (4)) + ((5) + (2))) * (2)) / (2)) * ((((7) + ((12) * ((2) + (4)))) * (3)) / (3)))) + (((2) * ((((10) + ((5) + (2))) + (9)) + (((((((8) + (5)) + (18)) * (2)) / (2)) + ((3) * (14))) + (10)))) + (((19) * ((7) * (3))) + (((13) * (((3) * ((4) + (3))) + ((4) * (3)))) + ((((((((18) + ((16) - (5))) * (4)) + ((8) * (4))) + ((((2) * ((5) * (5))) + ((12) + ((7) + (10)))) + (((((2) * (((((3) + (10)) * (4)) + ((3) * (14))) / (2))) - ((3) * (((6) + (4)) - (3)))) * (3)) / (3)))) + (((1) + (10)) * ((1) + (6)))) + ((4) + (19))) + ((((((5) * (10)) + ((5) + (4))) + ((3) * (5))) + (3)) * ((1) + (6)))))))) * (5)) + ((((((4) + ((4) + (3))) * ((2) * (17))) - ((((5) + (4)) * (5)) + ((2) * ((2) * (10))))) * ((5) * (5))) + ((((((2) * ((((((2) * ((((7) * (2)) + (8)) / (2))) - (5)) + ((11) + (9))) - ((3) + (5))) + (((2) * (4)) * (4)))) + ((((7) * ((7) * (2))) + (4)) + (((10) + (1)) + ((((2) * (((4) + ((1) + (6))) * (2))) + (((5) + (4)) * (3))) * (2))))) + (((((9) * (((6) + ((7) + (4))) + (2))) * (3)) - ((((6) * (3)) + ((3) + ((2) * ((3) + ((((3) + (8)) - (1)) * (5)))))) + (9))) + ((11) * (5)))) * (3)) + (((((4) * ((2) * (17))) / (8)) * ((14) / (2))) * (7))))) + (((((((((4) + ((4) * (14))) / (5)) - (1)) * (5)) + (((((2) * (4)) * (2)) + ((5) * (5))) + (1))) + ((((7) * ((4) + (3))) * (5)) + ((2) * ((2) * (((4) * ((5) * (2))) + (((4) + (3)) + (((4) * (7)) + (8)))))))) / (2)) * (2)))) + (((3) * (5)) * ((((2) * ((((((((3) * (6)) + (13)) - (((3) * (3)) - (2))) + (2)) * (2)) - (15)) + ((3) * ((2) * (17))))) / (2)) + ((6) * (4))))))) * (2)) * ((((2) + ((3) + ((7) * (3)))) / (2)) * (((((((3) * (((12) * (2)) + (7))) / (3)) * (2)) + ((9) - (3))) - (1)) * (3))))
```

## Time

This is how long it took me to complete each part.

- Part 1: 15:48
- Part 2: 47:08

## Puzzle

# Day 21: Monkey Math

[https://adventofcode.com/2022/day/21](https://adventofcode.com/2022/day/21)

## Description

### Part One

The [monkeys](https://adventofcode.com/2022/day/11) are back! You're worried they're going to try to steal your stuff again, but it seems like they're just holding their ground and making various monkey noises at you.

Eventually, one of the elephants realizes you don't speak monkey and comes over to interpret. As it turns out, they overheard you talking about trying to find the grove; they can show you a shortcut if you answer their _riddle_.

Each monkey is given a _job_: either to _yell a specific number_ or to _yell the result of a math operation_. All of the number-yelling monkeys know their number from the start; however, the math operation monkeys need to wait for two other monkeys to yell a number, and those two other monkeys might _also_ be waiting on other monkeys.

Your job is to _work out the number the monkey named `root` will yell_ before the monkeys figure it out themselves.

For example:

    root: pppw + sjmn
    dbpl: 5
    cczh: sllz + lgvd
    zczc: 2
    ptdq: humn - dvpt
    dvpt: 3
    lfqf: 4
    humn: 5
    ljgn: 2
    sjmn: drzm * dbpl
    sllz: 4
    pppw: cczh / lfqf
    lgvd: ljgn * ptdq
    drzm: hmdt - zczc
    hmdt: 32


Each line contains the name of a monkey, a colon, and then the job of that monkey:

*   A lone number means the monkey's job is simply to yell that number.
*   A job like `aaaa + bbbb` means the monkey waits for monkeys `aaaa` and `bbbb` to yell each of their numbers; the monkey then yells the sum of those two numbers.
*   `aaaa - bbbb` means the monkey yells `aaaa`'s number minus `bbbb`'s number.
*   Job `aaaa * bbbb` will yell `aaaa`'s number multiplied by `bbbb`'s number.
*   Job `aaaa / bbbb` will yell `aaaa`'s number divided by `bbbb`'s number.

So, in the above example, monkey `drzm` has to wait for monkeys `hmdt` and `zczc` to yell their numbers. Fortunately, both `hmdt` and `zczc` have jobs that involve simply yelling a single number, so they do this immediately: `32` and `2`. Monkey `drzm` can then yell its number by finding `32` minus `2`: _`30`_.

Then, monkey `sjmn` has one of its numbers (`30`, from monkey `drzm`), and already has its other number, `5`, from `dbpl`. This allows it to yell its own number by finding `30` multiplied by `5`: _`150`_.

This process continues until `root` yells a number: _`152`_.

However, your actual situation involves <span title="Advent of Code 2022: Now With Considerably More Monkeys">considerably more monkeys</span>. _What number will the monkey named `root` yell?_

### Part Two

Due to some kind of monkey-elephant-human mistranslation, you seem to have misunderstood a few key details about the riddle.

First, you got the wrong job for the monkey named `root`; specifically, you got the wrong math operation. The correct operation for monkey `root` should be `=`, which means that it still listens for two numbers (from the same two monkeys as before), but now checks that the two numbers _match_.

Second, you got the wrong monkey for the job starting with `humn:`. It isn't a monkey - it's _you_. Actually, you got the job wrong, too: you need to figure out _what number you need to yell_ so that `root`'s equality check passes. (The number that appears after `humn:` in your input is now irrelevant.)

In the above example, the number you need to yell to pass `root`'s equality test is _`301`_. (This causes `root` to get the same number, `150`, from both of its monkeys.)

_What number do you yell to pass `root`'s equality test?_

