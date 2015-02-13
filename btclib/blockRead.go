// blockRead.go
package btclib

import (
	//	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
	//	"gopkg.in/mgo.v2/bson"
	//	"sync"
)

func BlockCount() float64 {
	params := make([]int, 0)
	var blockCount, _ = callCmd("getblockcount", params)
	return blockCount.(float64)
}

func BlockHash(block int) string {
	params := make([]int, 1)
	params[0] = block
	var hash, _ = callCmd("getblockhash", params)
	return hash.(string)
}

func BlockInsert() {

	session, err := mgo.Dial(MongoD)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var blockCount, _ = callCmd("getblockcount", make([]int, 0))

	db := session.DB("btc").C("block")
	tx := session.DB("btc").C("tx")

	var retval map[string]interface{}
	db.Find(nil).Sort("-count").One(&retval)

	var start int
	if retval != nil {
		start = retval["count"].(int)
	} else {
		start = 0
	}

	var stop = int(blockCount.(float64))
	//	stop = start + 1000
	var startTime = time.Now()

	//	var throttle = make(chan int, maxConcurrency)
	//	var wg sync.WaitGroup
	paramsInt := make([]int, 1)
	paramsString := make([]string, 1)
	for i := start; i < stop; i++ {
		//		throttle <- 1
		//		wg.Add(1)
		//		go func(db *mgo.Collection, tx *mgo.Collection, i int) {
		paramsInt[0] = i
		var hash, _ = callCmd("getblockhash", paramsInt)
		paramsString[0] = hash.(string)
		var block, _ = callCmd("getblock", paramsString)
		var newBlock = block.(map[string]interface{})
		newBlock["count"] = int(i)
		err = db.Insert(newBlock)
		if err != nil {
			fmt.Println("Error block insert: ", err)
			fmt.Println(i)
		}
		var arr = newBlock["tx"].([]interface{})

		var tot = len(arr)
		//		fmt.Println(arr, tot)
		for a := 0; a < tot; a++ {
			//			fmt.Println(arr[a])
			//			TransactionInsert(arr[a].(string))
			txInsert(arr[a].(string), tx)
		}
		//			<-throttle
		//		}(db, tx, i)
		fmt.Print(".")
	}
	//	wg.Wait()
	fmt.Println("")
	db.Find(nil).Sort("-count").One(&retval)
	fmt.Println("End: ", retval["count"])
	fmt.Println(time.Since(startTime))
}