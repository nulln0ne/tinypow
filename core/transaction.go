package core

import (
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

// getTransactionRecord returns a string representation of a transaction.
// It takes a Transaction object and returns a string.
// The string is a concatenation of the transaction's sender, recipient, amount, and timestamp.
func getTransactionRecord(tx Transaction) string {
	return tx.Sender + tx.Recipient + strconv.Itoa(tx.Amount) + strconv.FormatInt(tx.Timestamp, 10)
}

// CalculateHash calculates the hash of the transaction.
// It takes a pointer to a Transaction and returns a string.
// The hash is calculated using the SHA-256 hashing algorithm.
// The hash is encoded to a hexadecimal string.
func (tx *Transaction) CalculateHash() string {
	record := getTransactionRecord(*tx)

	return Sha256Hash(record)
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
