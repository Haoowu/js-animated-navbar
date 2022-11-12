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
				Expect(handleError(nil, &response)).To(BeNil())
			})
		})

		// err
		Context("errPrev is not nil", func() {
			errPrev := errors.New("fake error")
			response := rpcResponse{
				Id:     1212,
				Result: []byte("{}"),
				Err:    nil,
			}
			It("should occur", func() {
				err := handleError(errPrev, &response)
				Expect(Ω(err).Should(HaveOccurred()))
			})

		})

		// RPC error is not null
		Context("RPC error",