package main

import (
	"aoc/src/cli/cmd"
	"aoc/src/cli/cmd/prettylog"
	"log/slog"
)

func main() {
	logger := slog.New(prettylog.NewHandler(nil))
	slog.SetDefault(logger)

	cmd.Execute()
}
