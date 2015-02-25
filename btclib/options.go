// options.go
package btclib

const (
	//	DnsSeed ['dnsseed.bluematt.me', 'dnsseed.bitcoin.dashjr.org', 'seed.bitcoinstats.com', 'bitseed.xf2.org']
	User     = "info"
	Password = "passw0rd"
	//	MongoD         = "10.39.81.90:49153"
	MongoD   = "10.39.81.90:27017"
	BitcoinD = "http://10.39.81.85:8332"
	//	maxConcurrency = 10
	maxBulk = 100
)

var Riak = []string{"10.39.80.181:8087", "10.39.80.182:8087", "10.39.80.183:8087", "10.39.80.184:8087", "10.39.80.185:8087", "10.39.80.186:8087"}
var Rethink = []string{"10.39.81.85:28115", "10.39.81.85:28215", "10.39.81.85:28315", "10.39.81.85:28415", "10.39.81.85:28515"}

type Options struct {
	User     string
	Password string
	Port     string
	Host     string
}

type RequestInt struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int64  `json:"id"`
	Method  string `json:"method"`
	Params  []int  `json:"params"`
}

type RequestString struct {
	Jsonrpc string   `json:"jsonrpc"`
	Id      int64    `json:"id"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

type Request struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      int64       `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}
