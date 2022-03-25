// Package operation define the zion operation -- the tx after being processed
package operation

import (
	"math/big"

	"github.com/vivijj/ziongo/types/transaction"
)

type ZionOp interface {
	isZionOp()
}

type NoopOp struct{}

func (op NoopOp) IsZionOp() {}

type DepositOp struct {
	Tx        transaction.DepositTx
	AccountId int
}

func (op DepositOp) isZionOp() {}

type PubkeyUpdateOp struct {
	Tx            transaction.PubkeyUpdateTx
	ConditionType uint
	MaxFee        *big.Int
}

func (op PubkeyUpdateOp) isZionOp() {}

type TransferOp struct {
	Tx             transaction.TransferTx
	ToId           int
	ConditionType  int
	MaxFee         *big.Int
	PutAddressInDa bool
}

func (op TransferOp) isZionOp() {}

type WithdrawOp struct {
	Tx            transaction.WithdrawTx
	MaxFee        *big.Int
	ConditionType uint
}

func (op WithdrawOp) isZionOp() {}
