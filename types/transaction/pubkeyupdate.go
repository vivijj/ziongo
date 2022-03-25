package transaction

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/vivijj/ziongo/crypto/babyjub"
)

// PubkeyUpdateTx will set the owner's public key associated with the account.
// without public key set, account is unable to execute any L2 transactions.
type PubkeyUpdateTx struct {
	AccountId  int
	Nonce      int
	ValidUntil int
	FeeToken   int
	Fee        *big.Int
	Account    common.Address
	PubKey     babyjub.PublicKey
	AuthData   []byte
}

func (tx PubkeyUpdateTx) isZionTx() {}

func (tx PubkeyUpdateTx) GetBytes() (out []byte) {

	out = append(out, []byte(PubKeyUpdate)...)
	out = append(out, IntToBytes(tx.AccountId)...)
	out = append(out, IntToBytes(tx.Nonce)...)
	out = append(out, IntToBytes(tx.ValidUntil)...)
	out = append(out, IntToBytes(tx.FeeToken)...)
	out = append(out, tx.Fee.Bytes()...)
	out = append(out, tx.Account.Bytes()...)
	out = append(out, tx.PubKey.X.Bytes()...)
	out = append(out, tx.PubKey.Y.Bytes()...)

	return
}

func (tx PubkeyUpdateTx) HashEncodeData() []byte {
	return crypto.Keccak256(tx.GetBytes())
}

func (tx PubkeyUpdateTx) IsAuthDataValid() bool {
	userAddr := tx.Account
	sig := tx.AuthData
	msgHash := tx.HashEncodeData()
	fmt.Println("msgHash is: ", msgHash)
	if len(sig) != 65 {
		return false
	}
	if sig[64] != 27 && sig[64] != 28 {
		return false
	}
	sig[64] -= 27
	pubkey, err := crypto.SigToPub(msgHash, sig)
	if err != nil {
		return false

	}
	recoverAddr := crypto.PubkeyToAddress(*pubkey)
	fmt.Println("recover address is: ", recoverAddr)
	if !bytes.Equal(recoverAddr.Bytes(), userAddr.Bytes()) {
		return false
	}
	return true
}
