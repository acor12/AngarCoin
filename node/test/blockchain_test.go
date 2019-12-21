package test

import (
	"os"
	"testing"

	"github.com/acor12/AngarCoin/node"
	"github.com/stretchr/testify/assert"
)

func TestBlockchainInit(t *testing.T) {
	blockchain := node.InitializeBlockchain()
	defer blockchain.Close()

	assert.NotEmpty(t, blockchain.Tip, blockchain.Tip)
	_, err := os.Stat("./blockchain")
	assert.Nil(t, err)
	os.Remove("./blockchain")
}

func TestLoadBlockchain(t *testing.T) {
	blockchain := node.InitializeBlockchain()
	blockchain.Close()

	blockchain = node.InitializeBlockchain()
	defer blockchain.Close()

	assert.NotEmpty(t, blockchain.Tip)
	_, err := os.Stat("./blockchain")
	assert.Nil(t, err)
	os.Remove("./blockchain")

}

func TestGetBlock(t *testing.T) {
	blockchain := node.InitializeBlockchain()
	defer blockchain.Close()

	tip := blockchain.Tip

	block := blockchain.GetBlock(tip)

	assert.NotNil(t, block)
	assert.Equal(t, block.Hash, tip)
	_, err := os.Stat("./blockchain")
	assert.Nil(t, err)
	os.Remove("./blockchain")
}
