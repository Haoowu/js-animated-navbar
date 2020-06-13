package bitcoind

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

func getNewTestServer(handler http.Handler) (testServer *httptest.Server, host string, port int, err error) {
	testServer = httptest.NewServer(handler)
	p := strings.Split(testServer.URL, ":")
	host = p[1][2:]
	pport, err := strconv.ParseInt(p[2], 10, 64)
	port = int(pport)
	return
}

var _ = Describe("Bitcoind