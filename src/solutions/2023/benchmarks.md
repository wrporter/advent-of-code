Benchmarks taken on an M1 Mac.

```shell
go test -benchmem -bench=. ./solutions/2023/...
```

```shell
goos: darwin
goarch: arm64
```

| Day    | Part 1    | Part 2    |
|--------|-----------|-----------|
| Day 1  | `809ns`   | `1.78μs`  |
| Day 2  | `6.70μs`  | `7.47μs`  |
| Day 3  | `37ns`    | `48ns`    |
| Day 4  | `3.94μs`  | `3.96μs`  |
| Day 5  | `4.04μs`  | `5.18μs`  |
| Day 6  | `343ns`   | `244ns`   |
| Day 7  | `1.00μs`  | `1.01μs`  |
| Day 8  | `2.93μs`  | `3.67μs`  |
| Day 9  | `823ns`   | `817ns`   |
| Day 10 | `1.25μs`  | `2.07μs`  |
| Day 11 | `632ns`   | `636ns`   |
| Day 12 | `2.24μs`  | `18.42μs` |
| Day 13 | `966ns`   | `950ns`   |
| Day 14 | `633ns`   | `18.94μs` |
| Day 15 | `229ns`   | `848ns`   |
| Day 16 | `1.63ms`  | `451ms`   |
| Day 17 | `49.79ms` | `80.70ms` |
| Day 18 | `49μs`    | `64μs`    |