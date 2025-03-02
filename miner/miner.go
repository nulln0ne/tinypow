package miner

import (
	"log"
	"time"

	"github.com/nulln0ne/tinypow/core"
	"github.com/nulln0ne/tinypow/mempool"
	"github.com/nulln0ne/tinypow/storage"
)

// StartMining starts the mining process.
// It takes a pointer to a Blockchain, a pointer to a Mempool, a pointer to a BlockStore, and an interval.
// It mines a block every interval and saves it to the BlockStore.
func StartMining(bc *core.Blockchain, mp *mempool.Mempool, store *storage.BlockStore, interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			txs := mp.GetTransactions()
			if len(txs) > 0 {
				newBlock := bc.AddBlock(txs)
				if err := store.SaveBlock(newBlock); err != nil {
					log.Printf("Failed to save block: %v", err)
				} else {
					log.Printf("Block mined and saved with hash: %s", newBlock.Hash)
				}
				mp.ClearTransactions()
			}
		}
	}()
}
