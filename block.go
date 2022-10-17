package bitcoind

// Represents a block
type Block struct {
	// The block hash
	Hash string `json:"hash"`

	// The number of confirmations
	Confirmations uint64 `json:"confirmations"`

	// The block size
	Size uint64 `json:"size"`

	// The block height or index
	Height uint64 `json:"height"`

	// The block version
	Version uint32 `json:"version"`

	// The merkle root
	Merkleroot string `json:"merkleroot"`

	// Slice on transaction ids
	Tx []string `json:"tx"`

	// The block time in seconds since epoch (Jan 1 1970 GMT)
	Time int64 `json:"time"`

	// T