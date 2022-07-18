package main

import "fmt"

func plus(a ...int) int {
	total := 0

	for _, item := range a {
		total += item
	}

	return total
}

func main() {
	result := plus(2, 3, 4, 5, 2, 3, 45, 5, 21, 4, 1, 34)

	fmt.Println(result)
}
