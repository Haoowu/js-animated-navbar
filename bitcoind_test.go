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

var _ = Describe("Bitcoind", func() {
	// We normaly just have to test calls that return data + err
	// server error handling is already tested in helpers_tests
	// But for the fisrt test we will do it as sample

	Describe("backupwallet", func() {
		// Will be used to test all case where only error could be returned
		Context("when success", func() {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `{"result":null,"error":null,"id":1400432805294160077}`)
			})
			ts, host, port, err := getNewTestServer(handler)
			if err != nil {
				log.Fatalln(err)
			}
			defer ts.Close()
			bitcoindClient, _ := New(host, port, "x", "fake", false)
			err = bitcoindClient.BackupWallet("/tmp/wallet.dat")
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
		// will be used to test all server error handling (ie when server reply whith error!=nil)
		Context("when error from server", func() {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `{"result":null,"error":{"code":6,"message":"fake error"},"id":1400425780999713481}`)
			})
			ts, host, port, err := getNewTestServer(handler)
			if err != nil {
				log.Fatalln(err)
			}
			defer ts.Close()
			bitcoindClient, _ := New(host, port, "x", "fake", false)
			err = bitcoindClient.BackupWallet("/tmp/wallet.dat")
			It("error should occured", func() {
				Expect(err).To(HaveOccurred())
			})

			It("error should be 'fake error'", func() {
				Expect(err).Should(MatchError("(6) fake error"))
			})
		})
	})

	Describe("dumpprivkey", func() {
		Context("when success", func() {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `{"result":"K7boEpon3igLpbVv6xebaR4bHALHPeLQSHhUJGiZ9S1U85ERWWi9","error":null,"id":1400433741655216321}`)
			})
			ts, host, port, err := getNewTestServer(handler)
			if err != nil {
				log.Fatalln(err)
			}
			defer ts.Close()
			bitcoindClient, _ := New(host, port, "x", "fake", false)
			privKey, err := bitcoindClient.DumpPrivKey("1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP3")
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should be a the pk", func() {
				Expect(privKey).To(Equal("K7boEpon3igLpbVv6xebaR4bHALHPeLQSHhUJGiZ9S1U85ERWWi9"))
			})
		})
	})

	/*Describe("encryptwallet", func() {
		Context("when success", func() {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `{"result":null,"error":null,"id":1400432805294160077}`)
			})
			ts, host, port, err := getNewTestServer(handler)
			if err != nil {
				log.Fatalln(err)
			}
			defer ts.Close()
			bitcoindClient, _ := New(host, port, "x", "fake", false)
			err = bitcoindClient.EncryptWallet("fakePasswd")
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})*/

	Describe("Testing GetAccount", func() {
		Context("when success", func() {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `{"result":"testAccount","error":null,"id":1400477642632278723}`)
			})
			ts, host, port, err := getNewTestServer(handler)
			if err != nil {
				log.Fatalln(err)
			}
			defer ts.Close()
			bitcoindClient, _ := New(host, port, "x", "fake", false)
			account, err := bitcoindClient.GetAccount("1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP2")
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should return string testAccount", func() {
				Expect(account).To(Equal("testAccount"))
			})
		})
	})

	Describe("Testing GetAccountAddress", func() {
		Context("when success", func() {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `{"result":"1Pyizp4HK7Bfz7CdbSwHHtprk7Ghumhxmy","error":null,"id":1400480276786253434}`)
			})
			ts, host, port, err := getNewTestServer(handler)
			if err != nil {
				log.Fatalln(err)
			}
			defer ts.Close()
			bitcoindClient, _ := New(host, port, "x", "fake", false)
			account, err := bitcoindClient.GetAccountAddress("testAccount")
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should return string 1Pyizp4HK7Bfz7CdbSwHHtprk7Ghumhxmy", func() {
				Expect(account).To(Equal("1Pyizp4HK7Bfz7CdbSwHHtprk7Ghumhxmy"))
			})
		})
	})

	Describe("Testing GetAddressesByAccount", func() {
		Context("when success", func() {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `{"result":["1Pyizp4HK7Bfz7CdbSwHHtprk7Ghumhxmy","1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP3","164s6WasTY9DruJRKq9SHdRjyj3KTw12aS","1obwJCPP9WvqJEG5QgGM97biLRkcwR55m"],"error":null,"id":1400480362380428320}`)
			})
			ts, host, port, err := getNewTestServer(handler)
			if err != nil {
				log.Fatalln(err)
			}
			defer ts.Close()
			bitcoindClient, _ := New(host, port, "x", "fake", false)
			addresses, err := bitcoindClient.GetAddressesByAccount("testAccount")
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should be a slice of string", func() {
				Expect(addresses).Should(BeAssignableToTypeOf([]string{}))

			})
			It(`should return slice "1Pyizp4HK7Bfz7CdbSwHHtprk7Ghumhxmy","1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP3","164s6WasTY9DruJRKq9SHdRjyj3KTw12aS","1obwJCPP9WvqJEG5QgGM97biLRkcwR55m"`, func() {
				Expect(addresses).To(Equal([]string{"1Pyizp4HK7Bfz7CdbSwHHtprk7Ghumhxmy", "1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP3", "164s6WasTY9DruJRKq9SHdRjyj3KTw12aS", "1obwJCPP9WvqJEG5QgGM97biLRkcwR55m"}))
			})
		})
	})

	Describe("Testing GetBalance", func() {
		Context("when success", func() {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `{"result":0.00066600,"error":null,"id":1400501795897598372}`)
			})
			ts, host, port, err