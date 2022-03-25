package core

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/vivijj/ziongo/types/block"
	"github.com/vivijj/ziongo/types/transaction"
	"github.com/vivijj/ziongo/utils/deque"
)

type MempoolState struct {
	TxsQueue    *deque.Deque[transaction.ZionTx]
	PriTxsQueue *deque.Deque[transaction.ZionPriTx]
}

func NewMempool() *MempoolState {
	return &MempoolState{
		TxsQueue:    deque.New[transaction.ZionTx](),
		PriTxsQueue: deque.New[transaction.ZionPriTx](),
	}
}

func (ms *MempoolState) AddTx(tx transaction.ZionTx) common.Hash {
	ms.TxsQueue.PushBack(tx)
	return transaction.ZionTxHash(tx)
}

func (ms MempoolState) AddPriorityTx(tx transaction.PriorityTx) common.Hash {
	ms.PriTxsQueue.PushBack(tx)
	return common.Hash{}
}

func (ms *MempoolState) ProposeNewBlock() block.ProposedBlock {
	numPriTxs := ms.PriTxsQueue.Len()
	priTxs := make([]transaction.ZionPriTx, 0, numPriTxs)
	for i := 0; i < numPriTxs; i++ {
		t := ms.PriTxsQueue.PopFront()
		priTxs = append(priTxs, t)
	}

	numTxs := ms.TxsQueue.Len()
	txs := make([]transaction.ZionTx, 0, numTxs)
	for i := 0; i < numTxs; i++ {
		t := ms.TxsQueue.PopFront()
		txs = append(txs, t)
	}

	return block.ProposedBlock{
		Txs:    txs,
		PriTxs: priTxs,
	}
}

func (ms *MempoolState) Run() {
	fmt.Println("Mempool handler is running.")
	select {}
}
