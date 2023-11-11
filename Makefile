dist:
	go build cmd/ao/ao.go

install:
	go install cmd/ao/ao.go

default:
	ao $(ARGS)

start:
	ao start

part1:
	ao end --part 1
part2:
	ao end --part 2

install-aoc:
	brew install scarvalhojr/tap/aoc-cli

