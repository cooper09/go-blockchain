package blockchain

import (
	"bytes"
	"crypto/sha256"
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