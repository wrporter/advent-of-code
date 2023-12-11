# Advent of Code

This is a repository for all my Advent of Code solutions over the years. I've done them all in Go, but in 2022 I wanted to try things out in TypeScript since that's what I was using the most at that time in my work. There was one particular day on which my code was too slow. I converted it to Go, and it was **10 times faster**! I decided I probably will stick with Go as my primary language in the future.

## CLI

I wrote a small CLI wrapper around some other utilities out there to automate things I do on each day.

### Usage

1. Install [aoc-cli](https://github.com/scarvalhojr/aoc-cli).
    - After you log in at https://adventofcode.com/, copy your `session` cookie and save it to `.adventofcode.session`.
2. Compile and install the binary with `go install cmd/ao/ao.go`.
3. For convenience, create a `.ao.yaml` file. Example:
    ```yaml
    year: 2022
    day: 13
    language: go
    leaderboardId: 11111
    memberId: 11111
    ```
4. Copy the directory [templates/0-template](templates/0-template). This is the base directory for the Go templates to generate files for a new day.
5. Start a day with `ao start -d 13`.
    - Writes new files from the templates for the day.
    - Starts a timer.
    - Downloads puzzle input from the Advent of Code site. Formats the file to remove any trailing newline characters to keep files consistent.
    - When a language is specified, the timer is scoped to that language's directory. This can be a nice way to time how long it takes you to write solutions in each language.
6. End a part with `ao end -d 13 -p 1`.
    - Records how long it took you to solve the part. Prefers the timestamp from a private leaderboard (both the `leaderboardId` and the `memberId` must be specified in `.ao.yaml`). Otherwise, uses the current time to determine how long the part took to solve.
    - Upon completing Part 2, downloads the full puzzle description.
7. Run the solution for a day with `ao run`.
    - Specify the year or day with the `--year` and `--day` flags.
    - Run all days for a year with `--day 0`.

```shell
$ ao help
A wrapper around the aoc CLI https://github.com/scarvalhojr/aoc-cli

Usage:
  ao [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  end         Record end time for a part. Download puzzle description after completing part 2.
  help        Help about any command
  run         Run solutions for a year and day (runs all days when day is 0)
  start       Start solving parts for a given day. Generates a solution template, downloads the puzzle input, and writes a starting time.

Flags:
  -c, --config string        config file (default is $HOME/.ao.yaml)
  -d, --day int              event day (default 11)
  -h, --help                 help for ao
  -l, --language string      specify writing templated files for a single language (default "all")
  -o, --output-path string   path to output files to (default is solutions/{year}/{day})
  -p, --part int             part to record time for (default 1)
  -y, --year int             event year (default 2023)

Use "ao [command] --help" for more information about a command.
```
