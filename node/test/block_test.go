package test

import (
	"testing"

	"github.com/acor12/AngarCoin/node"
	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {
	block := node.NewBlock(0, 0, []byte{})
	assert.NotEqual(t, len(block.Serialize()), 0)
}

func TestDeserialize(t *testing.T) {
	block := node.NewBlock(0, 0, []byte{})
	serialized := block.Serialize()
	assert.Equal(t, block.Index, node.Deserialize(serialized).Index)
}

func TestNewBlock(t *testing.T) {
	block := node.NewBlock(0, 4, []byte{})
	assert.NotNil(t, block.DataHash)
	assert.Nil(t, block.Hash)
}

func TestGeneratedGenesisBlock(t *testing.T) {
	block := node.GeneratedGenesisBlock()
	assert.Equal(t, block.Index, uint(0))
	assert.Equal(t, block.Difficulty, uint8(0))
	assert.Equal(t, block.Transaction, []uint8([]byte(nil)))
	assert.Equal(t, block.PrevHash, []byte{})
	assert.Equal(t, block.MinedBy, []byte{})
	assert.Equal(t, block.Nonce, uint(0))
	assert.Equal(t, block.DateCreated, int64(1576895721))
}
