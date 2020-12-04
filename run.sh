#!/usr/bin/env bash

find 2020 -name "main.go" -maxdepth 2 -print0 | sort --version-sort -z | xargs -0 -I{} go run {}
