package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

//Take Data from Block

//create a counter (nonce) which starts at 0

//create a hash of the data plus the counter

//check the has to see if it meets a set of requirements

//Requirements - first few bytes of the hash must contain 0's. The more 0's the more difficult

const Difficulty = 12

type ProofOfWork struct {
	Block *Block
	Target *big.Int
}//end struct

func NewProof(b *Block ) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}//end ProofOfWork

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][] byte {
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}//end initData

func (pow *ProofOfWork) Run() (int, []byte) {
	 var intHash big.Int
	 var hash [32]byte

	 nonce := 0 
	 for nonce < math.MaxInt64 {
		 data = pow.InitData(nonce)
		hash = sha256.Sum256()

		fmt.Printf("\rx%", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp() == -1 {
			break
		} else {
			nonce++
		}
	 }//end for
	 fmt.Println()
	 return nonce, hash[:]
}//end Run()

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}//end validate

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.write(buff, binary.BigEndian, num )
	if err != nil {
		log.Panic("Can't convert to Hex: ", err )
	}
	return buff.Bytes
}//end ToHex