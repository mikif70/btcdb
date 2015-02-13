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
func TransactionInsert(tx string) {
	session, err := mgo.Dial(DBHost)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	db := session.DB("btc").C("tx")

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
		fmt.Println(data)
		fmt.Println(err)
	}
}

func AllTxInsert() {
	session, err := mgo.Dial(DBHost)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	tx := session.DB("btc").C("tx")
	db := session.DB("btc").C("block")

	var rettx map[string]interface{}
	tx.Find(nil).Sort("-_id").One(&rettx)

	var txblock map[string]interface{}
	db.Find(bson.M{"tx": rettx["txid"]}).Sort("-count").One(&txblock)

	var lastblock map[string]interface{}
	db.Find(nil).Sort("-count").One(&lastblock)

	//	db.Find(bson.M)

	//	fmt.Println(txblock["count"], lastblock["count"], rettx["txid"])

	var startTime = time.Now()

	var block map[string]interface{}
	var startBlock = txblock["count"].(int)
	var stopBlock = lastblock["count"].(int)
	//	stopBlock = startBlock + 1000
	for i := startBlock; i < stopBlock; i++ {
		db.Find(bson.M{"count": i}).One(&block)
		var arr = block["tx"].([]interface{})
		var tot = len(arr)
		for a := 0; a < tot; a++ {
			TransactionInsert(arr[a].(string))
			//			txInsert(arr[a].(string), tx)
		}
	}

	var stopTime = time.Since(startTime)
	fmt.Println(stopTime)
}
