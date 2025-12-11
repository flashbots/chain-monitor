package l2

// Implementation based on https://github.com/flashbots/flashtestations/tree/7cc7f68492fe672a823dd2dead649793aac1f216
import (
	"encoding/binary"
	"fmt"

	"golang.org/x/crypto/sha3"
)

const (
	// Raw TDX v4 quote structure constants
	// Raw quote has a 48-byte header before the TD10ReportBody
	HEADER_LENGTH      = 48
	TD_REPORT10_LENGTH = 584
)

// ComputeWorkloadID computes the workload ID from Automata's serialized verifier output
// This corresponds to QuoteParser.parseV4Quote in Solidity
// The workload ID uniquely identifies a TEE workload based on its measurement registers
func ComputeWorkloadID(rawQuote []byte) ([32]byte, error) {
	var workloadID [32]byte

	// Validate quote length
	if len(rawQuote) < HEADER_LENGTH+TD_REPORT10_LENGTH {
		return workloadID, fmt.Errorf("invalid quote length: %d, expected at least %d",
			len(rawQuote), HEADER_LENGTH+TD_REPORT10_LENGTH)
	}

	// Skip the 48-byte header to get to the TD10ReportBody
	reportBody := rawQuote[HEADER_LENGTH:]

	// Extract fields exactly as parseRawReportBody does in Solidity
	// Using hardcoded offsets to match Solidity implementation exactly
	mrTd := reportBody[136 : 136+48]
	rtMr0 := reportBody[328 : 328+48]
	rtMr1 := reportBody[376 : 376+48]
	rtMr2 := reportBody[424 : 424+48]
	rtMr3 := reportBody[472 : 472+48]
	mrConfigId := reportBody[184 : 184+48]

	// Extract xFAM and tdAttributes (8 bytes each)
	// In Solidity, bytes8 is treated as big-endian for bitwise operations
	xfam := binary.BigEndian.Uint64(reportBody[128 : 128+8])
	tdAttributes := binary.BigEndian.Uint64(reportBody[120 : 120+8])

	xfamBytes := make([]byte, 8)
	tdAttributesBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(xfamBytes, xfam)
	binary.BigEndian.PutUint64(tdAttributesBytes, tdAttributes)

	// Concatenate all fields
	var concatenated []byte
	concatenated = append(concatenated, mrTd...)
	concatenated = append(concatenated, rtMr0...)
	concatenated = append(concatenated, rtMr1...)
	concatenated = append(concatenated, rtMr2...)
	concatenated = append(concatenated, rtMr3...)
	concatenated = append(concatenated, mrConfigId...)
	concatenated = append(concatenated, xfamBytes...)
	concatenated = append(concatenated, tdAttributesBytes...)

	// Compute keccak256 hash
	hash := sha3.NewLegacyKeccak256()
	hash.Write(concatenated)
	copy(workloadID[:], hash.Sum(nil))

	return workloadID, nil
}
