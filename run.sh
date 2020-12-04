#!/usr/bin/env bash

year=${1:-$(date +"%Y")}
: "${year:?Must pass the year}"

echo "--- Executing files for year: ${year}"
echo

find "${year}" -name "main.go" -maxdepth 2 -print0 | sort --version-sort -z | xargs -0 -I{} sh -c "go run {}; echo"
