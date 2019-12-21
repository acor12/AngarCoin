package node

import (
	"crypto/sha256"
	"encoding/json"
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

//Serialize block into bytes
func (b *Block) Serialize() []byte {
	data, _ := json.Marshal(b)
	return data
}

//Deserialize encoded block
func Deserialize(serializedBlock []byte) *Block {
	block := new(Block)
	json.Unmarshal(serializedBlock, block)
	return block
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
func NewBlock(index uint, difficulty uint8, prevHash string) *Block {

	block := &Block{
		Index:    index,
		PrevHash: []byte(prevHash),
		DataHash: setDataHash(prevHash),
	}

	return block
}
