package blockchain

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
		},
		[]byte{},
	)
	return data
}//end initData