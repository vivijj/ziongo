package transaction

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type DepositTx struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Token  uint16
}

func (tx DepositTx) isZionPriTx() {}
