package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}//end BlockChain

type Block struct {
	Hash		[]byte
	Data		[]byte
	PrevHash 	[]byte
}//end Block

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}//end DeriveHash

func CreateBlock (data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash }
	block.DeriveHash()
	return block
}//end CreateBlock

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks) - 1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new )
}//end AddBlock

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}//end Genesis

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}//end InitBlockchain

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

	}//end for
}//end main