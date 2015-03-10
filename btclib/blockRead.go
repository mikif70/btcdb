// blockRead.go
package btclib

import (
	//	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
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

func blockInsert(db *mgo.Collection, count int, data ...[]interface{}) {

}

func BlockInsert(opt string) {

	session, db, _ := openDB()
	defer session.Close()

	start, stop := startStop(db)
	if opt != "" {
		start, _ = strconv.Atoi(opt)
	}
	fmt.Printf("Start: %d - Stop: %d\n", start, stop)

	var startTime = time.Now()

	params := make([]interface{}, 1)
	//	blocks := make([]interface{}, 0)
	for i := start; i < stop; i++ {
		params[0] = i
		if checkCount(i, db) {
			fmt.Print(":")
			continue
		}
		var hash, _ = callCmd("getblockhash", params)
		if hash == nil {
			fmt.Printf("Error getting blockhash: %d", i)
			continue
		}
		params[0] = hash.(string)
		var block, _ = callCmd("getblock", params)
		if block == nil {
			fmt.Printf("Error getting block: %d", i)
			continue
		}
		var newBlock = block.(map[string]interface{})
		newBlock["count"] = int(i)
		go func() {
			err := db.Insert(newBlock)
			//		blocks = append(blocks, newBlock)
			//		if len(blocks) >= maxBulk {
			//		err := db.Insert(blocks...)
			//err := db.Insert(newBlock)
			if err != nil {
				fmt.Println("")
				fmt.Println("Error block insert: ", i, err)
			}
		}()
		//		blocks = make([]interface{}, 0)
		//		}
		fmt.Print(".")
	}
	//	if len(blocks) > 0 {
	//		err := db.Insert(blocks...)
	//		if err != nil {
	//			fmt.Println("")
	//			fmt.Println("Error block insert: ", err)
	//		}
	//		blocks = make([]interface{}, 0)
	//	}
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

func BlockVerify() {
	session, db, _ := openDB()
	defer session.Close()

	_, stop := startStop(db)

	for i := 0; i < stop; i++ {
		found, err := db.Find(bson.M{"count": i}).Count()
		if err != nil {
			fmt.Printf("Error found: %v\n", i)
		}
		if found == 0 {
			fmt.Printf("Not found: %v\n", i)
		} else if found > 1 {
			fmt.Printf("Found multi: %v - %v\n", found, i)
		}
	}
}
