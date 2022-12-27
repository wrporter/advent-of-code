package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"sort"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 21
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	food := parse(input)

	_, allergenicIngredients := getAllergenicIngredients(food)

	numSafeIngredients := 0
	for _, item := range food {
		for ingredient := range item.ingredients {
			if !allergenicIngredients[ingredient] {
				numSafeIngredients++
			}
		}
	}

	return numSafeIngredients
}

func part2(input []string) interface{} {
	food := parse(input)

	allergensToIngredients, badIngredients := getAllergenicIngredients(food)
	allergenIngredients := make(map[string]string)

	for len(badIngredients) > 0 {
		for allergen, ingredients := range allergensToIngredients {
			if len(ingredients) != 1 {
				continue
			}

			ingredient := first(ingredients)
			allergenIngredients[allergen] = ingredient

			for _, ingredients2 := range allergensToIngredients {
				delete(ingredients2, ingredient)
			}
			delete(badIngredients, ingredient)
		}
	}

	allergens := keys(allergenIngredients)
	sort.Strings(allergens)

	dangerous := make([]string, len(allergens))
	for i, allergen := range allergens {
		dangerous[i] = allergenIngredients[allergen]
	}

	return strings.Join(dangerous, ",")
}

func keys(set map[string]string) []string {
	var result []string
	for key := range set {
		result = append(result, key)
	}
	return result
}

func first(set map[string]bool) string {
	for key := range set {
		return key
	}
	return ""
}

func getAllergenicIngredients(food []Food) (map[string]map[string]bool, map[string]bool) {
	allergensToIngredients := make(map[string]map[string]bool)
	for _, item := range food {
		for allergen := range item.allergens {
			if _, ok := allergensToIngredients[allergen]; !ok {
				allergensToIngredients[allergen] = item.ingredients
			} else {
				allergensToIngredients[allergen] = intersect(allergensToIngredients[allergen], item.ingredients)
			}
		}
	}

	allergenicIngredients := make(map[string]bool)
	for _, ingredientSet := range allergensToIngredients {
		for ingredient := range ingredientSet {
			allergenicIngredients[ingredient] = true
		}
	}

	return allergensToIngredients, allergenicIngredients
}

func intersect(set1, set2 map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for item := range set1 {
		if set2[item] {
			result[item] = true
		}
	}
	return result
}

type Food struct {
	ingredients map[string]bool
	allergens   map[string]bool
}

func parse(input []string) []Food {
	var food []Food

	for _, line := range input {
		parts := strings.Split(line, " (contains ")

		ingredients := make(map[string]bool)
		for _, ingredient := range strings.Split(parts[0], " ") {
			ingredients[ingredient] = true
		}

		allergens := make(map[string]bool)
		for _, allergen := range strings.Split(strings.Trim(parts[1], ")"), ", ") {
			allergens[allergen] = true
		}

		food = append(food, Food{
			ingredients: ingredients,
			allergens:   allergens,
		})
	}

	return food
}
