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

		var tot = len(arr)
		for a := 0; a < tot; a++ {
			txInsert(arr[a].(string), tx)
		}
		fmt.Print(".")
	}
	fmt.Println("")

	endLog(db, startTime)
}
