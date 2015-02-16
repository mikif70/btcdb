// transactionRead.go
package btclib

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// GetTransaction get a transaction with hash = tx
func GetTransaction(tx *string) interface{} {
	params := make([]string, 1)
	params[0] = *tx
	var rettx, _ = callCmd("getrawtransaction", params)
	if rettx == nil {
		return nil
	}
	params[0] = rettx.(string)
	var strtx, _ = callCmd("decoderawtransaction", params)
	//	fmt.Println(strtx, err2)

	return strtx
}

// TransactionInsert insert all transaction to the DB
func TxInsert(tx string) {

	session, _, db := openDB()
	defer session.Close()

	txInsert(tx, db)
}

func txInsert(data string, db *mgo.Collection) {
	var retval = GetTransaction(&data)
	if retval == nil {
		fmt.Println("KO")
		return
	}
	err := db.Insert(retval.(map[string]interface{}))
	if err != nil {
		fmt.Println("Error tx insert: ", err)
		fmt.Println(data)
	}
}

func AllTxInsert() {

	session, db, tx := openDB()

	defer session.Close()

	var rettx map[string]interface{}
	tx.Find(nil).Sort("-_id").One(&rettx)

	var startBlock int
	if rettx != nil {
		var txblock map[string]interface{}
		db.Find(bson.M{"tx": rettx["txid"]}).Sort("-count").One(&txblock)
		startBlock = txblock["count"].(int)
	} else {
		startBlock = 0
	}

	var lastblock map[string]interface{}
	db.Find(nil).Sort("-count").One(&lastblock)
	var stopBlock = lastblock["count"].(int)
	//	stopBlock = startBlock + 1000

	var startTime = time.Now()

	var block map[string]interface{}
	for i := startBlock; i < stopBlock; i++ {
		db.Find(bson.M{"count": i}).One(&block)
		var arr = block["tx"].([]interface{})
		var tot = len(arr)
		for a := 0; a < tot; a++ {
			//			TxInsert(arr[a].(string))
			txInsert(arr[a].(string), tx)
		}
		fmt.Print(".")
	}

	fmt.Println("")
	var stopTime = time.Since(startTime)
	fmt.Println(stopTime)
}
