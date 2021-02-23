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
				fmt.Fprintln(w, `{"result":{"hash":"00000000000000003f8d1861d035e44d4297c49bd2517dc0a44ad73c7091926c","confirmations":503,"size":348678,"height":301043,"version":2,"merkleroot":"7867189c03d63d73624b2fab05009a270e15a53611a9efb6fac8fa61981229c9","tx":["d5de3d0622f7379d3e5afe67ff42336854e6b15a77e241fc13e77458bde32603","df8d461f53d509f0d0951df7251b5c63aa1c2380ca07ccb39d2cdd39338b2bef","5b97b6f1411db19dd1029505de8a726d4ab61d30edc48a7481d939dc88a7ddf7","3a275aa7d5027c46dad4433846e9370a671babdbebd601cbf23d69efb0108526","2f8eccfcfc1c7926f7764d77462c479baa2ec2a6dfad1c7a9ed9e9a6647018d3","d547ea33120a88cee685ddda2d64a850cb51af27b4f70458bbaf17cf6e8cbb56","40a5c85b47947ddd91b550572e5df6378b9fc349d4324e5af97b2a1ce4243540","64bc8130a85ef8f7a6acfac8bb813a16509957c44ac3e28d38970964a5fe874e","a64f1953958d580f81e30281defca57363a358911a0d39cb74ee909f3f17e29e","f61a76a408f900c76b7b1e4aec0a2d7861a8d540cce65bd60e1501bcc06ac830","a2c017e41b7f486d9e7e297bfd5c7d7399d33384cc97783976bb088cc378629b","2b93294738cb9763169b88e978fc38117967c78baee367222a4dd47e9f268977","0c3588f866a99d85f57cbe7594918d0020e7b0df2cadc485e599339d5eb8700a","3bf7dca0f2619016050fb3c10d2435b593c621bc342bb1d2164fa4c5dcd89c94","405d48c9897a766c99ff30e0a3c5c16ed120b54abd209d4fdda0ca08a89ccb66","fda807f186ac6951b79b5b6ddcd7b53b4ed0905a60672b4c9be3ddebf6d91a82","384ae85037d051dce938e4f6022d31f28891f2f50fbef21715d6c2067fa605e2","f647ebced14545e6903a5d03785f307010ac32c73052de520cfbdd219c38c26d","8e20033fda8f41ddc29d56eee0fe9034a1ba525888334be5684af43cab387129","995d203ea4981f464b4b8d9b6414b77ace7008ddb3e6874c8c407368ec890e88","cf970431a9611512ddeee367d561f67af3c405399d7eaf7f113d44c618033ae1","87ff5a5c5cb5d84ab1575391993b5f023970c60c1eb0ffdf0a6090d5960cfeea","5a05d6e5b34810b6c917e4868c8ffba5fa2a74200d2d185ae6779f077508c4dd","213a9a911d9fb92c9100672bf5fba7e26152d33be4ea11e464b0a7bc41aee8d6","541e6ee6da227ed199353029d776340fe5e00e25a88078d617472dbf8c9a3c1f","b81ad477640bdb59d0b3375eba7b08f9cbd14bcad4ef2cbd2f2adfdf7fccf221","cacaec1e361d63105acb8bc371440805a04931247a80b0a470343f4bbe9f29fe","2fe07baa219c3a050abfd06159f105c848ef03c0e90fb6fdd806573ca33e7ecf","b7af5609280f0115f88852c77a0f3fb47404ab10110bf3819509d87e300a5d8e","5a15e954f02d742f213819faefcb3ece38bc51ce7221334d58314a4253da9703","6e146b18e620795c19140739ccdffc08034e8f64b0afb6379e86f0e7331b0a7d","0ddf89c82c7e926ea478b919e0d8cc880b75052871499328d73107770c4ee013","4e4914867089d6c8b9f5255b1b45641929f5eb55a0d1a5973c37d6aba466fb6c","aca08e36253bce2a6fc2fa9fde65121224d31ed7ad95eb6fe1d48ba23ce481a8","019866fb68bca4745a37da039dc2b151ab8667edbed6c1891dcc7e6b5fa706b5","de0edd41957fd5def4d2655c9293ec1fac704dbdd68a0c8af1adeb53bee7f3f4","cee3fd755a5204cd88d908fed80e18e115ca323ec0631dda03f0f9032a81c892","601eb4ca2a93e163b50a00b143cb6ee84015a46ddb05eae98d662b48dd4bd6eb","b26e34c1674cf2df7cdefa7a9f07b26f3cc1d7b7018e1fa5ad2c95116b6d3594","6645cabe00ea1c8963ee1c2b2f4c5c3b5a4bde7a9052796b9c3435953b1cafaf","687495aacfd7d0a2d8ffbd3b9ba8d4934610e7a27904185155862ca2b6a58d2a","8e6d8b14140cdd9e424bcfaada3cd79328c600f372ab012a915d619fa24e38af","7f51b88c2620f0876bfe98ddd22f173756714fd59d3d6f4ac1b8b53182f2aebc","2e8bb1e35fcebb68dcd4c5a2da7cdb8a2b5277f087284074cba415cd54eef40c","92f240efb96df3ca2655d887bc2c3e8e4cdff8aec1af0a132b1fbd5ce19b3966","36e91789829e62d95c0cee572c4aa415fbcb3b243570962bfa380d99afac18da","3e5dfb0ac7689c091363eff5a50be568273ae17e8faab4e7af9b65c92df079fe","d9b1bb721b5742d3e4cf89ae517000a1a71e5b831fdafb0fa5b7a9254a1503d4","615375c538a3f5c8486542af8f76c9d932edf5b6ca32eb4e549ccdfd560c1afd","79a31dfff869eb5ffd6e276e0850746ab6a694955a5dfb193ddc68100acd20db","5805c37a60a59f41bdb552b832303183e8f3743efa065fbf38d9fdc3eab48035","a67cd0cdd58e1f3fb8630ddad1fa5eb5f2fdf13f805af5871d7a3520764a1e5d","91730b29a3a40b6f336bdf4ec3992557e30ef5e3a26ba286ab12fb3f186ce226","dffff83c16b071d41730463aa48d3d2bd1d2f92bd65b90023c769eebdd7c759a","51cadb44133fe4153e737f4fd5d515aaabe61f5c04e741a9219573dc068436c8","adfad5808a192e7c7f52e9ae8b8af83867ecc6e3f55d8fcaab8070acbb249c6c","602a6c8ce0c4cf0f29f9354afaa341b8d752b57ebd350d1a9e4837c481dd10c3","9946be73684abfc1a5c02e883aa58126335d7baf81c6f60b2bb7521ad409ff8a","deafc3d8c02b0474000fe40e7d1415d43218968ce1fc826394f417b6ddb0cb47","d61fce6b169c86f3b76438c732cb20f3f2228d49975d0dee51be7ce16bfd0afc","3cd5c00d023a4abefd848ab2d119afbc006f503c50228d82b5750a7238f22178","6f3543df7d70f5258a753ee1cfde37eb7ab8376c3324eea3af1c89a7002c817d","1183614b4755cd8d420d561888913d614ee8bf0be59dabd536d5ea076884eed3","ad9ec97f9c453bc93ed06c1a649b41315ef12eda12106fb1f28f48a315e535b5","884a5b44f2e2264824ea0f854007539e8e6bb80c088112c04d66830e11eec27a","157a87cfa4e3c110d77d5d9dea051795a14cd5ab56b40f7a597405b56cad2afd","6798cb129fd201ec62ece5c1497a3b725b3b628af3d7e4628117c92a0db396bf","24d734fd73252a166bbd933e8d276cec6f9b1a274ea67b0354426ffe50931042","b67f8dbb47555628a882e7b8c597991d88df22c40963c42f7cef58cc76efebae","2b8b71ed597ac1d0cab5842f4dd81343816d904325797514fadab5dda16a826b","e61b054cf5af308cd12e3ba845162b77bb45506a7e0374bad121728bed5328a4","d14b3abbcb9f1090541b5c8670179ee1d4109cf85e4e146c0544eca5b1d5f88a","0935c259422d1b571b9033f4af5e8e3c1e8af0177402b8b558d440a0742ffdd7","c45fb655a486f7d04dac75587c79e583e6588e5f846f3a21bc2ceea10f8dd986","c7be31f6e1457a0eea5f01f36985d38cf77e3bc661d7500bda0f05775a92983e","aaf2dd6c6b7c09b32a6bebe12deacb3a9a18eef364dc70354442693ca0ca4d46","093e19e5170c4447cd3f66b8d01c2c931dbd9fa6a4cd182b5004a96a88ebd998","67ad09a80fd215007c6829424dd477ebf65c0565cecfbc243d28c5611336540f","6a37c35653583829621d62e0405fe3a500dcfaa5dbcb20fd319ed09316615334","83e00befdf31a346d2f16659f1209402ee474d87ac80103f3436244198001ca8","f634aacf20b2941f3c50f52e2fd6c13d108a7885aff8a4e9699553de05e3b606","37f690c8b6b06e64c0a4a9c810cc77cee459362252c4ee4d9f37e9833ebf15d4","a209ce625aa54cfb1e80f5752c905d473659048019b29fcb421d26f5f3e5d66f","0260ceb76791f0e9af269a390cae6ce2f4b6f4dc0718b417f963706d1c337f2a","4ff4aa45aadb9df02bcaeccdf654efab9af26d8f76553c6ee1827cbfe6774df2","1383f5630919dec64efcdc812a8ade37c6f80eb053028c1654b490ee0c7f723a","f9bca8a99b11fb32594b451e9d1621cef215f990c90c7d5db3dd31423f0f0fda","cd01d769bc25891480b868624da012b6f1ff7242d345bd5e2c102bc693a11fd8","4818d5e0c1cf4f3efbef25c372496373738acc223a1604bd01f635c0b5d5f057","8c0fffca0c2448cfb29d8e869f83f1414d0c49184d518e4d244098133603b98c","638982e3f3918fad33d82709b7489e24e58e9bacfb7004af0554e8473d26d63e","d41a96f8e54527dd6b52d2a9ca71e91120e7eb2609366f924de4e3947364f28e","e9e9ea91795527a83544e24ccb8cf62095e1fdfceda6413221344351c2aa9594","17f59dda1f66a6cfdc4e02e0263bd77880926564233c5b4414aaf094075b003a","10d9a70a67410b46cfc188e670e729f5723dd913369b687ef374c7b1a7b32d11","b8582e06977cfe201a4a04a14af9dc0970944b5b21a9bc6d9f4d8025df99c544","12fe78d0bc7839ce0967e4fcc016e1f8fbf2beabbe0f45dd27bdafd3f651f2c3","15ff883293d8c889c525321d2fcf5a4a773303539dd2ca03ed7f6c4f48f24a02","4146ce046f387006073b3dd145eb9c27195349ecfdde081b6e096f5456916108","04f09823800c04c0051466ed81c238ef5901bc0341d8c6a1e85e97f6a76ecc78","b8bd3554e854aff12181d9ee889a4a3263abc2e5bb2bfb7317dd8a3ae08342a5","ff70f8c20399ffb39a7a31f537d440bb2fe6a3193c8642bab4c37dca107d7ce9","6c62e421c25500a99c3c402d9b52382305bf24b2da0db301319bd0f81dfc8877","1ab330ab5ea2718b9272db2fef0701389d5342d964ab5a7a1b1b6304a2ccc542","896ffe2e0f118603f9d8d46f3d7467e7970da1a0682e1b958ed23f23777d77a1","3c21522a4e02aaacfee57087a5b465aef64daa48e9561d1ee1a777b6084e9beb","09edbe76333b830bd3bcc4f8a2362e24207c48d56b398869a299678b16c41a5d","03b3a6ae805677e43aef06f5faedd24e68e1517b6917a7a89b46a6febc177742","3b9600f8d69297b8b2e11935dd3dda85be0a95b5b5d3b17b7860a3607790da22","b059aa436d586b05780663b242bc0b8793c7eaffb7a1f469d933a5af453993bb","ba53eb183903c976b9d62becf5c32bf03271c5b8d11c08cdd3a399f73f438a39","d41abf5cb846d27553ca06e845d15c38e96a7f941c5d5e8bb862e99b9b6f361f","268353f81699f70adbd87142291c7e87e2cd654185fd7e8066a9c8290bd55930","e5fe304ef03724d30e228c0165172f26a7c6f537f8f54a2109b9651807da73aa","e1fc8eb32f860474abc6321908abbd30d488d977618b4915e6697839ba6acced","237a44fd9ad6d31c52fc8abf2bfc801ea66bc31076090845cafa59f5a712415e","492b69618f2d830533aa9444ed9fc72d100375a44cb9453aa4fc4168b1c08181","7094fd0392cb9a2e0f7478a66b97b7ad7a5c3db53ea28367063ffe33e3ee82ec","5ac63d82244dc3a808e9b644c25fd9bb11dde4e54b147a622a93ea84dab86d23","2a17667a67f89fba35a861dad4e7b5b305e4dff05a44d31f4e66c68d8774e025","226027c0f5e576d5ad9e7db535e0b28194e6670daa83a72b915212d7d69a62a9","a2ee3f14a4066db7f73a280b98626c8a5111fc4819d2e91ae783e424d9905a49","4ef247b2360b078717011ea49ff40b557ab8ce35c3bbef684b9a3de5c9a258e6","098d98d94a69b44b8ccf9bff103de27088004f094e4a02ff234f74b009a19850","7fb17534fd6e4f396037641e4f0e0f604332bf7b5ba9ce2b71afc55ff687042c","5b6eb8fd664752576a94c12ed8f77cdd4223e9eb283b622dcf2be75b63ad266c","41b56dd6ddab125efce5e310ee5e32f52644773cccd4e5c3380acc81db821bf0","9db6c8d1578d31ae16a64e9a617eff75d1678f921ae7448c1e3a76a3a292f026","8d3e696f03fc58a188da5545c2033755c745e7be0a26a808020840a6f5aa696e","87fdfbcb8cd0a6190b2b9670a7edaad7bf4bcacf594e9ec5a5dfde7ce6693e84","434d4aec031fe81ca06538ef9bd72e9c07ad25350a03f10bb15e7358e678574d","c3c856b14b8832110e6f76322bff46d9dd77f09188eafa6138aa921548097bb2","2d5ac9e1a89369e239bfb80635eace5361d80efe09773863fb33a8656eb11411","08fcaf677d02536f9968439b13617c459d41aeceb8e52c7c327e1467266a38b0","875eb9b3dd65e0fabd1ee859cdf7e900a5cd96b1e5751380ea04721dc8430396","bdb53560db7a6ad7971e261a6a5355c3d41a8d1985f80515b5c31eefe348f3db","d28204aba16aca0e16d64b9b777a97f13eb0b0f6f44cefcf48ece1f6731c3524","e0469d7259bc009597290a438e66a2ae37ef20cfcdb38e14cbf95e5e8858f466","8ace466d3247c62164373415af77fe7c0e2b242fd7f665629f1de217f892c40a","5253e72c2df6f81375b920cb556192ad5e79a55867600b83e4bc3c7c5e864573","a8823728af127622807f402b158279d13837758dee389e4088aa8c052cf4fe44","78d705955d4b7fa26a4f2f86140ad435f7c36e6a045aea7990ba1627e08e8ca6","c8dfe7ee9c7cf4b314332f3482566e4c025eb303a1c0a22e33b8f6055d088cb2","f38ccd7d120fe82ef57d77740de4f0308ffb51f0b57c90a90ba10ad461acf3ca","093372feddc648bcdba4cb8296a50110807544fbbf9b70b599d17b91573fc66a","5e1e765d24a1916b6c28148630e433ea35446e27a4cdc4417e1002895e8bf73d","63701f1399a68ec58d4840986f6bf49be1d79cf5947cdd3aa7873a4d94e32ca4","7dcff0117581424cdccd9bdf82f559f76cde27c392832182d0cb6a86f43267e6","aee9b706790ad0da77cf6838a495a9a554d848865221aea31ac2c243ceda30ec","69f89bf591abb66734812a91c04fe0c5e99ec3517eaf51afbb2cfb674623b26b","8e80ae392f64e38ab559c0ec82df51226a57d54f88d6a34793515b44f4f0f7a5","ed648af2004b5fb93d32e3a8b78ea342ddd58412a89db32607564f9c27b2fb43","d273a93710f261f9e7408ebbbbbb8d719d0632e10d1fa58cb021ae74f60a1333","5f1cb94f30389287988af6634f7cc6c5aea49abec55a69f1d01ec7bce1517322","7a997171503c38cf740c0dcb1dbf7ba6dbf8173f910e975bdf30ddbe5eaa87b9","d2aaaa29ac73df4e71ba0b11a09b002e46d5411f803d30ac10faa6874095d0df","a246da382ec108c573e9c528a067af35e68db3402bef11ba886985d05a420e3c","6bdd826222631e40c812aba429aa182f7fcc2dde405dfe2b7c2cfca2dfd194d6","ca238460619584f10eddec52feae1c45fc2cce2d5b895a7da81e745ea3be8df1","f55ac75f845260149c63e515de857645b9c7807067b4342c1711eec046a7a3e4","8ab68b4a35bc4114ba553666af91dc3d286994b315eed710220e5c34c4a8c103","560eee525e5c499247504f80d1159498539bd6f15159a776fa03e4092e1574ba","a089b3e4388a869084bb6a8c5fccb743b4b4f414f00e489c2c2e116907c1bac2","9737393d3c8a0f60b454d709886718b939c78c1228a4081f7403cea5e5a66976","5c2ef0ecd978d998db5a2cdb5417ac8debb3d9be6c52412f4da6c1defa206603","aa7781a1d45679bf7e297242e183fe0d7e1b19da254cc339d2864ddd9df51845","bbee2db2e27a002da9b94755fd2825bc24e5ec3f044807040a432994a724c188","c056a94aaf7187e7ccffdd13ff409f9cf70e5b0cb7f5c89e3af847e7ea4befd1","52212ddd3f83dd58299fe51ef2c5c21abc1a982b4e8cb1d6dc1de932d712046e","2c080c35a11b1265fe064743f8ad2179f121d5bc8ac4dfacaabdd30083423f79","fa7d4ade2d6556400f0f5244d2e95e8ae45c9dbdfabbf7bf87d3ef9900a57be0","4cf4c95aad2b366875b90f091a1a88f9a747d6e5c9802d4b837b9e8f8f87874a","ac8bdf0bd4c0a396dbdad36ade72d1f7f4a7cd81ccc1b21392beee40c3088b72","76f299dac7977d6747cf94677b4ee874dfd92ae4d306f187bd005e422f2a3d7d","67270aec165cbf2b5c473e6b4ee64fe0e4656afce5588f6dbf978c6e647ac127","690c97a04465ab31eede64205b9a8bf32fd7972c34c8df83aea1e784eb6cac93","42268fa2e41033057793b8c57ad966e9e070eac19da8fc9be612c12be18c9f8e","848f582468c00d77be3556b8c2ec2334c0350d0c29157652dbe3a4d271bb1c9d","00ed4183815df18f7478e5a78c54826720e4f89c1249cb85445001d4593350f4","74eaccc8510676d8efb1075895b4c1681845851f8f7eaefe4365600a8db6f8ad","fbf9529bb60edea9503640ac4d8cd4fa629abd05b8477b62819a308059f765f2","ba068944f6a7ef42bae16c6c198ef22b19724fe803c783482c624ada3e7888ef","edc982ef71918361f0f03a299dc30c2c0d16eff5d928197e7546dbb2f0afc1ce","bd250d2e3d79dc4fd129c53872ff06eb23f4e3efca7b50b34f89ec0dc34b8165","03285cb37239e1f94b3cc517ffeb919cbb1fd66fb8bd146a3703bcfc19ae7b35","7753494520876fbdeaa20479e2556c2b3cb1d68fae199ad2879f4ef42d27b3a5","1f965dc7822e8aa821aeb9026363a6ead5c5377f1d1f6ec29b5d41b885512dcf","27834d7b9454e0824b9b1b918f7dbf8d11298f8d8d0198d0112ee575bd19d065","151163055af9685fa02861b61a2c11346701699c510364b5028e608092328da1","bd8bd405f16f2f5fe7106669546c4a04135fa3b8af4e8a1ea39b1555e9465ff6","c66d567a2a0bdbdebcbae4b2daa778e40074ff659862e6459c170e2b84a401bf","1f7f01297756ea81114b2d9c587881a3f179b6100ae86456e89265677168917f","2bc40a6658b58aba19d21dd016daf133342a3ba33104e35a7da2cac382563d14","7eb73864a0be8bbeaf513d3ec0f89cf39514ee4f1ffc4f0d86eb13c723078768","fe1ca0a1e59f666ffb5eaf05919434277cc848a59f7891bcd93182b20d604cf1","bb8a2ae0bca832af908506b70dfe5f409a7bca331c8f35723d4d35c4c8491e58","f794dc0f0a9d5c0f9a568d62b31645df78d4e0b0468e92cd5029bef77e4ef24b","d30b9197979dfb3b2a298530e3bfa3f36ce1d989e27318cf17b28e0074e98f90","c09549fe3860fdde04c7d9aa789ee2ae4e468505fa08940d0c3788de4df98f95","34ea4ec3f2d0b49d8ebab0cd20f74e8c350a1a4bfc6cf5d8c7fa212490148f7a","74cc5748a0cb0e776985a65a805461b7b3337fff94ba03d01f1242419eec2d6d","e24a57f7b60889a83e747180d4d6baa0c8ff139e774ef3b09e26f3346f774810","f1b0ded0b2b073507f380fd701c0a73b90616b1ff67bff52e1846ed452d86ad6","7a30e00233b5df9eb0dfcf7f2d9505b1f78ffa31536534f799de6a0fcdf67ab7","3bd7f9f4b4c4dd8334d65ce75f1044718b96ecdb0909a583dff8b2887364ed2c","8ae3d8d754ed179c2d54cdf7b1798eb1f75be681b43148e11adf8c4faf00d492","b22d4b521b9f2bd2d293136acc8a4707ee3a8bab4ae8147d613dd0972a2ff007","491e14be8a3031e6e9fa6e6bf9370a5a66382394e2bb0cd28f52001174520c60","25c724631113f0b6e18094d3d6ca94925e958969dffadaa9d171221acc7174f2","ffcc88e99d961d4881de2ba7b63fde91d6671e409d78820b6bb0425dce0790a0","cdac405f2d658bf5d5a3d939a3fc70ddcb10f99f23294da79a9dab8882769a47","1586037b9b597d46ace36c8d81a488fb66bde4f24a0605970d10fdfd2110e8bf","9196bb0770e882236ba9c52d6c7f86ed94a4e34464b14228b42712dec8bd1907","413b83ded913901b81c6252007460085c8e7984ea88992a988865a58a49c22bd","e5f0453dab1e89bc7e04db72a9595d2d847f5d54dbe57ab009039692c32f98dd","fb1cd5bffe86a629c4e0a8c5861bd5f80be90514d0aa73f7d9169bd6ee45f737","cdbe1d8f4a58ec637724db20157d1d1912e29d18a31c588ea8d212c3ae7eeea0","a64905fcc628a28573454ab6e9237c07e51aa955517503637a5a43b1c9d13b6d","561237db74a3f26bb846268426e6a4b911543967379bd6bbad78185bf86502c0","8a3961417c0e48e69e831e1f9dfc0c2eeda8af97cc385ac7ac210b4e2ae11c6a","a8016aea8c3164c38fc142e07df3c6451350bbee4f921be25fd7297f1f79e049","372937b38ecf0c674556b3ed1c1c5e3aa5710ee4f9348c63fa585e056d23f149","98647e0c001d822bdfe9dd4cae460f3b1bf3ea6e51cd1761089f069e4c4e101a","49cb525d2849807b6ab3550b58655647ef23647625acba9bdda6cf640d5e9632","667a382ea1b278c13eeb34d501591b63f45d702e530ea53e194dd0a7ab6a0ec0","5460b24dc95090e7332ea02a2872639dcf7e202df022792340e447298dfe37af","cba5c9fefc5e6ac19c5d0ddd07f2487c64e57a52be99f24f10dc75f8cf50fabd","2d14f87637804ea79b9a18d66b5cbba88cf41f9b5f03957f82ee8e5b80fb2310","38cd9761c681ed0d6d827c829ede86659656cd4070f21f6b272e371fbd7bad50","601ab673195dda28e8a5490b4c2a12110d0105892fa8bc90661c6e55bdb21c6c","08da9cd996186e04fab1608bb416920232fc9f9c64f8bb932386c63bfbda2cf3","8ecbae2f777e27549ba41209ddcc9733dd06f57633f2e0ae6c6fef17ad0c9152","9e032b3208437271557e48a09efc61523de94be08b2fd3e8123da7ec54802805","7d5838e8e73016c9fd7e92b631a5db82729506deb2f39b112793f0fa6d2ac5f3","c8492cbb3bedb03bfd29b2b2bbfc0470518f0170e9a8671468090d26e5c0edb8","029e8c2aebb479ff74156db891eddcf81e10a931e02263e27eb5aa0130d7aa26","4ac085095aab20d00e27252f7fb44bb8dafe4ae971dff42fec209ae90d4f5ecd","1a2f9a6e89c473179b73b20b43feba63cb49ce15fd9dec98e0aca01e4321d69e","d8dccf53e42872a8d6319fdf1aeac69f15527c9c56600f5ad8d0f82eb74e9004","dce0180e1884d75ef98a0ee7ea93cf9d7c6230684eb72b9f2e38e3d468357186","ac30e79d3cd9808b600642ed78c660029504d81648bc118f4e4a182147ca56cc","912f8c6f3c4d3c25acc524079b8b8736d57b23925e0cbf549de09fe5df908764","6b4c7b2bd545b880d0c73a7bc29aab62959665b3b48bb657d410a92481023a0c","15ffc033369e87f3e5ee6c0737a8ee452a584250457656466ca47156946c0baa","a62237fcbbc5e21bc912b6813cdede4b1f2c203e09f9ae94608200f586a3aa27","9ddc785a43dc7c4f7c7f027768fc6149507d0c0fcb1b9b955da6c994e0cb15bf","1e171d7d3e4d4fe37382dbcadf6319a73755327e7fb7ae1c98e9ae90ae51574d","64e9f7b175d141bdd1d3135be171f6d7a111b3b7bcc59ff6d4123eb6620d9ee5","d788d372bc4a1ee9bf9a66e2744aebc19d2d0cadcb2463aac419c000f64d8ac3","2d5289ff0d9032605341739ad95a5cb6f65001b7b793e6136385ca3858228064","569e8652618ec513ed2e14d146af376f8ad16c0d97cc6e564f21fc8e34f871ad","73f04c34f60531cd3ca1c411a21e6e9f54a1df0fa68312afa7820b74510f9d69","523efb064228674f0f6b591d84ffb5f6a12f3b0258dc5fff79ec5c80fac98b7d","e1910cce097956c57ea2c90afdaec1a0e2788b5204450462824d8cbe448473fb","6dc38ce66cf047003013048e001fa7e17658a505f570c686679198e361d35e25","5f62bfccd33d0c40227792c24fe0dc7bdc96e1bd3405149b2386be5b3950a599","1f4795fec88778fb2d1c8b178f78e8d99d81bcb333c44dd740bad7e1750b5eae","e063c714a6ac6cddc4207f909fe11cc2a5c1759f310d20d4d75a1c80b6d18171","21b073d33a9881be48bfe1fd6cfbebaff9389adcbf8300bcdbf5d4dfeb50239c","5f52e46c70fa573c8004f72da23e3109c15c680c9961a022f4601b7a39ec4d29","0e4ea305c4b1ae872cd82e249be01ef2bb0a4f7653b273e8b2b93b7843fd7b4b","cad1db70904a50dd0bd7d0356097c0f11882a12e2adff93c99430e50b526d8bc","d09897cb69bf2d6ca73abff271e06aa8536bc875984f7c77717965be2facac52","e1da6ef2e3da4de8035f7b1ac1f353f22e6f699e677e9222198992a3aa2126b8","43ca0398dece62e9dbc6a7723a5a3cfd75953feb7fe8dbe03b4ae4eae7039af9","38b9b977f0661534aca2a8f1d98692664683814b6b2f58cee3ac207b1c9c2b43","603aabfb3f5c31341d4122bd709219039194c00b0238105e0b514e84222f919e","2e9eb0786eae1227e2106cd939d816193a41383faad282ae6d09de3bc0bab15d","eb4d06527e60de01ba86520a1fed7970b58656e42517c44cc389a877970be54d","f600e1c1fbcfc667a4b5024d1becb20772c4a2c239fe5ae1aa60bd5b0479d6c6","1689a94c088e2a7718c96467fd8bb9699809b363a13b94fda9a2ae019bbabb93","53626b07f6698829bae790455626d1f2cc4716560713368b2986b070c87f6473","961bf292a28f2a58416234404fb2cbc4b54541a1a842ea82048bf8c9ff0c104f","f8ed404b56bf023163f85ac44c136cca873635cd96cf5e5c00d0fb7eed981d2c","15bf1b9ff7d6b974fadf7b5f091f099db889ebde5d20098eda516523fe8f772a","1dd2f7945428827bb28c130867e7cdc9d70c22dea57c970df7b41b765fd577cb","18b2f0ab2caf97d6efdf7c2a012486248d34b5bd06b6278b1491bd0e20567b02","acd5fde27a117e645b42ee550da1ae0c97af6b9d80137e25b45b78d3c2437331","a44a4ffa98eeeb13e7abf3c4403e9e125ffeca6baa349cf8d9bd96d6a3d0642e","e2999c485699b406c94da10f8f13674f994e0952ba066a7049aefc94d92f288d","5e6d4c4d4e0f5b517c0859fe7a1c1123476498db164f13e37ee2dce87effbf4a","e10ba18665d833663268cd41640b19e8b4b3f4215465fac33cc4c26cb7ac036f","76ac747470626760ca670defc620a10fe830f7360df158e1d420643efef3ca85","887f96bd92a2b6cfbbbaa3bac59df1b171e8277791f61527a8406e51a35614b6","d1054af96d50af4191219b3e7816d323d1d94f0bc7ff862cc79e4ed24cc1613a","3c37cb77a5f7c36d13998533697da58a6d4d4fb329a49d272f14d350e7ac73ee","854d3edb29e477606757f0777d345d97f9a824f33da1ffb36241e2e410261885","b74ed332cc65b07f4c0ffc647084a39ded87ba8235765436f2d9f09e8e63d878","6f30cfb2e57ebef15ce5184f0161927848a67676098e9dde98c2108e4d940dcf","e9a0783416c481de45ebe85f5fcb297763fd41453798872b64c66a302e14c326","b710adb9054f4562c57308c018a4faccfe55648d2eb3f69bf18f3b49b85fee87","bf4deec19d042856f263fceda3164a0c1a72a61bd483f4882163cbf7a783ed7a","bc552378f26031eb3af6725f98d9a94723df6723add71829993ca8ee7cfec132","348610a30c4325df5ae35271267fcb99c55bdab8126d4c51b90bbb4c5126d269","585753b150ba28b1b3e5ccb14f7a057cbd70249212973c248df0851842b041cf","1f33a69b001ed5d9c9014beffdcc6c37cb258447f9c619d9e812a58b8fcd4236","735f9706aa56120b9c2fdfb68bddcddb776ddc050e8efafeb042df922cecea64","cfc1a4aa92e1858f449b576e61c29368cc212bf15e9ec2ea72507163b3b1ff2e","db12fe45f88285ff13341073e0ab4d4c20a8422f5bab7aca7c94eaf1136f1f21","7725d07970d2dccdd5837e435a2cb0f86fa6608fb755ca10f664c2c996a184c9","23f80ecdae97c27314d025b30dc5cbaf22424fa3fd93f39ed8a0cf67f5b855df","8ab7c06989fe239d07d36c869a20228c3de05d544c09ad43f06352ffb873ba89","37c9ea1ae2d50e69fb6d0afb9750c5c276bbf2324f1430a4d48cd178d256035e","337b6af4312e7d1d7836331ce2159f97cf625753d93b4c32fd3d844ad09f196a","fc42505faf48cce3599d1b29f9f0fb155338e41adbe9af09f8b486a9d163140f","82278c4b7e62256b99c7a7552e086f91e9c3a20d6c151229b3a3cf77284c8cbb","ec52a1bf1997df3d246c22debd2dde46385576518f2f6a2ffc74971dac9c1daa","f11152dd7761b567b8518a19350f86eb0e05072e98af21d5adb5b79df9c177e2","49685225341466ac08d3df0b97d0fd8d0367d4ff8aed488a4581e2f31174fef6","17359bbae87cf494d40fc545733bffb2ca0e7210dd76df7b359dc2db56556df0","5d1147dcabcce55241649c936a035b61bb886f58d6e2769fe63517ea71f28368","ad8e3a74a9b11ecc94016898a4bbd380ab7696e18d8378d7898b500e1b52c9b5","21d10afadd5dc01f7da6ec2773b22cb464129e2cce46ece3881065454fdbe08d","df628c46e5a154ce11122c948db82b09ba81d5cfeb7453a0826c8b70efe81ac9","e80d285ace89b1d5f635449e567fd014f9d12ae98830b8ae98aaa0d4faf3928d","75ee9acc2cf70486f9a1327a80092720005910b220dce3eb93883db1847e9ca8","81387b1ffd6ea308c08668c3dc58f57825b84c130ef653d02930e3662e78617a","0a4cd513761eaaab478fa5213bf3432187a3b4ea18ecf9941d91a7307c7c26bd","6ba6cfad9775267d622125f67c99dfcac0f4229f5f68b8b3ba5d797ecc087cee","9263118cb9498ace984dd666dbe1829a3f53b21fabf96cacf9ad434a4df577c1","e151f5e3a74aff1eb61002187a8c14b2b54ea1dd310d0192ceff7882b08f9f3f","bb0f35a52b4e463b8298808a5bd6f256f6d562d6ec1f6e2da85f917774b05099","8cd9b223853ccbae69f088195980f3200a1eca4aaf2ba3debacd70273219e61d","badeda48dea7438d74732d72d9f0e81aa56ea73e2b5fede635041ff120199361","7d1b950a36412f408396c1a72b3ab5ea1e8ab6cf3445a864bc69b2ee93f25f97","da0c1eb2a7d5adaf85062cdd3ee97d54d0e5d2fe38b4c048f1109c3eb550f37b","6ba74e59c8473f054c0d425575f35b4a66d694d888cda137c38a211107e1f4cc","78480714e675e0c1c8534c2a8d56944b379dcc3c791a04f6b7406fcb8bb319cf","bd4e2c34de9e79a016983f450a2cedd89d611c43505e3265eec9742ad86df41d","4e3aae560fb91c2f54d05539d070723d676689edb83bcd35be73c9f16aaf8f54","2a032ecc3b62bb8f09b1564fdbcea71b06d8fb5c3969a03114780643d76f0283","298dcdafa638c43ab4b4f2e26019d06602150560f6987d0ba2bc36b2f028343f","d774101f2593ed5e30ee89fe61e5560e55bfbb0290ae2d7bced8f60feb34a60a","9ed532d20bd200720d41864808c38391fd4a025d520c606f0358d6f4d7ad1351","c5fdb45682e7be82a556fab7e2c4920da0431c3de7d426a95ffcd73dea30ac74","e7679e044e87550bd67eb4abd668e9a59386cd0886cc6c2b931fb890a766b6e3","3a957426ca27a3efbd95ff1ce118d01cff2f0f38d9cc4e8759dfce64a4930a74","bfc497bf23ef2f2b5a25e646454a43d0e5c6694241c9f02cae82e236ca780469","39fc150db910a569d714a608ffececc2dd6c89814f278243f3e474c72d3baada","76df15e13c9059e1ed49061714693ad23c6fa1a64a60f63431e42a858e4f89e4","35434afc19ae78295e8af48c7956c575bec077972f6d522c3f92da1e78cf318e","2442cc287cc612d9ee37db09c9c7b4607ee84cfaa8a936d54eb33d28b21bb7ac","b3b26d1ce87d784acc22a906f33b226714133dbf103f20a59e613d8651e36cf8","6f14900b0cd1b09ee7fafbe74696c58487b6ed46c616d43d9b331101291f5964","c884b07e9136c233c6184b19f3999035af4828121ccac5aa58d3f3ebaa5b0078","a88a0b0d4b38decdbf95b07115cee38946239426ca7d325fa0f9a10841e9c7f3","76bf4d36828416f7c6528fee940dbcd2f6230e82b5435693054d800741c30479","de923b255eda60b394ed166b827a1df9d4f7d2d10cc47dc440c4878502bfe469","6a15b50adb55c20a980674c84f1ed137372eb17d649249ed76c3bca2c99a1508","d0e1f8ac608f5d5cca4e1f08e5fded44f0d702e313f4c8a278bd219a091f68a8","a0ab57340b559fe2bcbdaf2640513bc1dca805736ab8df27385540b6c76e8812","af7dfb403d517b87f8edc3153abba5eddf0f48b5581d7b19f68ea35fc957adaa","500e63bf9f9ae89a2713de9a1e4bd32a04e0e2700501a5fb224fcd2251d2b832","3e5671f5687fc7f4779972f4fbf79f3d9c6e04a1e3a359bbe8f94a6ed1c63e37","3286cbd56a3dddbabeda928e1e3edd9ad1ff511f92b2c6b768feb7dfe75a79cd","3953e5b6f5b91eaf26d27c716d1dd389823ce22aeec5748c47068f15f439ea80","a4db2f178e04c2952f9db4fd40ac1fdf9863d4642dee2281f9ae601ffd03f462","060c2b70be0aace291752f2db27c665439d31f87c83f378c088c877611bc1d4e","891de04f832d67942f9fb5a331502d3adbd7e12ac039bc63c64bcb97589c920d","8682022ee4c936c7e4d5dadfe0cd7c1749c9eae77fcf7fdbdddab605e6ebaf13","af72d166a5f0581bc5a0e43dfb753c52805fe2acadaf96488dfc7b800f4dc44d","4c3bbe400deeb3993bafb138a12c62c23d18950ece7191586a5845ea1ae00347","63fc9d02aa61f97737c3e56c5839048f403d7a09813a00ac0d83ab92dc277742","ffad9e9a87d9cd9fa1818f6fd1fb5bc53f4f6ec8c94ee4377889bd1aaca10c93","090f3a9928409ae95550b44bd89b4b33c5f06ba945c2a7b34b17fcfeb21b1cd7","9a683e6464f3622e554125b6a6127ba37e402057131df8921b76baaff2a8bd5b","7fe6e183d32e9d94e2271a21397c21c013a13f0bb51dee28aab4963a7f7282d2","329016c39799bc916b64fe7192de79a02127daf0d53055d64c39e71dcee373a6","2202967cf8d563500aefd26ed461de2aa643f06655989df749bc0b0ec0fa8d39","b7a5e488c77a949f3e7ee0d14ded5147f20ff260dd084adf589d257fac1bbaa8","1002f1839f92e7b86948700e1d388095d6dcc304a16ee45e548d63e34d16f6c8","3248ab6e1d390b2dfefb85578652a9212d75625a11c2dcf3ce568d08d6918f09","0aa1c91c7b47c5b801b46025e30cff0547b854090f82cb5e0d21ccfebbf275a4","5653b8c0a01de6d6190b36b268f9c86099ba884c64a4ab10da203d44a39c536e","8e389a20a49ddc6d03775d6f5bca3f6372df618c3ce17285e0905e7014fdd9d0","73d71fa3957eeb7a8cfbf79bc458a58a5a584db5452da7a1bb8336c460862955","2c0b1b0785ca0820bc7338b2dcce46662cb0bf3462c14a008d2c9b874b5a4790","42b64a046c7326a156faf10e9c3c1cc873f30c166b35cfaf5eab85c83ed80082","2022b1d326fe10df6c4acc7c55f2d622e5b9f5c3f4c79433b7ce69086fb88fb1","510157ae914a3804d612868924b974b5548d334ebb716ed46e47e16d468c13c7","d06a6a2857b18b4e2ff7ad634347a28d389b1cc93c49bd63c8a39587613ab681","60ef334f5e6609bcfb5b15dc035f494ee4ab0dea393e9723e145601d437d740c","4a539a26cbb8b7792cdf1be23bec78d9f93526a060d3828308b99c8062b24fbb","ea032ba3d60540b4e204b69a9b523a7a6d433a28bf284cbc45b985459b8fe866","c0b54c8e9642d9d73d6fe69392c49ee0d23a4b56877f87c94d9e10c8c0256f14","50204b57d3d49223e51c3e2432fc02cb78a101735c6b4b58499688e17e74ff53","c472d3e720115841a4dbea5e88337220e92fceb5fc3c7125831c65cc1fb0bda4","fe46609fb1159add021de154ae37ef6a01162571873bd5bfdac2c9d9177d95ef","dcb6780b5843f58383300a65009a8a8a1fd95f92bd892a133c4653861ab7d467","bfe52d8a9f24251ad1467c9ff3ba42378b7510b1ca6117f8bbf13d5691d1b87b","27fbcc1eefa63d778cb1c526ba7ce659a98e75f8aadceaf74a377361666957cb","1dad9e4a67b4f504e251a080ddb616a4f85fd35e3b761f074ffb78da50b672a7","4db32a3b09e9afec4f39dd2cf2159cfa7b4995ad673f9471688ddcb871e7dd76","8d701892ab43dfbaea40fa21e0d6d1131a0a2fef28c7f4f1292286f44c68d0a2","735ba988371d7068fba93da5cae532f8f492d6c0471403c7e9af6ebf8b80f36e","15d910ba5691f15939580e492783765910ffb42b21a236817469f49009fcc7cf","1df9c4034d9843d5a908ad75953788fbdf8795f4efe2405fdf5c93188da4f276","93708420b1f937a05157bd742f0dd580b810424f70b8d2d63c49177ba89297d8","4994f6d1de6daf77987b6d1edce7b3e825eb63bf0f4143937e97e1597264f0b5","eed43fb30cf0e85336ea5d4b4641da726704d43e4935e58433d02b011e9b668c","9afd5cdf2645ba8744e7c578713611992d5c2a879030281d4eac837e09ff113e","7cd04211f0db1a80cb7655092528084cbe907b7ace4a23f2db1392e66ecc3e6c","8b540480b2a1fba10ac97b88ac7283b52736e02150d54302b8ba64874693e9dd","013b8745283b7ac7246b38d595690f94e32dd7f94c6d6b48d72adabdf5a94ef2","31e92ee1ceeaf10d9fb7a84567d333b0977fe67febd7e01951932f86d1f61494","0178fc116025caa1f51b632e7d8a98ab8a547748ea3e6a4a1a3dd8095ab6de4e","c100708554838fe740bf321caeb564151b5deaa521a522e0884cd7768a9ff285","943036b3b56ae6955569f724262117219b31bba0fc13bda59934fc14ec891ed9","32097db3ab188f185df590cff0b807ac958861fc1be50ef8219a675e0bac928c","a560d40f4de3a80a076b6fbd070a944bf1a3d523002e108dfc84db644c7d4a59","01c8ce2df11eb393981eba87e1939a49206468ca270da4d1ad7a009b29ad0d66","9268fe46950c7d69839e84388efc6503227e52eacb3a3833f7a772c14b509259","a4a4d03832951a9d9179063d0064281eccd4a46a5427ed11230aeb57b26e04da","a88463737404f73586d2d1ebab0ec9a551c766e0cc35a20c878bb8743aa94145","390fb172c36288f586cdbc4c8d885b808933f2149c7a988dcc8811aa8711a92e","c92efcdd24e04b5a23aebbc8a918f595b2ed3166d52b6b6615c193c8c6ed518f","64d42439f8a6748190441f3a864d740cc6fa54c28c4863122dbbb18c3fa52586","17dc079035d91e4e5d637260690abefded5edd7a0855c48feb1e84e868816276","b77187dfddd934cd2e25e4a6e3bfde5046241ae9cccb8ae2946bf0087e32dbe7","9a4f7e611443b404a3c6ade9dc4f56263cc091b5406a2af4cbc758a4264dcbab",