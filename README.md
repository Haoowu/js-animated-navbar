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
		SERVER_PORT        = p