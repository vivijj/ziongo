// Package transaction define the tx in the zion network(tx from l2 directly & priority tx from contract)
package transaction

import (
	"crypto/sha256"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

type TxType string

const (
	Noop         TxType = "Noop"
	Deposit      TxType = "Deposit"
	Withdraw     TxType = "Withdraw"
	Transfer     TxType = "Transfer"
	PubKeyUpdate TxType = "PubKeyUpdate"
)

// ZionTx is the L2 transaction(transfer, pubkey update, withdraw) init from user directly.
type ZionTx interface {
	isZionTx()

	// GetBytes Encode the transaction data as the byte sequence according to zion protocol.
	GetBytes() []byte

	// // AuxData will return the auxiliary data of this specific transaction.
	// AuxData(txIndex int) []byte
}

func ZionTxHash(tx ZionTx) common.Hash {
	txBytes := tx.GetBytes()
	return sha256.Sum256(txBytes)
}

// ZionPriTx is the transaction init from contract(only include deposit now).
type ZionPriTx interface {
	// flag function
	isZionPriTx()
}

// PriorityTx is description with the metadata required for server to process it.
type PriorityTx struct {
	// Hash of corresponding l1 transaction.
	L1Hash common.Hash
	// Block in which L1 transaction was included.
	L1Block int
	// Transaction index in the L1 Block
	L1BlockIndex int
	// Priority transaction
	Data ZionPriTx
}

func PriTxHash(tx PriorityTx) common.Hash {
	return sha256.Sum256(tx.GetBytes())
}

func (ptx PriorityTx) isZionPriTx() {}

// GetBytes will return the bytes that consist of the tx layer info.
func (ptx PriorityTx) GetBytes() (out []byte) {
	out = append(out, ptx.L1Hash.Bytes()...)
	out = append(out, IntToBytes(ptx.L1Block)...)
	out = append(out, IntToBytes(ptx.L1BlockIndex)...)
	return
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func IntToBytes[T Integer](i T) []byte {
	return []byte(strconv.Itoa(int(i)))
}
