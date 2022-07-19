package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) sayHello() {
	fmt.Printf("Hello! My name is %s and I'm %d", p.name, p.age)
}

func main() {
	terry := person{"terry", 12}

	terry.sayHello()

}
