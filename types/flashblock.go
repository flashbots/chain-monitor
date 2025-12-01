package types

type Flashblock struct {
	PayloadId string `json:"payload_id"`
	Index     int    `json:"index"`

	Metadata FlashblockMetadata `json:"metadata"`

	Diff FlashblockDiff `json:"diff"`
}

type FlashblockMetadata struct {
	BlockNumber uint64 `json:"block_number"`
}

type FlashblockDiff struct {
	BlockHash string `json:"block_hash"`

	StateRoot       string `json:"state_root"`
	ReceiptsRoot    string `json:"receipts_root"`
	WithdrawalsRoot string `json:"withdrawals_root"`
}

func (fb Flashblock) Equal(another Flashblock) bool {
	return fb.PayloadId == another.PayloadId &&
		fb.Index == another.Index &&
		fb.Metadata.Equal(another.Metadata) &&
		fb.Diff.Equal(another.Diff)
}

func (fbm FlashblockMetadata) Equal(another FlashblockMetadata) bool {
	return fbm.BlockNumber == another.BlockNumber
}

func (fbd FlashblockDiff) Equal(another FlashblockDiff) bool {
	return fbd.BlockHash == another.BlockHash &&
		fbd.StateRoot == another.StateRoot &&
		fbd.ReceiptsRoot == another.ReceiptsRoot &&
		fbd.WithdrawalsRoot == another.WithdrawalsRoot
}
