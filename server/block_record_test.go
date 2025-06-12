package server

import (
	"encoding/json"
	"math/big"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestBlockRecordMarshalling(t *testing.T) {
	b := &blockRecord{
		Number: big.NewInt(42),
		Hash:   ethcommon.Hash(ethcommon.FromHex("0x4242424242424242424242424242424242424242424242424242424242424242")),
		Landed: true,
	}

	bytes, err := json.Marshal(b)
	assert.NoError(t, err)
	assert.Equal(t, `{"number":42,"hash":"0x4242424242424242424242424242424242424242424242424242424242424242","Landed":true}`, string(bytes))

	b2 := &blockRecord{}
	err = json.Unmarshal(bytes, b2)
	assert.NoError(t, err)

	assert.Equal(t, b.Number, b2.Number)
	assert.Equal(t, b.Hash, b2.Hash)
	assert.Equal(t, b.Landed, b2.Landed)
}
