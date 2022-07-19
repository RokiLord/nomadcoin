package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	genesisBlcok := block{"Genesis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlcok.data + genesisBlcok.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlcok.hash = hexHash
	fmt.Println(genesisBlcok)
}
