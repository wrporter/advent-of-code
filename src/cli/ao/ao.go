package main

import (
	"aoc/src/cli/ao/cmd"
	"aoc/src/cli/ao/cmd/prettylog"
	"log/slog"
)

func main() {
	logger := slog.New(prettylog.NewHandler(nil))
	slog.SetDefault(logger)

	cmd.Execute()
}
