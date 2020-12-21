package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type direction int

type food struct {
	ingredients []string
	allergens   []string
}

var foodRegex *regexp.Regexp

func main() {
	// file, err := os.Open("../test21.txt")
	file, err := os.Open("../input21.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	foodRegex = regexp.MustCompile(`(.*) \(contains (.*)\)`)

	foods := make([]food, 0)
	allergenMap := make(map[string][]string)
	ingredientToAllergenMap := make(map[string]string)
	for sc.Scan() {
		line := sc.Text()
		if line != "" {
			submatches := foodRegex.FindStringSubmatch(line)
			ings := strings.Split(submatches[1], " ")
			algs := strings.Split(submatches[2], ", ")
			foods = append(foods, food{ingredients: ings, allergens: algs})

			for _, alg := range algs {
				updateAllergens(alg, ings, &allergenMap)

				updateIngredientMap(alg, &allergenMap, &ingredientToAllergenMap)
			}
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	for {
		definite := true
		for alg, potentialMatches := range allergenMap {
			if len(potentialMatches) > 1 {
				definite = false
			}
			updateIngredientMap(alg, &allergenMap, &ingredientToAllergenMap)
		}
		if definite {
			break
		}
	}

	log.Println(foods)
	log.Println(allergenMap)
	log.Println(ingredientToAllergenMap)

	var res int
	for _, fo := range foods {
		for _, ing := range fo.ingredients {
			if _, ok := ingredientToAllergenMap[ing]; !ok {
				res++
			}
		}
	}

	log.Println(res)
}

func updateIngredientMap(allergen string, allergenMap *map[string][]string, ingredientToAllergenMap *map[string]string) {
	potentialMatches := (*allergenMap)[allergen]

	if len(potentialMatches) == 1 {
		match := potentialMatches[0]
		(*ingredientToAllergenMap)[match] = allergen
		for alg1, ings1 := range *allergenMap {
			if allergen == alg1 {
				continue
			}
			idx := -1
			for i, ing1 := range ings1 {
				if ing1 == match {
					idx = i
					break
				}
			}
			if idx >= 0 {
				(*allergenMap)[alg1] = remove(ings1, idx)
			}
		}
	}
}

func updateAllergens(allergen string, ingredients []string, allergenMap *map[string][]string) {
	if potentialIngs, ok := (*allergenMap)[allergen]; ok {
		(*allergenMap)[allergen] = intersection(potentialIngs, ingredients)
	} else {
		cop := make([]string, len(ingredients))
		copy(cop, ingredients)
		(*allergenMap)[allergen] = cop
	}
}

func reverse(value string) string {
	// Convert string to rune slice.
	// ... This method works on the level of runes, not bytes.
	data := []rune(value)
	result := []rune{}

	// Add runes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	// Return new string.
	return string(result)
}

func unique(slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
