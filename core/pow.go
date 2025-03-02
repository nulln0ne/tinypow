package core

import "math/big"

// ProofOfWork is a struct that represents the proof of work algorithm
// It contains a difficulty level and a target hash
type ProofOfWork struct {
	Difficulty int      `json:"difficulty"`
	Target     *big.Int `json:"target"`
}

// NewProofOfWork creates a new ProofOfWork object.
// It takes a difficulty level and returns a pointer to a ProofOfWork object.
// The difficulty level is the number of leading zeros in the target hash.
func NewProofOfWork(difficulty int) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty*4))

	return &ProofOfWork{
		Difficulty: difficulty,
		Target:     target,
	}
}

// isValidHash checks if the hash is valid.
// It takes a hash and returns a boolean.
// The hash is valid if it is less than the target.
func (pow *ProofOfWork) isValidHash(hash string) bool {
	hashInt := new(big.Int)
	hashInt.SetString(hash, 16)
	return hashInt.Cmp(pow.Target) == -1
}

// Mine mines a block.
// It takes a pointer to a Block object and mines the block.
// It returns the nonce and the hash of the block.
func (pow *ProofOfWork) Mine(block *Block) (int, string) {
	for {
		block.Nonce++
		block.Hash = CalculateHash(*block)
		if pow.isValidHash(block.Hash) {
			break
		}
	}
	return block.Nonce, block.Hash
}

// Validate validates a block
// It takes a hash and returns a boolean
// The block is valid if the hash is less than the target.
func (pow *ProofOfWork) Validate(hash string) bool {
	hashInt := new(big.Int)
	hashInt.SetString(hash, 16)

	return hashInt.Cmp(pow.Target) == -1
}
