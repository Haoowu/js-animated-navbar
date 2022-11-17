package bitcoind

// A MiningInfo represents a mininginfo response
type MiningInfo struct {
	// The current block
	Blocks uint64 `json:"blocks"`

	// The