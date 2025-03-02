package core

import (
	"encoding/hex"
)

// EmptyMerkleRoot represents the root value of an empty Merkle tree
const (
	EmptyMerkleRoot = ""
)

// CalculateMerkleRoot computes the Merkle root hash from a list of transactions
// Returns EmptyMerkleRoot if no transactions are provided
func CalculateMerkleRoot(transactions []Transaction) string {
	if len(transactions) == 0 {
		return EmptyMerkleRoot
	}

	hashes := make([][]byte, len(transactions))
	for i, tx := range transactions {
		hashBytes, err := hex.DecodeString(tx.Hash)
		if err != nil {
			hashBytes = Sha256([]byte(tx.Hash))
		}
		hashes[i] = hashBytes
	}

	rootHash := buildMerkleTree(hashes)
	return hex.EncodeToString(rootHash)
}

// buildMerkleTree recursively builds a Merkle tree from a list of hash values
// Returns the root hash of the tree
func buildMerkleTree(hashes [][]byte) []byte {
	hashCount := len(hashes)

	if hashCount == 0 {
		return Sha256([]byte{})
	}

	if hashCount == 1 {
		return hashes[0]
	}

	nextLevel := make([][]byte, 0, (hashCount+1)/2)

	for i := 0; i < hashCount; i += 2 {
		left := hashes[i]
		var right []byte

		if i+1 < hashCount {
			right = hashes[i+1]
		} else {
			right = left
		}

		combined := append(left, right...)
		nextLevel = append(nextLevel, Sha256(combined))
	}

	return buildMerkleTree(nextLevel)
}
