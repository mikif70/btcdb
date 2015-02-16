// btcDB
package main

import (
	"flag"
	"fmt"
	"github.com/mikif70/btcdb/btclib"
	"strconv"
)

func main() {
	var cmd = flag.String("c", "", "cmd")
	var opt = flag.String("o", "", "option")
	flag.Parse()

	switch *cmd {
	case "count":
		fmt.Println("Tot Blocks: ", btclib.BlockCount())
	case "hash":
		var i, _ = strconv.Atoi(*opt)
		fmt.Printf("Hash of %s: %s\n", *opt, btclib.BlockHash(i))
	case "tx":
		fmt.Printf("TX of %s: %s\n", *opt, btclib.GetTransaction(opt))
	case "pushtx":
		btclib.TxInsert(*opt)
	case "alltx":
		btclib.AllTxInsert()
	case "blocktx":
		btclib.BlockTxInsert()
	default:
		btclib.BlockInsert()
	}
}
