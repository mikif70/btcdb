// blockRead.go
package btclib

import (
	//	"encoding/json"
	"fmt"
	//	"gopkg.in/mgo.v2"
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

	session, db, _ := openDB()
	defer session.Close()

	start, stop := startStop(db)

	var startTime = time.Now()

	params := make([]interface{}, 1)
	blocks := make([]interface{}, 0)
	for i := start; i < stop; i++ {
		params[0] = i
		var hash, _ = callCmd("getblockhash", params)
		params[0] = hash.(string)
		var block, _ = callCmd("getblock", params)
		var newBlock = block.(map[string]interface{})
		newBlock["count"] = int(i)
		blocks = append(blocks, newBlock)
		if len(blocks) >= maxBulk {
			err := db.Insert(blocks...)
			//err := db.Insert(newBlock)
			if err != nil {
				fmt.Println("")
				fmt.Println("Error block insert: ", err)
				fmt.Println(i)
			}
			blocks = make([]interface{}, 0)
		}
		fmt.Print(".")
	}
	fmt.Println("")

	endLog(db, startTime)
}

func BlockTxInsert() {

	session, db, tx := openDB()
	defer session.Close()

	start, stop := startStop(db)

	var startTime = time.Now()

	paramsInt := make([]int, 1)
	paramsString := make([]string, 1)
	for i := start; i < stop; i++ {
		paramsInt[0] = i
		var hash, _ = callCmd("getblockhash", paramsInt)
		paramsString[0] = hash.(string)
		var block, _ = callCmd("getblock", paramsString)
		var newBlock = block.(map[string]interface{})
		newBlock["count"] = int(i)
		err := db.Insert(newBlock)
		if err != nil {
			fmt.Println("Error block insert: ", err)
			fmt.Println(i)
		}
		var arr = newBlock["tx"].([]interface{})

		//		var tot = len(arr)
		//		for a := 0; a < tot; a++ {
		txInsert(tx, i, arr)
		//		}
		fmt.Print(".")
	}
	fmt.Println("")

	endLog(db, startTime)
}
