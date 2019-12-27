package wallet

import (
	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

//Wallet struct
type Wallet struct {
	PublicKey  []byte
	PrivateKey []byte
	Address    []byte
	wallet     *hdwallet.Wallet
	account    accounts.Account
}

//Sign the given data
func (w *Wallet) Sign(data []byte) []byte {
	signature, _ := w.wallet.SignText(w.account, data)
	return signature
}

// GenerateMnemonic return a new mnemonic string
func GenerateMnemonic() string {
	mnemonic, _ := hdwallet.NewMnemonic(256)
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
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		Address:    account.Address.Bytes(),
		wallet:     wallet,
		account:    account,
	}
}
