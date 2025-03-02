package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nulln0ne/tinypow/core"
	"github.com/nulln0ne/tinypow/http"
	"github.com/nulln0ne/tinypow/mempool"
	"github.com/nulln0ne/tinypow/miner"
	"github.com/nulln0ne/tinypow/storage"
	"github.com/syndtr/goleveldb/leveldb"
)

// loadBlockchain loads the blockchain from the database.
// It returns a pointer to a Blockchain.
func loadBlockchain(store *storage.BlockStore, difficulty int) *core.Blockchain {
	blockchain := &core.Blockchain{
		Difficulty: difficulty,
	}
	index := 0
	for {
		block, err := store.GetBlock(index)
		if err != nil {
			if err == leveldb.ErrNotFound {
				log.Printf("No block found for index %d. Ending blockchain load.", index)
				break
			}
			log.Printf("Error loading block %d: %v", index, err)
			break
		}
		log.Printf("Loaded block %d with hash: %s", index, block.Hash)
		blockchain.Blocks = append(blockchain.Blocks, *block)
		index++
	}
	if len(blockchain.Blocks) == 0 {
		log.Println("No existing blockchain found. Creating genesis block.")
		blockchain = core.NewBlockchain(difficulty)
		if err := store.SaveBlock(&blockchain.Blocks[0]); err != nil {
			log.Fatalf("Failed to save genesis block: %v", err)
		}
	}
	return blockchain
}

func main() {
	store, err := storage.NewBlockStore("storage/tinypow.db")
	if err != nil {
		log.Fatalf("Failed to create block store: %v", err)
	}
	defer store.Close()

	difficulty := 4

	bc := loadBlockchain(store, difficulty)
	mp := mempool.NewMempool()

	miner.StartMining(bc, mp, store, 10*time.Second)

	server := http.NewServer(bc, mp, store)

	errCh := make(chan error, 1)
	go func() {
		errCh <- server.Start(":8080")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errCh:
		log.Fatalf("Server error: %v", err)
	case <-c:
		log.Println("Shutting down...")
	}
}
