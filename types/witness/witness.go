package witness

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/vivijj/ziongo/crypto/babyjub"
	"github.com/vivijj/ziongo/types/operation"
	"github.com/vivijj/ziongo/utils/hasher"
)

var (
	BalanceHasher = hasher.NewPoseidonHasher(5)
)

type BalanceLeaf struct {
	Balance *big.Int
}

func (b BalanceLeaf) Hash() *big.Int {
	return BalanceHasher.HashBi([]*big.Int{b.Balance})
}

type AccountNode struct {
	Address     common.Address
	PublicKeyX  string
	PublicKeyY  string
	Nonce       int
	BalanceRoot string
}

type BalanceUpdateWitness struct {
	TokenId    int
	Proof      []string
	RootBefore string
	RootAfter  string
	Before     BalanceLeaf
	After      BalanceLeaf
}

type AccountUpdateWitness struct {
	AccountId     int
	Proof         []string
	RootBefore    string
	RootAfter     string
	AccountBefore AccountNode
	AccountAfter  AccountNode
}

type AccountOriginWitness struct {
	AccountUpdate     AccountUpdateWitness
	BalanceUpdateMain BalanceUpdateWitness
	BalanceUpdateFee  BalanceUpdateWitness
}

type AccountOutcomeWitness struct {
	AccountUpdate AccountUpdateWitness
	BalanceUpdate BalanceUpdateWitness
}

type Witness struct {
	SignatureFrom     *babyjub.Signature
	AccountMerkleRoot string

	BalanceUpdateFrom    BalanceUpdateWitness
	BalanceUpdateFeeFrom BalanceUpdateWitness
	AccountUpdateFrom    AccountUpdateWitness

	BalanceUpdateTo BalanceUpdateWitness
	AccountUpdateTo AccountUpdateWitness

	BalanceUpdateOperator BalanceUpdateWitness
	AccountUpdateOperator AccountUpdateWitness

	NumConditionalTransactionAfter int
}

type TxWitness struct {
	TxType  string
	Tx      operation.ZionOp
	Witness Witness
}

type NoopWitness struct {
	TxType string
}
