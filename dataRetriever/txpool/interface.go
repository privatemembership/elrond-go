package txpool

import (
	"github.com/ElrondNetwork/elrond-go/storage"
	"github.com/ElrondNetwork/elrond-go/storage/txcache"
)

type txCache interface {
	storage.Cacher

	AddTx(tx *txcache.WrappedTransaction) (ok bool, added bool)
	GetByTxHash(txHash []byte) (*txcache.WrappedTransaction, bool)
	RemoveTxByHash(txHash []byte) error
	ImmunizeTxsAgainstEviction(keys [][]byte)
	ForEachTransaction(function txcache.ForEachTransaction)
}
