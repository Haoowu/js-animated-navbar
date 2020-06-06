
// Package Bitcoind is client librari for bitcoind JSON RPC API
package bitcoind

import (
	"encoding/json"
	"errors"
	"strconv"
)

const (
	// VERSION represents bicoind package version
	VERSION = 0.1
	// DEFAULT_RPCCLIENT_TIMEOUT represent http timeout for rcp client
	RPCCLIENT_TIMEOUT = 30
)

// A Bitcoind represents a Bitcoind client
type Bitcoind struct {
	client *rpcClient
}

// New return a new bitcoind
func New(host string, port int, user, passwd string, useSSL bool, timeoutParam ...int) (*Bitcoind, error) {
	var timeout int = RPCCLIENT_TIMEOUT
	// If the timeout is specified in timeoutParam, allow it.
	if len(timeoutParam) != 0 {
		timeout = timeoutParam[0]
	}

	rpcClient, err := newClient(host, port, user, passwd, useSSL, timeout)
	if err != nil {
		return nil, err
	}
	return &Bitcoind{rpcClient}, nil
}

// BackupWallet Safely copies wallet.dat to destination,
// which can be a directory or a path with filename on the remote server
func (b *Bitcoind) BackupWallet(destination string) error {
	r, err := b.client.call("backupwallet", []string{destination})
	return handleError(err, &r)
}

// DumpPrivKey return private key as string associated to public <address>
func (b *Bitcoind) DumpPrivKey(address string) (privKey string, err error) {
	r, err := b.client.call("dumpprivkey", []string{address})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &privKey)
	return
}

// EncryptWallet encrypts the wallet with <passphrase>.
func (b *Bitcoind) EncryptWallet(passphrase string) error {
	r, err := b.client.call("encryptwallet", []string{passphrase})
	return handleError(err, &r)
}

// GetAccount returns the account associated with the given address.
func (b *Bitcoind) GetAccount(address string) (account string, err error) {
	r, err := b.client.call("getaccount", []string{address})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &account)
	return
}

// GetAccountAddress Returns the current bitcoin address for receiving
// payments to this account.
// If account does not exist, it will be created along with an
// associated new address that will be returned.
func (b *Bitcoind) GetAccountAddress(account string) (address string, err error) {
	r, err := b.client.call("getaccountaddress", []string{account})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &address)
	return
}

// GetAddressesByAccount return addresses associated with account <account>
func (b *Bitcoind) GetAddressesByAccount(account string) (addresses []string, err error) {
	r, err := b.client.call("getaddressesbyaccount", []string{account})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &addresses)
	return
}

// GetBalance return the balance of the server or of a specific account
//If [account] is "", returns the server's total available balance.
//If [account] is specified, returns the balance in the account
func (b *Bitcoind) GetBalance(account string, minconf uint64) (balance float64, err error) {
	r, err := b.client.call("getbalance", []interface{}{account, minconf})
	if err = handleError(err, &r); err != nil {
		return
	}
	balance, err = strconv.ParseFloat(string(r.Result), 64)
	return
}

type BlockHeader struct {
	Hash              string
	Confirmations     int
	Height            int
	Version           uint32
	VersionHex        string
	Merkleroot        string
	Time              int64
	Mediantime        int64
	Nonce             uint32
	Bits              uint32
	Difficulty        float64
	Chainwork         string
	Txes              int    `json:"nTx"`
	Previousblockhash string `json:"omitempty"`
	Nextblockhash     string `json:"omitempty"`
}

func (b *Bitcoind) GetBlockheader(blockHash string) (*BlockHeader, error) {
	r, err := b.client.call("getblockheader", []string{blockHash})
	if err = handleError(err, &r); err != nil {
		return nil, err
	}

	var blockHeader BlockHeader
	err = json.Unmarshal(r.Result, &blockHeader)

	return &blockHeader, nil
}

// GetBestBlockhash returns the hash of the best (tip) block in the longest block chain.
func (b *Bitcoind) GetBestBlockhash() (bestBlockHash string, err error) {
	r, err := b.client.call("getbestblockhash", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &bestBlockHash)
	return
}

// GetBlock returns information about the block with the given hash.
func (b *Bitcoind) GetBlock(blockHash string) (block Block, err error) {
	r, err := b.client.call("getblock", []string{blockHash})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &block)
	return
}

// GetRawBlock returns information about the block with the given hash.
func (b *Bitcoind) GetRawBlock(blockHash string) (str string, err error) {
	r, err := b.client.call("getblock", []interface{}{blockHash, false})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &str)
	return
}

// GetBlockCount returns the number of blocks in the longest block chain.
func (b *Bitcoind) GetBlockCount() (count uint64, err error) {
	r, err := b.client.call("getblockcount", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	count, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

// GetBlockHash returns hash of block in best-block-chain at <index>
func (b *Bitcoind) GetBlockHash(index uint64) (hash string, err error) {
	r, err := b.client.call("getblockhash", []uint64{index})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &hash)
	return
}

// getBlockTemplateParams reperesent parameters for GetBlockTemplate
type getBlockTemplateParams struct {
	Mode         string   `json:"mode,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
}

// TODO a finir
// GetBlockTemplate Returns data needed to construct a block to work on.
// See BIP_0022 for more info on params.
func (b *Bitcoind) GetBlockTemplate(capabilities []string, mode string) (template string, err error) {
	params := getBlockTemplateParams{
		Mode:         mode,
		Capabilities: capabilities,
	}
	// TODO []interface{}{mode, capa}
	r, err := b.client.call("getblocktemplate", []getBlockTemplateParams{params})
	if err = handleError(err, &r); err != nil {
		return
	}
	return
}

type ChainTip struct {
	// The height of the current tip
	Height int
	// The hash of the highest tip
	Hash string
	// The length of the associated blockchain branch
	BranchLen int
	// The status of the current tip.
	Status string
}

func (b *Bitcoind) GetChainTips() (tips []ChainTip, err error) {
	r, err := b.client.call("getchaintips", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &tips)
	return
}

// GetConnectionCount returns the number of connections to other nodes.
func (b *Bitcoind) GetConnectionCount() (count uint64, err error) {
	r, err := b.client.call("getconnectioncount", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	count, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

// GetDifficulty returns the proof-of-work difficulty as a multiple of
// the minimum difficulty.
func (b *Bitcoind) GetDifficulty() (difficulty float64, err error) {
	r, err := b.client.call("getdifficulty", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	difficulty, err = strconv.ParseFloat(string(r.Result), 64)
	return
}

// GetGenerate returns true or false whether bitcoind is currently generating hashes
func (b *Bitcoind) GetGenerate() (generate bool, err error) {
	r, err := b.client.call("getgenerate", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &generate)
	return
}

// GetHashesPerSec returns a recent hashes per second performance measurement while generating.
func (b *Bitcoind) GetHashesPerSec() (hashpersec float64, err error) {
	r, err := b.client.call("gethashespersec", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &hashpersec)
	return
}

// GetInfo return result of "getinfo" command (Amazing !)
func (b *Bitcoind) GetInfo() (i Info, err error) {
	r, err := b.client.call("getinfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &i)
	return
}

// GetMiningInfo returns an object containing mining-related information
func (b *Bitcoind) GetMiningInfo() (miningInfo MiningInfo, err error) {
	r, err := b.client.call("getmininginfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &miningInfo)
	return
}

// GetNewAddress return a new address for account [account].
func (b *Bitcoind) GetNewAddress(account ...string) (addr string, err error) {
	// 0 or 1 account
	if len(account) > 1 {
		err = errors.New("Bad parameters for GetNewAddress: you can set 0 or 1 account")
		return
	}
	r, err := b.client.call("getnewaddress", account)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &addr)
	return
}

// GetPeerInfo returns data about each connected node
func (b *Bitcoind) GetPeerInfo() (peerInfo []Peer, err error) {
	r, err := b.client.call("getpeerinfo", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &peerInfo)
	return
}

// GetRawChangeAddress Returns a new Bitcoin address, for receiving change.
// This is for use with raw transactions, NOT normal use.
func (b *Bitcoind) GetRawChangeAddress(account ...string) (rawAddress string, err error) {
	// 0 or 1 account
	if len(account) > 1 {
		err = errors.New("Bad parameters for GetRawChangeAddress: you can set 0 or 1 account")
		return
	}
	r, err := b.client.call("getrawchangeaddress", account)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &rawAddress)
	return