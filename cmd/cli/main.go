package main

import (
	"github.com/wrporter/advent-of-code/cmd/cli/cmd"
	"github.com/wrporter/advent-of-code/cmd/cli/cmd/prettylog"
	"log/slog"
)

func main() {
	logger := slog.New(prettylog.NewHandler(nil))
	slog.SetDefault(logger)

	cmd.Execute()
}
