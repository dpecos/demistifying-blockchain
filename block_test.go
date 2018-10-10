package main

import (
	"time"
	"testing"
)

var blockchain *BlockChain

func createBlock() *Block {
	if blockchain == nil {
		blockchain = &BlockChain{
			Challenge: "0",
			Blocks:    make(map[string]*Block),
		}
	}

	blockchain.Append("Test block")

	return blockchain.LastBlock
}

func TestHashChangesForNumber(t *testing.T) {
	block := createBlock()
	hash := block.Hash
	block.Number = 999
	newHash := calculateHash(block)

	if hash == newHash && hash != "" && newHash != "" {
		t.Errorf("checksum does not consider block's 'Number'")
	}
}

func TestHashChangesForTS(t *testing.T) {
	block := createBlock()
	hash := block.Hash
	block.Timestamp = time.Now()
	newHash := calculateHash(block)

	if hash == newHash && hash != "" && newHash != "" {
		t.Errorf("checksum does not consider block's 'Timestamp'")
	}
}

func TestHashChangesForData(t *testing.T) {
	block := createBlock()
	hash := block.Hash
	block.Data = "A different value"
	newHash := calculateHash(block)

	if hash == newHash && hash != "" && newHash != "" {
		t.Errorf("checksum does not consider block's 'Data'")
	}
}

func TestHashChangesForNonce(t *testing.T) {
	block := createBlock()
	hash := block.Hash
	block.Nonce = -1
	newHash := calculateHash(block)

	if hash == newHash && hash != "" && newHash != "" {
		t.Errorf("checksum does not consider block's 'Nonce'")
	}
}

func TestHashChangesForPrevious(t *testing.T) {
	block := createBlock()
	hash := block.Hash
	block.Previous = "this-is-not-a-hash"
	newHash := calculateHash(block)

	if hash == newHash && hash != "" && newHash != "" {
		t.Errorf("checksum does not consider block's 'Previous'")
	}
}