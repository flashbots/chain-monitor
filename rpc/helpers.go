package rpc

import (
	"encoding/json"
	"errors"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type rpcBlock struct {
	Hash         *ethcommon.Hash        `json:"hash"`
	Transactions []rpcTransaction       `json:"transactions"`
	UncleHashes  []ethcommon.Hash       `json:"uncles"`
	Withdrawals  []*ethtypes.Withdrawal `json:"withdrawals,omitempty"`
}

type rpcTransaction struct {
	tx *ethtypes.Transaction
	txExtraInfo
}

func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}

type txExtraInfo struct {
	BlockNumber *string            `json:"blockNumber,omitempty"`
	BlockHash   *ethcommon.Hash    `json:"blockHash,omitempty"`
	From        *ethcommon.Address `json:"from,omitempty"`
}

type senderFromServer struct {
	addr      ethcommon.Address
	blockhash ethcommon.Hash
}

func (s *senderFromServer) Sender(tx *ethtypes.Transaction) (ethcommon.Address, error) {
	if s.addr == (ethcommon.Address{}) {
		return ethcommon.Address{}, errors.New("sender not cached")
	}
	return s.addr, nil
}

func (s *senderFromServer) ChainID() *big.Int {
	panic("can't sign with senderFromServer")
}

func (s *senderFromServer) Equal(other ethtypes.Signer) bool {
	os, ok := other.(*senderFromServer)
	return ok && os.blockhash == s.blockhash
}

func (s *senderFromServer) Hash(tx *ethtypes.Transaction) ethcommon.Hash {
	panic("can't sign with senderFromServer")
}

func (s *senderFromServer) SignatureValues(tx *ethtypes.Transaction, sig []byte) (R, S, V *big.Int, err error) {
	panic("can't sign with senderFromServer")
}
