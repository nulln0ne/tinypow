package storage

import (
	"encoding/json"
	"fmt"

	"github.com/nulln0ne/tinypow/core"
	"github.com/syndtr/goleveldb/leveldb"
)

type BlockStore struct {
	db *leveldb.DB
}

func NewBlockStore(path string) (*BlockStore, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return &BlockStore{db: db}, nil
}
func (store *BlockStore) SaveBlock(block *core.Block) error {
	key := fmt.Sprintf("block_%d", block.Index)
	data, err := json.Marshal(block)
	if err != nil {
		return err
	}

	return store.db.Put([]byte(key), data, nil)
}

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

func (store *BlockStore) Close() error {
	return store.db.Close()
}
