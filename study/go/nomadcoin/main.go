package main

import (
	"github.com/RokiLord/nomadcoin/explorer"
	"github.com/RokiLord/nomadcoin/rest"
)

func main() {
	explorer.Start(4000)
	rest.Start(4000)
}
