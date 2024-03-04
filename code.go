package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)


type Block struct {
	Index     int
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
}

type Blockchain struct {
	Chain []Block
}

func (bc *Blockchain) NewBlock(data string) Block {
	var prevBlock Block
	if len(bc.Chain) > 0 {
		prevBlock = bc.Chain[len(bc.Chain)-1]
	}

	newBlock := Block{
		Index:     len(bc.Chain),
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = calculateBlockHash(newBlock)

	bc.Chain = append(bc.Chain, newBlock)
	return newBlock
}

func calculateBlockHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PrevHash)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}


func (bc *Blockchain) DisplayAllBlocks() {
	for _, block := range bc.Chain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Previous Hash: %s\n\n", block.PrevHash)
	}
}

// ModifyBlock modifies a block in the blockchain.
func (bc *Blockchain) ModifyBlock(index int, newData string) {
	if index >= 0 && index < len(bc.Chain) {
		block := &bc.Chain[index]
		block.Data = newData
		block.Hash = calculateBlockHash(*block)
	}
}

func main() {

	bc := Blockchain{}
	bc.NewBlock("First Block Data")
	bc.NewBlock("Second Block Data")

	fmt.Println("All Blocks in the Blockchain:")
	bc.DisplayAllBlocks()

	// Modify a block
	bc.ModifyBlock(1, "Modified Data for Second Block")

	// Display all blocks after modification
	fmt.Println("\nAll Blocks in the Blockchain after Modification:")
	bc.DisplayAllBlocks()
}
