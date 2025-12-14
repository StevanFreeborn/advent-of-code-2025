# Advent of Code 2025

This repository contains my solutions for the [Advent of Code 2025](https://adventofcode.com/2025) challenges.

## Languages

I am using Go for this year's challenge again.

## Structure

Each day's solution is contained in a separate package.

### Prerequisite

- [Go](https://go.dev/doc/install)

### Running the Solutions

To run a specific day's solution you can run the packages tests:

```pwsh
go test ./cmd/01
```

## Challenges

| Day | Problem                                         | Solution              | Notes                                                                                                                                                                                                                 |
|-----|-------------------------------------------------|-----------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 01  | [Problem](https://adventofcode.com/2025/day/1)  | [Solution](./cmd/01/) | Part 1 seemed straight forward enough. Part 2 completely stumped me. The trickiest part was not counting zero's twice when you ended and then started from there.                                                     |
| 02  | [Problem](https://adventofcode.com/2025/day/2)  | [Solution](./cmd/02/) | Part 1 should have been super straight-forward, but some whitespace screwed up my parsing of the final range. I some how was able to stumble over a solution to part 2 rather blindly.                                |
| 03  | [Problem](https://adventofcode.com/2025/day/3)  | [Solution](./cmd/03/) | Part 1 intuitively makes you think you only care about highest nums, but then you need to consider placement. I once again tripped over the solution using a stack. However chat morally supported me the whole time. |
| 04  | [Problem](https://adventofcode.com/2025/day/4)  | [Solution](./cmd/04/) | Part 1 was by far the most straight-forward challenge so far. This was easiest day of AoC this year by far.                                                                                                           |
| 05  | [Problem](https://adventofcode.com/2025/day/5)  | [Solution](./cmd/05/) | Part 1 was fairly straight forward. Merging ranges was trickiest bit. I feel like I'm being led into a trap with how straight forward today was.                                                                      |
| 06  | [Problem](https://adventofcode.com/2025/day/6)  | [Solution](./cmd/06/) | Part 1 I don't think it is the best solution, but it does solve it. I kicked ass on part 2. 🥳                                                                                                                        |
| 07  | [Problem](https://adventofcode.com/2025/day/7)  | [Solution](./cmd/07/) | FIFO for life. 🤣. Fuck me...part 2 I found really difficult and the use of DFS was obvious, but doing the memoizing was not.                                                                                         |
| 08  | [Problem](https://adventofcode.com/2025/day/8)  | [Solution](./cmd/08/) |                                                                                                                                                                                                                       |
| 09  | [Problem](https://adventofcode.com/2025/day/9)  | [Solution](./cmd/09/) |                                                                                                                                                                                                                       |
| 10  | [Problem](https://adventofcode.com/2025/day/10) | [Solution](./cmd/10/) |                                                                                                                                                                                                                       |
| 11  | [Problem](https://adventofcode.com/2025/day/11) | [Solution](./cmd/11/) |                                                                                                                                                                                                                       |
| 12  | [Problem](https://adventofcode.com/2025/day/12) | [Solution](./cmd/12/) |                                                                                                                                                                                                                       |
