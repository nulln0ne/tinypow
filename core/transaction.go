package core

import (
	"encoding/hex"
	"strconv"
	"time"
)

type Transaction struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    int    `json:"amount"`
	Timestamp int64  `json:"timestamp"`
	Hash      string `json:"hash"`
}

func (tx *Transaction) CalculateHash() string {
	record := tx.Sender + tx.Recipient + strconv.Itoa(tx.Amount) + strconv.FormatInt(tx.Timestamp, 10)
	hash := Sha256([]byte(record))

	return hex.EncodeToString(hash[:])
}

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
