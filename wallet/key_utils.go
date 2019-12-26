package wallet

import (
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

//Wallet struct
type Wallet struct {
	publicKey  []byte
	privateKey []byte
	address    []byte
}

// GenerateMnemonic return a new mnemonic string
func GenerateMnemonic() string {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return mnemonic
}

//GenerateWallet returns new wallet from given mnemonic
func GenerateWallet(mnemonic string) *Wallet {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		println(err.Error())
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, _ := wallet.Derive(path, false)

	privateKey, _ := wallet.PrivateKeyBytes(account)
	publicKey, _ := wallet.PublicKeyBytes(account)

	return &Wallet{
		publicKey:  publicKey,
		privateKey: privateKey,
		address:    account.Address.Bytes(),
	}
}
