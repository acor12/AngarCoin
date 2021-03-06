package node

import (
	"crypto/sha256"
	"encoding/json"
)

// Block keeps block headers
type Block struct {
	Index       uint
	Transaction []byte
	Difficulty  uint8
	PrevHash    []byte
	MinedBy     []byte
	DataHash    []byte
	Nonce       uint
	DateCreated int64
	Hash        []byte
}

//GeneratedGenesisBlock  method
func GeneratedGenesisBlock() *Block {
	dataHash := setDataHash(1, 0, []byte{}, []byte{})

	return &Block{
		Index:       1,
		Difficulty:  0,
		Transaction: nil,
		PrevHash:    []byte{},
		MinedBy:     []byte{},
		Nonce:       0,
		DateCreated: 1576895721,
		DataHash:    dataHash,
		Hash:        setHash(dataHash, 0, 1576895721),
	}
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

// SetMinedData set the nonce, the date created and the hash to the block
// return sfalse if the block is invalid otherwise returns true
func (b *Block) SetMinedData(nonce uint, dateCreated int64) bool {
	b.Nonce = nonce
	b.DateCreated = dateCreated

	return b.IsValid()
}

// IsValid returns true if the block hash valid data
func (b *Block) IsValid() bool {
	//TODO: complete this method
	if b.Nonce == 0 || b.Difficulty == 0 ||
		len(b.PrevHash) == 0 || len(b.DataHash) == 0 {
		return false
	}

	return true
}

// setHash calculates hash
func setHash(dataHash []byte, nonce uint64, DataCreated int64) []byte {
	blockData := map[string]interface{}{
		"Nonce":       nonce,
		"DataCreated": DataCreated,
		"DataHash":    dataHash,
	}

	jsonBlock, _ := json.Marshal(blockData)
	hasher := sha256.New()
	hasher.Write(jsonBlock)

	return hasher.Sum(nil)
}

// setDataHash calculates data hash
func setDataHash(index uint, difficulty uint8, prevHash []byte, minedBy []byte) []byte {

	dataHash := map[string]interface{}{
		"Index":       index,
		"Difficulty":  difficulty,
		"Transaction": nil,
		"PrevHash":    prevHash,
		"MinedBy":     minedBy,
	}
	jsonDataHash, _ := json.Marshal(dataHash)
	hasher := sha256.New()
	hasher.Write(jsonDataHash)

	return hasher.Sum(nil)
}

// NewBlock creates and returns Block
func NewBlock(index uint, difficulty uint8, prevHash []byte, minedBy []byte) *Block {

	block := &Block{
		Index:    index,
		PrevHash: prevHash,
		DataHash: setDataHash(index, difficulty, prevHash, minedBy),
	}

	return block
}
