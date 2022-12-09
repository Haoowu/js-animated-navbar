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