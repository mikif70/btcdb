// rethinkDB
package btclib

import (
	"fmt"
	"log"
	//	"strconv"

	rt "github.com/dancannon/gorethink"
)

func TestRT() {
	var session *rt.Session
	session, err := rt.Connect(rt.ConnectOpts{
		Address:  Rethink[0],
		Database: "test",
		MaxIdle:  10,
		MaxOpen:  10,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	session.SetMaxOpenConns(5)

	_, err = rt.Table("test").Insert("test").RunWrite(session)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func openRT() *rt.Session {
	//	var session *rt.Session
	session, err := rt.Connect(rt.ConnectOpts{
		Address:  Rethink[0],
		Database: "btc",
		MaxIdle:  10,
		MaxOpen:  30,
	})
	if err != nil {
		log.Fatalln("Open error: ", err.Error())
	}

	return session
}

func insertRT(session *rt.Session, table string, data []interface{}) {

	resp, err := rt.Table(table).Insert(data).RunWrite(session)
	if err != nil {
		fmt.Println()
		fmt.Println("Write Error: ", err.Error(), resp)

	}

	//	fmt.Println("Inserted: ", resp.Inserted)
}

func getTxStartStop(session *rt.Session) (int, int) {
	bs, err := rt.Table("block").Max("count").Field("count").Run(session)
	defer bs.Close()
	var stop interface{}
	if err == nil {
		bs.One(&stop)
	} else {
		stop, _ = callCmd("getblockcount", make([]int, 0))
	}

	tx, err := rt.Table("transaction").Max("count").Field("count").Run(session)
	defer tx.Close()
	var start interface{}
	if err == nil {
		tx.One(&start)
	} else {
		start = 0.0
	}

	return int(start.(float64)), int(stop.(float64))
}

func getBlockStartStop(session *rt.Session) (int, int) {
	//.db("btc").table("block").max("count").getField("count");
	var stop, _ = callCmd("getblockcount", make([]int, 0))

	res, _ := rt.Table("block").Max("count").Field("count").Run(session)
	defer res.Close()

	var start interface{}
	res.One(&start)

	fmt.Println(start, stop)

	return int(start.(float64)), int(stop.(float64))
}

func getTransactions(txArray []interface{}, count int) (txs []interface{}) {
	tot := len(txArray)
	txs = make([]interface{}, 0)
	var retval = make(map[string]interface{})

	for i := 0; i < tot; i++ {
		retval = GetTransaction(txArray[i])
		if retval == nil {
			fmt.Println("Error")
			continue
		}
		retval["count"] = count
		txs = append(txs, retval)
	}

	return txs
}

func InsertTxRt() {
	session := openRT()
	defer session.Close()

	start, stop := getTxStartStop(session)

	if start < stop {
		//		params := make([]interface{}, 1)
		//		txs := make([]interface{}, 0)
		for i := start; i <= stop; i++ {
			txArray, err := rt.Table("block").GetAllByIndex("count", i).Field("tx").Run(session)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				continue
			}
			var retval []interface{}
			txArray.One(&retval)
			txs := getTransactions(retval, i)
			//			fmt.Printf("%d - %+v\n", i, txs)
			fmt.Printf("%d\n", i)
			if len(txs) > 0 {
				go insertRT(session, "transaction", txs)
			}
			txArray.Close()
		}
	}
}

func InsertBlocksRt() {

	session := openRT()
	defer session.Close()

	start, stop := getBlockStartStop(session)

	if start < stop {
		params := make([]interface{}, 1)
		blocks := make([]interface{}, 0)
		for i := start; i <= stop; i++ {
			params[0] = i
			var hash, _ = callCmd("getblockhash", params)
			if hash == nil {
				fmt.Println()
				fmt.Println(i)
				continue
			}
			params[0] = hash.(string)
			var block, _ = callCmd("getblock", params)
			if block == nil {
				fmt.Println()
				fmt.Println(i)
				continue
			}
			var newBlock = block.(map[string]interface{})
			newBlock["count"] = int(i)
			blocks = append(blocks, newBlock)
			if len(blocks) >= maxBulk {
				go insertRT(session, "block", blocks)
				blocks = make([]interface{}, 0)
			}
			fmt.Print(".")
		}
		if len(blocks) >= 0 {
			//		fmt.Println("inserting... ", blocks)
			insertRT(session, "block", blocks)
			//		blocks = make([]interface{}, 0)
		}
		fmt.Println("")
	}
}

func BlockVerifyRt() {
	session := openRT()
	defer session.Close()

	_, stop := getBlockStartStop(session)

	for i := 0; i < stop; i++ {
		resp, err := rt.Table("block").GetAllByIndex("count", i).Field("count").Run(session)
		if err != nil {
			fmt.Printf("Error: %d - %s\n", i, err.Error())
		}
		resp.Close()
	}

}
