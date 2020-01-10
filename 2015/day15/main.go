package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"regexp"
)

var regex = regexp.MustCompile(`([a-zA-Z]+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`)

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func getOptimalRecipe(ingredientStrings []string, teaspoons int, calorieAmount int) int {
	var ingredients []Ingredient
	for _, ingredientString := range ingredientStrings {
		match := regex.FindStringSubmatch(ingredientString)
		ingredients = append(ingredients, Ingredient{
			name:       match[1],
			capacity:   conversion.StringToInt(match[2]),
			durability: conversion.StringToInt(match[3]),
			flavor:     conversion.StringToInt(match[4]),
			texture:    conversion.StringToInt(match[5]),
			calories:   conversion.StringToInt(match[6]),
		})
	}

	max := 0
	properties := []string{"capacity", "durability", "flavor", "texture"}

	sumCombo(teaspoons, len(ingredients), func(values []int) {
		propScores := make(map[string]int)
		for i, ingredient := range ingredients {
			propScores["capacity"] += values[i] * ingredient.capacity
			propScores["durability"] += values[i] * ingredient.durability
			propScores["flavor"] += values[i] * ingredient.flavor
			propScores["texture"] += values[i] * ingredient.texture
			propScores["calories"] += values[i] * ingredient.calories
		}

		if calorieAmount > 0 && propScores["calories"] != calorieAmount {
			return
		}

		score := 1
		for _, property := range properties {
			if propScores[property] < 0 {
				score = 0
				break
			}
			score *= propScores[property]
		}
		max = ints.Max(max, score)
	})

	return max
}

func sumCombo(target int, size int, emit func(values []int)) {
	current := make([]int, size)
	sumComboRec(target, current, 0, target, size, emit)
}

func sumComboRec(leftOver int, current []int, index int, target int, size int, emit func(values []int)) {
	if leftOver < 0 {
		return
	}

	if index == size {
		if leftOver == 0 {
			emit(current)
		}
		return
	}

	for value := current[index] + 1; value < target; value++ {
		current[index] = value
		sumComboRec(leftOver-value, current, index+1, target, size, emit)
	}
	// backtrack
	current[index] = 0
}

func main() {
	lines, _ := file.ReadFile("./2015/day15/input.txt")
	//lines := []string{
	//	"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
	//	"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
	//}
	fmt.Println(getOptimalRecipe(lines, 100, -1))
	fmt.Println(getOptimalRecipe(lines, 100, 500))
}
