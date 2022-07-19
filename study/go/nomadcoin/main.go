package main

import "fmt"

func main() {
	foods := [4]string{"potato", "pizza", "pasta"}
	for _, i := range foods {
		fmt.Println(i)
	}

	slice := []string{"potato", "pizza", "pasta"}

	fmt.Printf("%v", slice)

	slice = append(slice, "tomato")
	fmt.Printf("%v", slice)
	fmt.Println(len(foods))
	fmt.Println(len(slice))
}
