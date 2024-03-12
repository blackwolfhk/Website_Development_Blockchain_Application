package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

func calculateHash(index int, timestamp string, data string, prevHash string) string {
	hashInput := fmt.Sprintf("%d%s%s%s", index, timestamp, data, prevHash)
	hash := sha256.New()
	hash.Write([]byte(hashInput))
	return hex.EncodeToString(hash.Sum((nil)))
}

func createBlock(prevBlock Block, data string) Block {
	timestamp := time.Now().Format(time.RFC3339)
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: timestamp,
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock.Index, newBlock.Timestamp, newBlock.Data, newBlock.PrevHash)
	return newBlock

}

func main() {
	// Genesis block (first block in the blockchain).
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock.Index, genesisBlock.Timestamp, genesisBlock.Data, genesisBlock.PrevHash)

	fmt.Printf("Genesis Block: %+v\n", genesisBlock)

	// Create and display additional blocks.
	block1 := createBlock(genesisBlock, "Data for Block 1")
	block2 := createBlock(block1, "Data for Block 2")

	fmt.Printf("Block 1: %+v\n", block1)
	fmt.Printf("Block 2: %+v\n", block2)
}
