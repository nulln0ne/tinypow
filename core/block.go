package core

import (
	"strconv"
)

// Block is a struct that represents a block in the blockchain.
// It contains the index, timestamp, transactions, previous hash, hash, and nonce.
type Block struct {
	Index        int           `json:"index"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     string        `json:"prevHash"`
	Hash         string        `json:"hash"`
	Nonce        int           `json:"nonce"`
}

// getBlockRecord returns a string representation of a block.
// It takes a Block object and returns a string.
// The string is a concatenation of the block's index, timestamp, merkle root, previous hash, and nonce.
func getBlockRecord(b Block) string {
	merkleRoot := CalculateMerkleRoot(b.Transactions)
	return strconv.Itoa(b.Index) +
		strconv.FormatInt(b.Timestamp, 10) +
		merkleRoot +
		b.PrevHash +
		strconv.Itoa(b.Nonce)
}

// CalculateBlockHash calculates the hash of a block.
// It takes a Block object and returns a string.
// The hash is calculated using the SHA-256 hashing algorithm.
// The hash is encoded to a hexadecimal string.
func CalculateHash(b Block) string {
	record := getBlockRecord(b)

	return Sha256Hash(record)
}

// IsBlockValid checks if a block is valid by verifying the index, previous hash, and the integrity of each transaction.
// It takes new and old Block objects and returns a boolean.
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	for _, tx := range newBlock.Transactions {
		if tx.CalculateHash() != tx.Hash {
			return false
		}
	}

	return CalculateHash(newBlock) == newBlock.Hash
}
