package http

import (
	"encoding/json"
	"net/http"

	"github.com/nulln0ne/tinypow/core"
	"github.com/nulln0ne/tinypow/mempool"
	"github.com/nulln0ne/tinypow/storage"
)

// Server is a struct that represents a server.
// It contains a pointer to a Blockchain, a pointer to a Mempool, and a pointer to a BlockStore.
type Server struct {
	Blockchain *core.Blockchain
	Mempool    *mempool.Mempool
	Store      *storage.BlockStore
}

// NewServer creates a new server.
// It takes a pointer to a Blockchain, a pointer to a Mempool, and a pointer to a BlockStore.
// It returns a pointer to a Server.
func NewServer(bc *core.Blockchain, mp *mempool.Mempool, store *storage.BlockStore) *Server {
	return &Server{
		Blockchain: bc,
		Mempool:    mp,
		Store:      store,
	}
}

// handleAddTransaction handles the add transaction request.
// It takes a pointer to a ResponseWriter and a pointer to a Request.
// It adds a transaction to the mempool.
func (s *Server) handleAddTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Sender    string `json:"sender"`
		Recipient string `json:"recipient"`
		Amount    int    `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tx := core.NewTransaction(req.Sender, req.Recipient, req.Amount)
	s.Mempool.AddTransaction(tx)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tx); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// handleGetBlocks handles the get blocks request.
// It takes a pointer to a ResponseWriter and a pointer to a Request.
// It returns the blockchain.
func (s *Server) handleGetBlocks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewEncoder(w).Encode(s.Blockchain.Blocks); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// Start starts the server.
// It takes an address and returns an error.
// It handles the requests and returns the response.
func (s *Server) Start(addr string) error {
	http.HandleFunc("/transactions", s.handleAddTransaction)
	http.HandleFunc("/blocks", s.handleGetBlocks)

	return http.ListenAndServe(addr, nil)
}
