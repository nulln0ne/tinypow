# TinyPow

A lightweight Proof-of-Work blockchain implementation in Go.

## Overview

TinyPow is a minimalist blockchain implementation that demonstrates the core principles of a blockchain with Proof-of-Work consensus. It includes implementations of:

- Blockchain data structure
- Blocks with transactions
- Merkle tree for transaction verification
- SHA-256 hashing
- Proof-of-Work mining algorithm

This project serves as an educational resource for understanding how blockchains work under the hood.

## Features

- **Blockchain Structure**: Complete blockchain implementation with blocks, transactions, and hashing.
- **Proof-of-Work Mining**: Adjustable difficulty mining algorithm.
- **Transaction System**: Simple transaction model with sender, recipient, and amount.
- **Merkle Tree**: Efficient verification of transaction integrity.
- **Chain Validation**: Methods to verify the integrity of the entire blockchain.

## Installation

```bash
git clone https://github.com/nulln0ne/tinypow.git
cd tinypow
go mod tidy
```


## Core Components

- **Block**: The basic unit of the blockchain, containing transactions and metadata.
- **Transaction**: Represents a transfer of value from one user to another.
- **Blockchain**: A chain of blocks linked by cryptographic hashes.
- **ProofOfWork**: The consensus algorithm that secures the blockchain.

## Implementation Details

### Proof-of-Work Algorithm

The Proof-of-Work algorithm requires miners to find a hash value below a target, which is determined by the difficulty level. The difficulty determines how many leading zeros are required in the hash.

### Block Validation

Each block is validated by:
1. Checking that its index is one more than the previous block's index
2. Verifying that its previous hash matches the hash of the previous block
3. Ensuring all transactions are valid
4. Confirming that the block's hash is valid according to the PoW algorithm

## Testing

Run the tests with:

```bash
go test ./...
```
