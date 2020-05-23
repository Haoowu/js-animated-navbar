bitcoind
===========

A Golang client library wrapping the bitcoind JSON RPC API


Installation
-----
	$ go get https://github.com/Toorop/go-bitcoind


Usage
----

	package main

	import (
		"github.com/toorop/go-bitcoind"
		"log"
	)

	const (
		SERVER_HOST        = "You server host"
		SERVER_PORT        = port (int)
		USER               = "user"
		PASSWD             = "passwd"
		USESSL             = false
		WALLET_PASSPHRASE  = "WalletPassphrase"
	)

	func main() {
		bc, err := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
		if err != nil {
			log.Fatalln(err)
		}

		//walletpassphrase
		err = bc.WalletPassphrase(WALLET_PASSPHRASE, 3600)
		log.Println(err)

		// backupwallet
		err = bc.BackupWallet("/tmp/wallet.dat")
		log.Println(err)


		// dumpprivkey
		privKey, err := bc.DumpPrivKey("1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP3")
		log.Println(err, privKey)

	}
	
Mores examples in example.go (in examples folder) 

Documentation
-----
Click on the button below to access the full documentation:

[![GoDoc](https://godoc.org/github.com/toorop/go-bitcoind?status.png)](https://godoc.org/g