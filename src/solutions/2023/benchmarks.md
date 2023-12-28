Benchmarks taken on:

```
Chip: Apple M1 Max
Cores: 8
Memory: 32 GB
CPU: 3.2 GHz
```

Helpful tools:

```shell
GOPROXY= GOPRIVATE= GONOPROXY= go install golang.org/x
/perf/cmd/benchstat 
GOPROXY= GOPRIVATE= GONOPROXY= go install golang.org/x/tools/cmd/benchcmp@latest
```

Run benchmarks for all solutions:

```shell
go test -benchmem -bench=. ./src/solutions/2023/...
```

Compare benchmarks:

```shell
go test ./src/solutions/2023/23/go -bench=BenchmarkSolution_Part2 -benchmem -benchtime=5s | tee old.txt
go test ./src/solutions/2023/23/go -bench=BenchmarkSolution_Part2 -benchmem -benchtime=5s | tee new.txt
benchcmp old.txt new.txt
```

| Day    | Part 1    | Part 2    |
|--------|-----------|-----------|
| Day 1  | `453μs`   | `584μs`   |
| Day 2  | `240μs`   | `263μs`   |
| Day 3  | `57μs`    | `5.14ms`  |
| Day 4  | `564μs`   | `576μs`   |
| Day 5  | `43μs`    | `145μs`   |
| Day 6  | `541ns`   | `299ns`   |
| Day 7  | `648μs`   | `620μs`   |
| Day 8  | `804μs`   | `1.08ms`  |
| Day 9  | `199μs`   | `200μs`   |
| Day 10 | `630μs`   | `2.24ms`  |
| Day 11 | `128μs`   | `127μs`   |
| Day 12 | `555μs`   | `4.97ms`  |
| Day 13 | `183μs`   | `178μs`   |
| Day 14 | `33μs`    | `34.9ms`  |
| Day 15 | `90μs`    | `393μs`   |
| Day 16 | `1.63ms`  | `451ms`   |
| Day 17 | `49.79ms` | `80.70ms` |
| Day 18 | `49μs`    | `64μs`    |
