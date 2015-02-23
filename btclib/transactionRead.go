// transactionRead.go
package btclib

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// TransactionInsert insert all transaction to the DB
/*
func TxInsert(tx string) {

	session, _, db := openDB()
	defer session.Close()

	txs := make([]interface{}, 1)
	txs[0] = tx

	txInsert(db, txs[:])
}
*/

func txInsert(db *mgo.Collection, count int, data ...[]interface{}) {
	tot := len(data[0])
	txs := make([]interface{}, 0)
	var retval = make(map[string]interface{})
	for i := 0; i < tot; i++ {
		retval = GetTransaction(data[0][i])
		if retval == nil {
			fmt.Println("KO")
			continue
		}
		retval["count"] = count
		txs = append(txs, retval)
	}
	if len(txs) > 0 {
		err := db.Insert(txs...)
		if err != nil {
			fmt.Println("Error tx insert: ", err)
			fmt.Println(data)
		}
	}
}

func AllTxInsert() {

	session, db, tx := openDB()

	defer session.Close()

	var rettx map[string]interface{}
	tx.Find(nil).Sort("-_id").One(&rettx)
	fmt.Println(rettx)

	var startBlock = 0
	if rettx != nil {
		var txblock map[string]interface{}
		db.Find(bson.M{"tx": rettx["txid"]}).Sort("-count").One(&txblock)
		if txblock != nil {
			startBlock = txblock["count"].(int)
		}
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
		//		var tot = len(arr)
		//		for a := 0; a < tot; a++ {
		//			TxInsert(arr[a].(string))
		txInsert(tx, i, arr)
		//		}
		fmt.Print(".")
	}

	fmt.Println("")
	var stopTime = time.Since(startTime)
	fmt.Println(stopTime)
}
