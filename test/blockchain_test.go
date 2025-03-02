package test

import (
	"testing"

	"github.com/nulln0ne/tinypow/core"
)

func TestNewTransaction(t *testing.T) {
	tx := core.NewTransaction("John", "Jane", 50)

	if tx.Sender != "John" {
		t.Errorf("Expected sender 'John', got '%s'", tx.Sender)
	}
	if tx.Recipient != "Jane" {
		t.Errorf("Expected recipient 'Jane', got '%s'", tx.Recipient)
	}
	if tx.Amount != 50 {
		t.Errorf("Expected amount 50, got %d", tx.Amount)
	}
	if tx.Hash == "" {
		t.Error("Transaction hash is empty")
	}
	if tx.Timestamp == 0 {
		t.Error("Timestamp is not set")
	}
}

func TestBlockCreationAndValidation(t *testing.T) {
	blockchain := core.NewBlockchain(4)
	genesisBlock := blockchain.Blocks[0]

	if genesisBlock.Index != 0 {
		t.Errorf("Genesis block index expected 0, got %d", genesisBlock.Index)
	}
	if genesisBlock.Hash == "" {
		t.Error("Genesis block hash is empty")
	}

	tx1 := core.NewTransaction("John", "Jane", 100)
	tx2 := core.NewTransaction("Jane", "John", 50)
	block := blockchain.AddBlock([]core.Transaction{*tx1, *tx2})

	if block.Index != 1 {
		t.Errorf("Block index expected 1, got %d", block.Index)
	}
	if block.Hash == "" {
		t.Error("Block hash is empty")
	}

	if !blockchain.IsChainValid() {
		t.Error("Blockchain should be valid")
	}
}

func TestChainInvalidation(t *testing.T) {
	blockchain := core.NewBlockchain(4)
	tx := core.NewTransaction("John", "Jane", 100)
	block := blockchain.AddBlock([]core.Transaction{*tx})

	block.Transactions[0].Amount = 9999
	block.Hash = core.CalculateHash(*block)

	if blockchain.IsChainValid() {
		t.Error("Blockchain should be invalid after tampering with a block")
	}
}
