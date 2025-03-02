package core

import (
	"time"
)

// Blockchain is a struct that represents a blockchain.
// It contains the blocks and the difficulty.
type Blockchain struct {
	Blocks     []Block `json:"blocks"`
	Difficulty int     `json:"difficulty"`
}

// NewBlockchain creates a new blockchain.
// It takes a difficulty and returns a pointer to a Blockchain.
func NewBlockchain(difficulty int) *Blockchain {
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Transactions: []Transaction{},
		PrevHash:     "",
		Nonce:        0,
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)

	return &Blockchain{
		Blocks:     []Block{genesisBlock},
		Difficulty: difficulty,
	}
}

// GetLatestBlock returns the latest block in the blockchain.
// It takes a pointer to a Blockchain and returns a pointer to a Block.
func (bc *Blockchain) GetLatestBlock() *Block {
	return &bc.Blocks[len(bc.Blocks)-1]
}

// AddBlock adds a new block to the blockchain.
// It takes a slice of transactions and returns a pointer to a Block.
func (bc *Blockchain) AddBlock(transactions []Transaction) *Block {
	latestBlock := bc.GetLatestBlock()
	newBlock := Block{
		Index:        latestBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     latestBlock.Hash,
		Nonce:        0,
	}

	pow := NewProofOfWork(bc.Difficulty)
	pow.Mine(&newBlock)

	bc.Blocks = append(bc.Blocks, newBlock)

	return &newBlock
}

// IsChainValid checks if the blockchain is valid.
// It takes a pointer to a Blockchain and returns a boolean.
func (bc *Blockchain) IsChainValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if !IsBlockValid(currentBlock, previousBlock) {
			return false
		}

		pow := NewProofOfWork(bc.Difficulty)
		if !pow.Validate(currentBlock.Hash) {
			return false
		}
	}
	return true
}
