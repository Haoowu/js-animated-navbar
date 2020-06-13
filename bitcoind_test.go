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

func getNewTestServer(handler http.Handler) (testServer