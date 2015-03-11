// options.go
package btclib

import ()

const (
	//	DnsSeed ['dnsseed.bluematt.me', 'dnsseed.bitcoin.dashjr.org', 'seed.bitcoinstats.com', 'bitseed.xf2.org']
	User     = "info"
	Password = "passw0rd"
	//	MongoD         = "10.39.81.90:49153"
	MongoD   = "10.39.81.90:27017"
	BitcoinD = "http://10.39.81.85:8332"
	//	maxConcurrency = 10
)

var (
	Riak    = []string{"10.39.80.181:8087", "10.39.80.182:8087", "10.39.80.183:8087", "10.39.80.184:8087", "10.39.80.185:8087", "10.39.80.186:8087"}
	Rethink = []string{"10.39.80.182:28015", "10.39.80.183:28015", "10.39.80.184:28015", "10.39.80.185:28015", "10.39.80.186:28015", "10.39.80.181:28015"}
	Mongo   = []string{"10.39.80.182:27018", "10.39.80.183:27018", "10.39.80.184:27018", "10.39.80.185:27018", "10.39.80.186:27018", "10.39.80.181:27018"}
	maxBulk = 20
)

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
