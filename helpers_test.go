package bitcoind

import (
	//. "github.com/Toorop/go-bitcoind"
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helpers", func() {
	Describe("Parse errors", func() {
		Context("No error", func() {
			response := rpcResponse{
				Id:     1212,
				Result: []byte("{}"),
				Err:    nil,
			}
			It("should return nil", func() {
				Expect(handleError(n