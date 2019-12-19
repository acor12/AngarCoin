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

	assert.NotEmpty(t, blockchain.Tip)
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
