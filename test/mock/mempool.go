package mock

import (
	"github.com/vapor/protocol"
	"github.com/vapor/protocol/bc/types"
)

type Mempool struct {
	txs []*protocol.TxDesc
}

func newMempool() *Mempool {
	return &Mempool{
		txs: []*protocol.TxDesc{},
	}
}

func (m *Mempool) AddTx(tx *types.Tx) {
	m.txs = append(m.txs, &protocol.TxDesc{Tx: tx})
}

func (m *Mempool) GetTransactions() []*protocol.TxDesc {
	return m.txs
}