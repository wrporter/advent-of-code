#!/usr/bin/env bash

year=${1:-$(date +"%Y")}
: "${year:?Must pass the year}"

green='\033[1;92m'
reset='\033[0m'

echo -e "${green}===============================${reset}"
echo -e "${green}= Saving Santa for year: ${year} =${reset}"
echo -e "${green}===============================${reset}"
echo

find "solutions/${year}" -name "main.go" -maxdepth 2 -print0 | sort --version-sort -z | xargs -0 -I{} sh -c "go run {}; echo"
