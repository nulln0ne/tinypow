package core

import (
	"encoding/hex"
	"strconv"
	"time"
)

// Transaction is a struct that represents a transaction in the blockchain.
// It contains the sender, recipient, amount, timestamp, and hash.
type Transaction struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    int    `json:"amount"`
	Timestamp int64  `json:"timestamp"`
	Hash      string `json:"hash"`
}

// CalculateHash calculates the hash of the transaction.
// It takes a pointer to a Transaction and returns a string.
// The hash is calculated using the SHA-256 hashing algorithm.
// The hash is encoded to a hexadecimal string.
func (tx *Transaction) CalculateHash() string {
	record := tx.Sender + tx.Recipient + strconv.Itoa(tx.Amount) + strconv.FormatInt(tx.Timestamp, 10)
	hash := Sha256([]byte(record))

	return hex.EncodeToString(hash[:])
}

// NewTransaction creates a new transaction.
// It takes a sender, recipient, and amount and returns a pointer to a Transaction.
// The timestamp is set to the current time.
// The hash is calculated using the CalculateHash method.
func NewTransaction(sender, recipient string, amount int) *Transaction {
	tx := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
	tx.Timestamp = time.Now().Unix()
	tx.Hash = tx.CalculateHash()

	return &tx
}
