// Package block include some zion network block definitions
package block

import (
	"github.com/vivijj/ziongo/types/operation"
	"github.com/vivijj/ziongo/types/transaction"
	"github.com/vivijj/ziongo/types/witness"
)

type ProposedBlock struct {
	Txs    []transaction.ZionTx
	PriTxs []transaction.ZionPriTx
}

// PendingBlock is an intermediate state of the block in the zion network.
// Contains the information about (so far) executed transactions and
// meta-information related to the block creating process.
type PendingBlock struct {
	TimeStamp           int
	ChunksLeft          int
	PendingOpBlockIndex int
	NumConditionalTx    int
	PreviousRootHash    string
	PendingIteration    int
	FailedTxCache       map[string]string
	SuccessOperations   []ExecutedOperation
	FailedTxs           []ExecutedOperation
	Witness             []witness.Witness
	AuxData             []byte
}

type WitnessBlock struct {
	BlockType             int
	BlockNumber           int
	TxWitness             []witness.TxWitness
	MerkleRootBefore      string
	TimeStamp             int
	OperatorAccountId     int
	AccountUpdateOperator witness.AccountUpdateWitness
	MerkleRootAfter       string
	BlockSize             int
}

// Block zion network block
type Block struct {
	BlockNumber int
	// state of chain root hash after execute this block
	NewRootHash string
	// ID of zion network operator
	Operator int
	// List of operation executed in the block(L1 & L2).
	BlockTransactions    []ExecutedOperation
	ProcessedPriTxBefore int
	ProcessedPriTxAfter  int
	BlockSize            int
	TimeStamp            int
}

// ExecutedOperation Representation of executed operations, which can be either L1 or L2.
type ExecutedOperation interface {
	isExecutedOperation()
}

// ExecutedTx Executed L2 transactions.
// some part of this struct should not exist when failed, but
// we don't want to use pointer, so the FailReason and BlockIndex
// is meaningful when success and fail
type ExecutedTx struct {
	Tx      transaction.ZionTx
	Success bool
	// if fail, Op is nil
	Op         operation.ZionOp
	FailReason string
	BlockIndex int
	CreatedAt  int64
}

func (_ ExecutedTx) isExecutedOperation() {}

// ExecutedPriorityTx L1 priority transactions, can't fail in L2.
type ExecutedPriorityTx struct {
	PriTx      transaction.PriorityTx
	Op         operation.ZionOp
	BlockIndex int
	CreatedAt  int64
}

func (_ ExecutedPriorityTx) isExecutedOperation() {}
