package mempool

import (
	"sync"

	"github.com/nulln0ne/tinypow/core"
)

// Mempool is a struct that represents a mempool for storing pending transactions.
// It contains a slice of transactions and a mutex for thread safety.
type Mempool struct {
	txs []core.Transaction
	mu  sync.Mutex
}

// NewMempool creates a new Mempool.
// It returns a pointer to a Mempool.
func NewMempool() *Mempool {
	return &Mempool{
		txs: make([]core.Transaction, 0),
	}
}

// AddTransaction adds a transaction to the mempool.
// It takes a pointer to a Transaction and adds it to the mempool.
func (mp *Mempool) AddTransaction(tx *core.Transaction) {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	mp.txs = append(mp.txs, *tx)
}

// GetTransactions returns a copy of the transactions in the mempool.
// It returns a slice of transactions.
func (mp *Mempool) GetTransactions() []core.Transaction {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	txsCopy := make([]core.Transaction, len(mp.txs))
	copy(txsCopy, mp.txs)
	return txsCopy
}

// ClearTransactions clears the transactions in the mempool.
func (mp *Mempool) ClearTransactions() {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	mp.txs = []core.Transaction{}
}
