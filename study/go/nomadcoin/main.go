package main

import "fmt"

func main() {
	x := 45245235235

	fmt.Printf("%b\n", x)
	fmt.Printf("%o\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%U\n", x)

	xAsBinary := fmt.Sprintf("%b", x)
	fmt.Println(xAsBinary)

}
