package node

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

// Block keeps block headers
type Block struct {
	Index       uint
	Difficulty  uint8
	PrevHash    []byte
	MinedBy     []byte
	DataHash    []byte
	Nonce       uint
	DataCreated int64
	Hash        []byte
}

// setHash calculates hash
func setHash(dataHash []byte, data string, DataCreated int64) []byte {
	blockData := map[string]interface{}{
		"Nonce":       nil,
		"DataCreated": DataCreated,
		"DataHash":    dataHash,
	}
	jsonBlock, _ := json.Marshal(blockData)
	hasher := sha256.New()
	hasher.Write(jsonBlock)

	return hasher.Sum(nil)
}

// setDataHash calculates data hash
func setDataHash(prevHash string) []byte {

	dataHash := map[string]interface{}{
		"Index":       nil,
		"Difficulty":  nil,
		"Transaction": nil,
		"PrevHash":    prevHash,
	}
	jsonDataHash, _ := json.Marshal(dataHash)
	hasher := sha256.New()
	hasher.Write(jsonDataHash)
	return hasher.Sum(nil)
}

// NewBlock creates and returns Block
func NewBlock(data string, prevHash string) *Block {

	dataCreated := time.Now().Unix()
	dataHash := setDataHash(prevHash)
	hash := setHash(dataHash, data, dataCreated)

	block := &Block{
		PrevHash:    []byte(prevHash),
		DataHash:    dataHash,
		DataCreated: dataCreated,
		Hash:        hash,
	}

	return block
}
