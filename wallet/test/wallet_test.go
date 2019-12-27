package test

import (
	"github.com/acor12/AngarCoin/wallet"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMnemonic(t *testing.T) {
	mnemonic := wallet.GenerateMnemonic()
	assert.NotEmpty(t, mnemonic)
}

func TestGenerateWallet(t *testing.T) {
	mnemonic := wallet.GenerateMnemonic()
	assert.NotEmpty(t, mnemonic)

	wallet, err := wallet.GenerateWalletFromMnemonic(mnemonic)

	assert.Nil(t, err)
	assert.NotNil(t, wallet)
}

func TestSign(t *testing.T) {
	mnemonic := wallet.GenerateMnemonic()
	assert.NotEmpty(t, mnemonic)

	wallet, err := wallet.GenerateWalletFromMnemonic(mnemonic)

	assert.Nil(t, err)
	assert.NotNil(t, wallet)

	data := []byte("data test")

	signature, err := wallet.Sign(data)
	assert.Nil(t, err)
	assert.NotNil(t, signature)
}

func TestValidateSignatureSign(t *testing.T) {
	mnemonic := wallet.GenerateMnemonic()
	assert.NotEmpty(t, mnemonic)

	account, err := wallet.GenerateWalletFromMnemonic(mnemonic)

	assert.Nil(t, err)
	assert.NotNil(t, account)

	data := []byte("data test")

	signature, err := account.Sign(data)
	assert.Nil(t, err)
	assert.NotNil(t, signature)

	assert.True(t, wallet.VerifySignature(account.PublicKey, data, signature))
}
