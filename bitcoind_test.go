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
			ts, host, port, err := getNewTestServer(handler)
			if err != nil {
				log.Fatalln(err)
			}
			defer ts.Close()
			bitcoindClient, _ := New(host, port, "x", "fake", false)
			balance, err := bitcoindClient.GetBalance("testAccount", 10)
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should return float64 0.000666", func() {
				Expect(balance).Should(BeNumerically("==", 0.000666))
			})
		})
	})

	Describe("Testing GetBestBlockhash", func() {
		Context("when success", func() {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `{"result":"000000000000000056f14ed49ba8bf0bef7c98b5965058cc6ff02ab00fc26d82","error":null,"id":1400502065079564568}`)
			})
			ts, host, port, err := getNewTestServer(handler)
			if err != nil {
				log.Fatalln(err)
			}
			defer ts.Close()
			bitcoindClient, _ := New(host, port, "x", "fake", false)
			bestblockhash, err := bitcoindClient.GetBestBlockhash()
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should return", func() {
				Expect(bestblockhash).Should(Equal("000000000000000056f14ed49ba8bf0bef7c98b5965058cc6ff02ab00fc26d82"))
			})
		})
	})

	Describe("Testing GetBlock", func() {
		Context("when success", func() {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `{"result":{"hash":"00000000000000003f8d1861d035e44d4297c49bd2517dc0a44ad73c7091926c","confirmations":503,"size":348678,"height":301043,"version":2,"merkleroot":"7867189c03d63d73624b2fab05009a270e15a53611a9efb6fac8fa61981229c9","tx":["d5de3d0622f7379d3e5afe67ff42336854e6b15a77e241fc13e77458bde32603","df8d461f53d509f0d0951df7251b5c63aa1c2380ca07ccb39d2cdd39338b2bef","5b97b6f1411db19dd1029505de8a726d4ab61d30edc48a7481d939dc88a7ddf7","3a275aa7d5027c46dad4433846e9370a671babdbebd601cbf23d69efb0108526","2f8eccfcfc1c7926f7764d77462c479baa2ec2a6dfad1c7a9ed9e9a6647018d3","d547ea33120a88cee685ddda2d64a850cb51af27b4f70458bbaf17cf6e8cbb56","40a5c85b47947ddd91b550572e5df6378b9fc349d4324e5af97b2a1ce4243540","64bc8130a85ef8f7a6acfac8bb813a16509957c44ac3e28d38970964a5fe874e","a64f1953958d580f81e30281defca57363a358911a0d39cb74ee909f3f17e29e","f61a76a408f900c76b7b1e4aec0a2d7861a8d540cce65bd60e1501bcc06ac830","a2c017e41b7f486d9e7e297bfd5c7d7399d33384cc97783976bb088cc378629b","2b93294738cb9763169b88e978fc38117967c78baee367222a4dd47e9f268977","0c3588f866a99d85f57cbe7594918d0020e7b0df2cadc485e599339d5eb8700a","3bf7dca0f2619016050fb3c10d2435b593c621bc342bb1d2164fa4c5dcd89c94","405d48c9897a766c99ff30e0a3c5c16ed120b54abd209d4fdda0ca08a89ccb66","fda807f186ac6951b79b5b6ddcd7b53b4ed0905a60672b4c9be3ddebf6d91a82","384ae85037d051dce938e4f6022d31f28891f2f50fbef21715d6c2067fa605e2","f647ebced14545e6903a5d03785f307010ac32c73052de520cfbdd219c38c26d","8e20033fda8f41ddc29d56eee0fe9034a1ba525888334be5684af43cab387129","995d203ea4981f464b4b8d9b6414b77ace7008ddb3e6874c8c407368ec890e88","cf970431a9611512ddeee367d561f67af3c405399d7eaf7f113d44c618033ae1","87ff5a5c5cb5d84ab1575391993b5f023970c60c1eb0ffdf0a6090d5960cfeea","5a05d6e5b34810b6c917e4868c8ffba5fa2a74200d2d185ae6779f077508c4dd","213a9a911d9fb92c9100672bf5fba7e26152d33be4ea11e464b0a7bc41aee8d6","541e6ee6da227ed199353029d776340fe5e00e25a88078d617472dbf8c9a3c1f","b81ad477640bdb59d0b3375eba7b08f9cbd14bcad4ef2cbd2f2adfdf7fccf221","cacaec1e361d63105acb8bc371440805a04931247a80b0a470343f4bbe9f29fe","2fe07baa219c3a050abfd06159f105c848ef03c0e90fb6fdd806573ca33e7ecf","b7af5609280f0115f88852c77a0f3fb47404ab10110bf3819509d87e300a5d8e","5a15e954f02d742f213819faefcb3ece38bc51ce7221334d58314a4253da9703","6e146b18e620795c19140739ccdffc08034e8f64b0afb6379e86f0e7331b0a7d","0ddf89c82c7e926ea478b919e0d8cc880b75052871499328d73107770c4ee013","4e4914867089d6c8b9f5255b1b45641929f5eb55a0d1a5973c37d6aba466fb6c","aca08e36253bce2a6fc2fa9fde65121224d31ed7ad95eb6fe1d48ba23ce481a8","019866fb68bca4745a37da039dc2b151ab8667edbed6c1891dcc7e6b5fa706b5","de0edd41957fd5def4d2655c9293ec1fac704dbdd68a0c8af1adeb53bee7f3f4","cee3fd755a5204cd88d908fed80e18e115ca323ec0631dda03f0f9032a81c892","601eb4ca2a93e163b50a00b143cb6ee84015a46ddb05eae98d662b48dd4bd6eb","b26e34c1674cf2df7cdefa7a9f07b26f3cc1d7b7018e1fa5ad2c95116b6d3594","6645cabe00ea1c8963ee1c2b2f4c5c3b5a4bde7a9052796b9c3435953b1cafaf","687495aacfd7d0a2d8ffbd3b9ba8d4934610e7a27904185155862ca2b6a58d2a","8e6d8b14140cdd9e424bcfaada3cd79328c600f372ab012a915d619fa24e38af","7f51b88c2620f0876bfe98ddd22f173756714fd59d3d6f4ac1b8b53182f2aebc","2e8bb1e35fcebb68dcd4c5a2da7cdb8a2b5277f087284074cba415cd54eef40c","92f240efb96df3ca2655d887bc2c3e8e4cdff8aec1af0a132b1fbd5ce19b3966","36e91789829e62d95c0cee572c4aa415fbcb3b243570962bfa380d99afac18da","3e5dfb0ac7689c091363eff5a50be568273ae17e8faab4e7af9b65c92df079fe","d9b1bb721b5742d3e4cf89ae517000a1a71e5b831fdafb0fa5b7a9254a1503d4","615375c538a3f5c8486542af8f76c9d932edf5b6ca32eb4e549ccdfd560c1afd","79a31dfff869eb5ffd6e276e0850746ab6a694955a5dfb193ddc68100acd20db","5805c37a60a59f41bdb552b832303183e8f3743efa065fbf38d9fdc3eab48035","a67cd0cdd58e1f3fb8630ddad1fa5eb5f2fdf13f805af5871d7a3520764a1e5d","91730b29a3a40b6f336bdf4ec3992557e30ef5e3a26ba286ab12fb3f186ce226","dffff83c16b071d41730463aa48d3d2bd1d2f92bd65b90023c769eebdd7c759a","51cadb44133fe4153e737f4fd5d515aaabe61f5c04e741a9219573dc068436c8","adfad5808a192e7c7f52e9ae8b8af83867ecc6e3f55d8fcaab8070acbb249c6c","602a6c8ce0c4cf0f29f9354afaa341b8d752b57ebd350d1a9e4837c481dd10c3","9946be73684abfc1a5c02e883aa58126335d7baf81c6f60b2bb7521ad409ff8a","deafc3d8c02b0474000fe40e7d1415d43218968ce1fc826394f417b6ddb0cb47","d61fce6b169c86f3b76438c732cb20f3f2228d49975d0dee51be7ce16bfd0afc","3cd5c00d023a4abefd848ab2d119afbc006f503c50228d82b5750a7238f22178","6f3543df7d70f5258a753ee1cfde37eb7ab8376c3324eea3af1c89a7002c817d","1183614b4755cd8d420d561888913d614ee8bf0be59dabd536d5ea076884eed3","ad9ec97f9c453bc93ed06c1a649b41315ef12eda12106fb1f28f48a315e535b5","884a5b44f2e2264824ea0f854007539e8e6bb80c088112c04d66830e11eec27a","157a87cfa4e3c110d77d5d9dea051795a14cd5ab56b40f7a597405b56cad2afd","6798cb129fd201ec62ece5c1497a3b725b3b628af3d7e4628117c92a0db396bf","24d734fd73252a166bbd933e8d276cec6f9b1a274ea67b0354426ffe50931042","b67f8dbb47555628a882e7b8c597991d88df22c40963c42f7cef58cc76efebae","2b8b71ed597ac1d0cab5842f4dd81343816d904325797514fadab5dda16a826b","e61b054cf5af308cd12e3ba845162b77bb45506a7e0374bad121728bed5328a4","d14b3abbcb9f1090541b5c8670179ee1d4109cf85e4e146c0544eca5b1d5f88a","0935c259422d1b571b9033f4af5e8e3c1e8af0177402b8b558d440a0742ffdd7","c45fb655a486f7d04dac75587c79e583e6588e5f846f3a21bc2ceea10f8dd986","c7be31f6e1457a0eea5f01f36985d38cf77e3bc661d7500bda0f05775a92983e","aaf2dd6c6b7c09b32a6bebe12deacb3a9a18eef364dc70354442693ca0ca4d46","093e19e5170c4447cd3f66b8d01c2c931dbd9fa6a4cd182b5004a96a88ebd998","67ad09a80fd215007c6829424dd477ebf65c0565cecfbc243d28c5611336540f","6a37c35653583829621d62e0405fe3a500dcfaa5dbcb20fd319ed09316615334","83e00befdf31a346d2f16659f1209402ee474d87ac80103f3436244198001ca8","f634aacf20b2941f3c50f52e2fd6c13d108a7885aff8a4e9699553de05e3b606","37f690c8b6b06e64c0a4a9c810cc77cee459362252c4ee4d9f37e9833ebf15d4","a209ce625aa54cfb1e80f5752c905d473659048019b29fcb421d26f5f3e5d66f","0260ceb76791f0e9af269a390cae6ce2f4b6f4dc0718b417f963706d1c337f2a","4ff4aa45aadb9df02bcaeccdf654efab9af26d8f76553c6ee1827cbfe6774df2","1383f5630919dec64efcdc812a8ade37c6f80eb053028c1654b490ee0c7f723a","f9bca8a99b11fb32594b451e9d1621cef215f990c90c7d5db3dd31423f0f0fda","cd01d769bc25891480b868624da012b6f1ff7242d345bd5e2c102bc693a11fd8","4818d5e0c1cf4f3efbef25c372496373738acc223a1604bd01f635c0b5d5f057","8c0fffca0c2448cfb29d8e869f83f1414d0c49184d518e4d244098133603b98c","638982e3f3918fad33d82709b7489e24e58e9bacfb7004af0554e8473d26d63e","d41a96f8e54527dd6b52d2a9ca71e91120e7eb2609366f924de4e3947364f28e","e9e9ea91795527a83544e24ccb8cf62095e1fdfceda6413221344351c2aa9594","17f59dda1f66a6cfdc4e02e0263bd77880926564233c5b4414aaf094075b003a","10d9a70a67410b46cfc188e670e729f5723dd913369b687ef374c7b1a7b32d11","b8582e06977cfe201a4a04a14af9dc0970944b5b21a9bc6d9f4d8025df99c544","12fe78d0bc7839ce0967e4fcc016e1f8fbf2beabbe0f45dd27bdafd3f651f2c3","15ff883293d8c889c525321d2fcf5a4a773303539dd2ca03ed7f6c4f48f24a02","4146ce046f387006073b3dd145eb9c27195349ecfdde081b6e096f5456916108","04f09823800c04c0051466ed81c238ef5901bc0341d8c6a1e85e97f6a76ecc78","b8bd3554e854aff12181d9ee889a4a3263abc2e5bb2bfb7317dd8a3ae08342a5","ff70f8c20399ffb39a7a31f537d440bb2fe6a3193c8642bab4c37dca107d7ce9","6c62e421c25500a99c3c402d9b52382305bf24b2da0db301319bd0f81dfc8877","1ab330ab5ea2718b9272db2fef0701389d5342d964ab5a7a1b1b6304a2ccc542","896ffe2e0f118603f9d8d46f3d7467e7970da1a0682e1b958ed23f23777d77a1","3c21522a4e02aaacfee57087a5b465aef64daa48e9561d1ee1a777b6084e9beb","09edbe76333b830bd3bcc4f8a2362e24207c48d56b398869a299678b16c41a5d","03b3a6ae805677e43aef06f5faedd24e68e1517b6917a7a89b46a6febc177742","3b9600f8d69297b8b2e11935dd3dda85be0a95b5b5d3b17b7860a3607790da22","b059aa436d586b05780663b242bc0b8793c7eaffb7a1f469d933a5af453993bb","ba53eb183903c976b9d62becf5c32bf03271c5b8d11c08cdd3a399f73f438a39","d41abf5cb846d27553ca06e845d15c38e96a7f941c5d5e8bb862e99b9b6f361f","268353f81699f70adbd87142291c7e87e2cd654185fd7e8066a9c8290bd55930","e5fe304ef03724d30e228c0165172f26a7c6f537f8f54a2109b9651807da73aa","e1fc8eb32f860474abc6321908abbd30d488d977618b4915e6697839ba6acced","237a44fd9ad6d31c52fc8abf2bfc801ea66bc31076090845cafa59f5a712415e","492b69618f2d830533aa9444ed9fc72d100375a44cb9453aa4fc4168b1c08181","7094fd0392cb9a2e0f7478a66b97b7ad7a5c3db53ea28367063ffe33e3ee82ec","5ac63d82244dc3a808e9b644c25fd9bb11dde4e54b147a622a93ea84dab86d23","2a17667a67f89fba35a861dad4e7b5b305e4dff05a44d31f4e66c68d8774e025","226027c0f5e576d5ad9e7db535e0b28194e6670daa83a72b915212d7d69a62a9","a2ee3f14a4066db7f73a280b98626c8a5111fc4819d2e91ae783e424d9905a49","4ef247b2360b078717011ea49ff40b557ab8ce35c3bbef684b9a3de5c9a258e6","098d98d94a69b44b8ccf9bff103de27088004f094e4a02ff234f74b009a19850","7fb17534fd6e4f396037641e4f0e0f604332bf7b5ba9ce2b71afc55ff687042c","5b6eb8fd664752576a94c12ed8f77cdd4223e9eb283b622dcf2be75b63ad266c","41b56dd6ddab125efce5e310ee5e32f52644773cccd4e5c3380acc81db821bf0","9db6c8d1578d31ae16a64e9a617eff75d1678f921ae7448c1e3a76a3a292f026","8d3e696f03fc58a188da5545c2033755c745e7be0a26a808020840a6f5aa696e","87fdfbcb8cd0a6190b2b9670a7edaad7bf4bcacf594e9ec5a5dfde7ce6693e84","434d4aec031fe81ca06538ef9bd72e9c07ad25350a03f10bb15e7358e678574d","c3c856b14b8832110e6f76322bff46d9dd77f09188eafa6138aa921548097bb2","2d5ac9e1a89369e239bfb80635eace5361d80efe09773863fb33a8656eb11411","08fcaf677d02536f9968439b13617c459d41aeceb8e52c7c327e1467266a38b0","875eb9b3dd65e0fabd1ee859cdf7e900a5cd96b1e5751380ea04721dc8430396","bdb53560db7a6ad7971e261a6a5355c3d41a8d1985f80515b5c31eefe348f3db","d28204aba16aca0e16d64b9b777a97f13eb0b0f6f44cefcf48ece1f6731c3524","e0469d7259bc009597290a438e66a2ae37ef20cfcdb38e14cbf95e5e8858f466","8ace466d3247c62164373415af77fe7c0e2b242fd7f665629f1de217f892c40a","5253e72c2df6f81375b920cb556192ad5e79a55867600b83e4bc3c7c5e864573","a8823728af127622807f402b158279d13837758dee389e4088aa8c052cf4fe44","78d705955d4b7fa26a4f2f86140ad435f7c36e6a045aea7990ba1627e08e8ca6","c8dfe7ee9c7cf4b314332f3482566e4c025eb303a1c0a22e33b8f6055d088cb2","f38ccd7d120fe82ef57d77740de4f0308ffb51f0b57c90a90ba10ad461acf3ca","093372feddc648bcdba4cb8296a50110807544fbbf9b70b599d17b91573fc66a","5e1e765d24a1916b6c28148630e433ea35446e27a4cdc4417e1002895e8bf73d","63701f1399a68ec58d4840986f6bf49be1d79cf5947cdd3aa7873a4d94e32ca4","7dcff0117581424cdccd9bdf82f559f76cde27c392832182d0cb6a86f43267e6","aee9b706790ad0da77cf6838a495a9a554d848865221aea31ac2c243ceda30ec","69f89bf591abb66734812a91c04fe0c5e99ec3517eaf51afbb2cfb674623b26b","8e80ae392f64e38ab559c0ec82df51226a57d54f88d6a34793515b44f4f0f7a5","ed648af2004b5fb93d32e3a8b78ea342ddd58412a89db32607564f9c27b2fb43","d273a93710f261f9e7408ebbbbbb8d719d0632e10d1fa58cb021ae74f60a1333","5f1cb94f30389287988af6634f7cc6c5aea49abec55a69f1d01ec7bce1517322","7a997171503c38cf740c0dcb1dbf7ba6dbf8173f910e975bdf30ddbe5eaa87b9","d2aaaa29ac73df4e71ba0b11a09b002e46d5411f803d30ac10faa6874095d0df","a246da382ec108c573e9c528a067af35e68db3402bef11ba886985d05a420e3c","6bdd826222631e40c812aba429aa182f7fcc2dde405dfe2b7c2cfca2dfd194d6","ca238460619584f10eddec52feae1c45fc2cce2d5b895a7da81e745ea3be8df1","f55ac75f845260149c63e515de857645b9c7807067b4342c1711eec046a7a3e4","8ab68b4a35bc4114ba553666af91dc3d286994b315eed710220e5c34c4a8c103","560eee525e5c499247504f80d1159498539bd6f15159a776fa03e4092e1574ba","a089b3e4388a869084bb6a8c5fccb743b4b4f414f00e489c2c2e116907c1bac2","9737393d3c8a0f60b454d