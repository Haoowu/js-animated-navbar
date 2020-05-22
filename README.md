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
			log.Fatalln(e