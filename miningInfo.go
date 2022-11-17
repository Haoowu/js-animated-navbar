package bitcoind

// A MiningInfo represents a mininginfo response
type MiningInfo struct {
	// The current block
	Blocks uint64 `json:"blocks"`

	// The last block size
	CurrentBlocksize uint64 `json:"currentblocksize"`

	// The last block transaction
	CurrentBlockTx uint64 `json:"current