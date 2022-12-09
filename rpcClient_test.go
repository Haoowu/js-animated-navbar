package bitcoind

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"
)

var _ = Describe("RpcClient", func() {
	Describe("Initialise a new rpcClient", func() {
		Context("when initialisation succeeded", func() {
			client, err := newClient("127.0.0.1", 8334, "user", "paswd", false, 30)
			It("err should be nil", func() {
				Expect(err).To(BeNil())
			})
			It("return must be a new rpcClient address ", func() {
				Expect(client).Should(BeAssignableToTypeOf(&rpcClient{}))
			})
		})

		Context("when initialisation failed (empty host)", func() {
			client, err := New("", 8334, "user", "paswd", false)
			It("err should occu