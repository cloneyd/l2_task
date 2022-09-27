package main

import (
	"fmt"
	"sort"
	"strings"
)

func checkAnagram(first string, second string) bool {
	counter := make(map[rune]int)

	runeFirst := []rune(first)
	runeSecond := []rune(second)

	if len(runeFirst) != len(runeSecond) {
		return false
	}

	for i := 0; i < len(runeFirst); i++ {
		counter[runeFirst[i]]++
		counter[runeSecond[i]]--
	}

	for _, val := range counter {
		if val != 0 {
			return false
		}
	}

	return true
}

func searchAnagrams(words *[]string) *map[string]*[]string {
	anagrams := make(map[string]*[]string)

	for _, word := range *words {
		lower := strings.ToLower(word)

		if _, ok := anagrams[word]; !ok {
			var anagramFound bool

			for key := range anagrams {
				anagramFound = checkAnagram(key, lower)
				if anagramFound {
					*anagrams[key] = append(*anagrams[key], lower)
				}
			}

			if !anagramFound {
				anagrams[lower] = new([]string)
				*anagrams[lower] = append(*anagrams[lower], lower)
			}
		}
	}

	for _, val := range anagrams {
		sort.Strings(*val)
	}

	return &anagrams
}

func main() {
	res := searchAnagrams(&[]string{"тяпка", "пятка", "пятак", "листок", "слиток", "столик"})

	for key, val := range *res {
		fmt.Printf("%s: %q\n", key, val)
	}
}
