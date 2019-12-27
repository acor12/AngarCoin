package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// Signature struct representation
type Signature struct {
	R, S *big.Int
}

// Wallet struct representation
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	address    common.Address
}

//Sign create ECDSA signature
func (w *Wallet) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, w.PrivateKey, data)
	if err != nil {
		return nil, err
	}

	return &Signature{
		R: r,
		S: s,
	}, nil
}

//PrivateKeyBytes returns private key in bytes
func (w *Wallet) PrivateKeyBytes() []byte {
	return crypto.FromECDSA(w.PrivateKey)
}

//PublicKeyBytes returns public key in bytes
func (w *Wallet) PublicKeyBytes() []byte {
	return crypto.FromECDSAPub(w.PublicKey)
}

//GenerateMnemonic generate new mnemonic
func GenerateMnemonic() string {
	mnemonic, _ := hdwallet.NewMnemonic(256)
	return mnemonic
}

//GenerateWalletFromMnemonic create new wallet from mnemonic
func GenerateWalletFromMnemonic(mnemonic string) (*Wallet, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, _ := wallet.Derive(path, false)

	privBytes, _ := wallet.PrivateKeyBytes(account)
	privateKey, _ := crypto.ToECDSA(privBytes)

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
		address:    crypto.PubkeyToAddress(privateKey.PublicKey),
	}, nil
}

//VerifySignature check if the signature is valid
func VerifySignature(publicKey *ecdsa.PublicKey, data []byte, signature *Signature) bool {
	return ecdsa.Verify(publicKey, data, signature.R, signature.S)
}
