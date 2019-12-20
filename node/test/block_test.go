package test

import (
	"github.com/acor12/AngarCoin/node"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialize(t *testing.T) {
	block := node.NewBlock([]byte{}, []byte{})
	assert.NotEqual(t, len(block.Serialize()), 0)
}

func TestDeserialize(t *testing.T) {
	block := node.NewBlock([]byte{}, []byte{})
	serialized := block.Serialize()
	assert.Equal(t, block.Index, node.Deserialize(serialized).Index)
}
