package transaction

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/vivijj/ziongo/crypto/babyjub"
	"github.com/vivijj/ziongo/crypto/poseidon"
)

var (
	withdrawHasher = poseidon.NewHasher(9)
)

// WithdrawTx  perform a withdrawal of funds from L2 account to L1 account
type WithdrawTx struct {
	// account id of the transaction initiator.
	AccountId  int
	Nonce      int
	ValidUntil int
	FeeToken   int
	Fee        *big.Int

	From            common.Address
	To              common.Address
	Token           int
	Amount          *big.Int
	Signature       babyjub.Signature
	MinGas          *big.Int
	ExtraData       []byte
	OnchainDataHash common.Hash
}

func (tx WithdrawTx) isZionTx() {}

// EncodeBi Encode the transaction data as the *big.Int by using poseidon hash
func (tx WithdrawTx) EncodeBi() *big.Int {
	var out []*big.Int
	out = append(out, big.NewInt(int64(tx.AccountId)))
	out = append(out, big.NewInt(int64(tx.Token)))
	out = append(out, tx.Amount)
	out = append(out, big.NewInt(int64(tx.FeeToken)))
	out = append(out, tx.Fee)
	out = append(out, tx.OnchainDataHash.Big())
	out = append(out, big.NewInt(int64(tx.ValidUntil)))
	out = append(out, big.NewInt(int64(tx.Nonce)))

	return withdrawHasher.HashBi(out)
}

// GetBytes Encode the transaction data as the byte sequence
func (tx WithdrawTx) GetBytes() (out []byte) {

	out = append(out, []byte(Withdraw)...)
	out = append(out, IntToBytes(tx.AccountId)...)
	out = append(out, IntToBytes(tx.Nonce)...)
	out = append(out, IntToBytes(tx.ValidUntil)...)
	out = append(out, IntToBytes(tx.FeeToken)...)
	out = append(out, tx.Fee.Bytes()...)
	out = append(out, tx.From.Bytes()...)
	out = append(out, tx.To.Bytes()...)
	out = append(out, IntToBytes(tx.Token)...)
	out = append(out, tx.Amount.Bytes()...)
	return
}
