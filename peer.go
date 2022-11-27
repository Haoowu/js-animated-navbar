package bitcoind

type Peer struct {
	// The ip address and port of the peer
	Addr string `json:"addr"`

	// Local address
	Addrlocal str