package main

import (
	"fmt"

	"github.com/RokiLord/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Thrid Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("Block Hash: %s\n", block.Hash)
		fmt.Printf("Block PrevHash: %s\n", block.PrevHash)
	}
}
