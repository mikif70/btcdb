// btcDB
package main

import (
	"flag"
	"fmt"
	"github.com/mikif70/btcdb/btclib"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
)

func main() {
	var cmd = flag.String("c", "", "cmd")
	var opt = flag.String("o", "", "option")
	var profile = flag.String("profile", "", "write cpu profile to file")
	flag.Parse()

	if *profile != "" {
		f, err := os.Create(*profile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	switch *cmd {
	case "count":
		fmt.Println("Tot Blocks: ", btclib.BlockCount())
	case "hash":
		var i, _ = strconv.Atoi(*opt)
		fmt.Printf("Hash of %s: %s\n", *opt, btclib.BlockHash(i))
	case "tx":
		fmt.Printf("TX of %s: %s\n", *opt, btclib.GetTransaction(opt))
	case "alltx":
		btclib.AllTxInsert()
	case "blocktx":
		btclib.BlockTxInsert()
	case "rttest":
		btclib.TestRT()
	case "rtinsert":
		btclib.InsertRt()
	default:
		btclib.BlockInsert()
	}
}
