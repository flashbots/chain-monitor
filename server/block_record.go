package server

import (
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type blockRecord struct {
	Number *big.Int       `json:"number"`
	Hash   ethcommon.Hash `json:"hash"`
	Landed bool
}
