package main

import (
	"fmt"
)

// ---------- Generic Function ----------
func First[T any](s []T) T {
	return s[0]
}

// ---------- Generic Type ----------
type Pair[T any, U any] struct {
	First  T
	Second U
}

func main() {
	// Using the generic function
	numbers := []int{10, 20, 30}
	words := []string{"Go", "Generics", "Demo"}

	fmt.Println("First number:", First(numbers)) // Output: 10
	fmt.Println("First word:", First(words))     // Output: Go

	// Using the generic type
	p1 := Pair[int, string]{First: 100, Second: "Hundred"}
	p2 := Pair[string, bool]{First: "FeatureEnabled", Second: true}

	fmt.Printf("Pair 1: %+v\n", p1) // Output: Pair 1: {First:100 Second:Hundred}
	fmt.Printf("Pair 2: %+v\n", p2) // Output: Pair 2: {First:FeatureEnabled Second:true}
}
