package main

import (
	"fmt"

	"github.com/RokiLord/nomadcoin/person"
)

func main() {
	terry := person.Person{}
	terry.SetDetails("terry", 12)
	fmt.Println(terry)
}
