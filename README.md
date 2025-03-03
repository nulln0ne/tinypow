# TinyPoW

A lightweight Proof-of-Work blockchain implementation in Go.

## Overview

TinyPoW is a minimalist blockchain implementation that demonstrates the core principles of a blockchain with Proof-of-Work consensus. It includes implementations of:

- Blockchain data structure
- Blocks with transactions
- Merkle tree for transaction verification
- SHA-256 hashing
- Proof-of-Work mining algorithm
- Transaction mempool
- Persistent storage with LevelDB
- HTTP API for blockchain interaction


This project serves as an educational resource for understanding how blockchains work under the hood.

## Features

- **Blockchain Structure**: Complete blockchain implementation with blocks, transactions, and hashing.
- **Proof-of-Work Mining**: Adjustable difficulty mining algorithm with automatic block mining.
- **Transaction System**: Simple transaction model with sender, recipient, and amount.
- **Merkle Tree**: Efficient verification of transaction integrity.
- **Chain Validation**: Methods to verify the integrity of the entire blockchain.
- **Mempool**: Transaction queue for pending transactions.
- **Persistent Storage**: Blockchain data is stored using LevelDB.
- **HTTP API**: RESTful API for interacting with the blockchain.


## Installation

```bash
git clone https://github.com/nulln0ne/tinypow.git
cd tinypow
go mod tidy
```

## Usage

### Running the Node

```bash
go run cmd/tinypow/main.go
```

This will start:
- A blockchain node with the specified difficulty
- An HTTP server on port 8080
- A miner that creates new blocks every 10 seconds

### API Endpoints

#### Add a Transaction
```
POST /transactions

{
  "sender": "Alice",
  "recipient": "Bob",
  "amount": 100
}
```

#### Get All Blocks
```
GET /blocks
```

## Core Components

- **Block**: The basic unit of the blockchain, containing transactions and metadata.
- **Transaction**: Represents a transfer of value from one user to another.
- **Blockchain**: A chain of blocks linked by cryptographic hashes.
- **ProofOfWork**: The consensus algorithm that secures the blockchain.
- **Mempool**: Queue for storing pending transactions before they're mined into blocks.
- **BlockStore**: Persistent storage interface for saving blockchain data.
- **Miner**: Background process that creates new blocks at regular intervals.
- **Server**: HTTP server that exposes the blockchain API.


## Implementation Details

### Proof-of-Work Algorithm

The Proof-of-Work algorithm requires miners to find a hash value below a target, which is determined by the difficulty level. The difficulty determines how many leading zeros are required in the hash.

### Block Validation

Each block is validated by:
1. Checking that its index is one more than the previous block's index
2. Verifying that its previous hash matches the hash of the previous block
3. Ensuring all transactions are valid
4. Confirming that the block's hash is valid according to the PoW algorithm

### Storage

Blockchain data is persistently stored using LevelDB. When the application starts, it attempts to load existing blocks from the database. If no blocks are found, a new blockchain with a genesis block is created.

## Testing

Run the tests with:

```bash
go test ./...
```
