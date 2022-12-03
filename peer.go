package bitcoind

type Peer struct {
	// The ip address and port of the peer
	Addr string `json:"addr"`

	// Local address
	Addrlocal string `json:"addrlocal"`

	// The services
	Services string `json:"services"`

	// The time in seconds since epoch (Jan 1 1970 GMT) of the 