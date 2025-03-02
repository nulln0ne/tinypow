package storage

import (
	"encoding/json"
	"fmt"

	"github.com/nulln0ne/tinypow/core"
	"github.com/syndtr/goleveldb/leveldb"
)

// BlockStore is a struct that represents a block store.
// It contains a pointer to a leveldb.DB.
type BlockStore struct {
	db *leveldb.DB
}

// NewBlockStore creates a new block store.
// It takes a path and returns a pointer to a BlockStore and an error.
// It opens a leveldb database and returns a pointer to a BlockStore.
func NewBlockStore(path string) (*BlockStore, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return &BlockStore{db: db}, nil
}

// SaveBlock saves a block to the database.
// It takes a pointer to a Block and returns an error.
// It marshals the block and saves it to the database.
func (store *BlockStore) SaveBlock(block *core.Block) error {
	key := fmt.Sprintf("block_%d", block.Index)
	data, err := json.Marshal(block)
	if err != nil {
		return err
	}

	return store.db.Put([]byte(key), data, nil)
}

// GetBlock gets a block from the database.
// It takes an index and returns a pointer to a Block and an error.
// It unmarshals the block and returns it.
func (store *BlockStore) GetBlock(index int) (*core.Block, error) {
	key := fmt.Sprintf("block_%d", index)
	data, err := store.db.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}

	var block core.Block
	if err := json.Unmarshal(data, &block); err != nil {
		return nil, err
	}

	return &block, nil
}

// Close closes the block store.
// It takes a pointer to a BlockStore and returns an error.
// It closes the leveldb database.
func (store *BlockStore) Close() error {
	return store.db.Close()
}
