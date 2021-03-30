package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/cooper09/go-blockchain/blockchain"
)

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s/n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}//end for
}//end mainrun main