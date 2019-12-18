package node

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

// Block keeps block headers
type Block struct {
	Index         uint
	Difficulty    int8
	PrevBlockHash []byte
	MinedBy       []byte
	BlockDataHash []byte
	Nonce         uint
	DataCreated   int64
	BlockHash     []byte
}

// SetHash calculates and sets block hash
func SetHash(PrevBlockHash []byte, blockDataHash []byte, DataCreated int64) []byte {
	blockData := map[string]interface{}{
		"PrevBlockHash": PrevBlockHash,
		"blockDataHash": blockDataHash,
		"DataCreated":   DataCreated,
	}
	jsonData, _ := json.Marshal(blockData)
	hasher := sha256.New()
	hasher.Write(jsonData)

	return hasher.Sum(nil)
}

// NewBlock creates and returns Block
func NewBlock(blockDataHash []byte, prevBlockHash []byte) *Block {
	block := &Block{
		PrevBlockHash: prevBlockHash,
		BlockDataHash: blockDataHash,
		DataCreated:   time.Now().Unix(),
	}

	SetHash(prevBlockHash, blockDataHash, block.DataCreated)

	return block
}
