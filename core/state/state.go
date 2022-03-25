package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"

	"github.com/vivijj/ziongo/crypto/babyjub"
	"github.com/vivijj/ziongo/types/fr"
	"github.com/vivijj/ziongo/types/smt"
	"github.com/vivijj/ziongo/types/transaction"
)

type TransitionVariant struct {
	SigFrom *babyjub.Signature

	AccFromId      int
	AccFromAddr    common.Address
	AccFromPubkeyX string
	AccFromPubkeyY string
	AccFromNonce   int

	BalanceFromTokenMainId      int
	BalanceFromTokenMainBalance *big.Int
	BalanceFromTokenFeeId       int
	BalanceFromTokenFeeBalance  *big.Int

	AccToId                   int
	AccToAddr                 common.Address
	BalanceToTokenMainId      int
	BalanceToTokenMainBalance *big.Int

	BalanceOperatorFeeBalance *big.Int
	NumConditionalIncrement   int
}

type State struct {
	BlockNumber     int
	NextFreeId      int
	AccountIdByAddr map[common.Address]int
	Accounts        map[int]accounts.Account
	AccountTree     smt.SparseQuadMerkleTree
}

func (s *State) RootHash() fr.Repr {
	return s.AccountTree.RootHash()
}

func (s *State) ExecuteTx(tx transaction.ZionTx, curCond int, operatorId int) {

}
