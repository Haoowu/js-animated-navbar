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
	Height uint64 `jso