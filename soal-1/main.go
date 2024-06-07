package main

import (
	"fmt"
)

// Function to calculate the weights of all substrings
func calculateWeights(s string) map[int]struct{} {
	weights := make(map[int]struct{})
	count := make(map[rune]int)

	for _, char := range s {
		count[char] = count[char] + 1
		for length := 1; length <= count[char]; length++ {
			weight := int(char-'a'+1) * length
			weights[weight] = struct{}{}
		}
	}

	return weights
}

// Function to check the queries against the weights
func checkQueries(s string, queries []int) []string {
	weights := calculateWeights(s)
	results := make([]string, len(queries))

	for i, query := range queries {
		if _, exists := weights[query]; exists {
			results[i] = "Yes"
		} else {
			results[i] = "No"
		}
	}

	return results
}

func main() {
	s := "abbcccd"
	queries := []int{1, 3, 9, 8}

	results := checkQueries(s, queries)
	for _, result := range results {
		fmt.Println(result)
	}
}
