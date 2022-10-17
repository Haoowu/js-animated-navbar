package bitcoind

// Represents a block
type Block struct {
	// The block hash
	Hash string `json:"hash"`

	// The number of confirmations
	Confirmations uint