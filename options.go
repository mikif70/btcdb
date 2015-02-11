package main

const (
	//	DnsSeed ['dnsseed.bluematt.me', 'dnsseed.bitcoin.dashjr.org', 'seed.bitcoinstats.com', 'bitseed.xf2.org']
	User     = "info"
	Password = "passw0rd"
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
