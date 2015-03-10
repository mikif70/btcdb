// mongoDB.go
package btclib

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func openDB() (*mgo.Session, *mgo.Collection, *mgo.Collection) {
	session, err := mgo.Dial(Mongo[0])
	if err != nil {
		panic(err)
	}

	db := session.DB("btc").C("block")
	tx := session.DB("btc").C("tx")

	return session, db, tx
}

func startStop(db *mgo.Collection) (int, int) {
	var blockCount, _ = callCmd("getblockcount", make([]int, 0))

	var retval map[string]interface{}
	db.Find(nil).Sort("-count").One(&retval)

	var start int
	if retval != nil {
		start = retval["count"].(int) + 1
	} else {
		start = 0
	}

	var stop = int(blockCount.(float64))
	//	stop = start + 1000

	return start, stop
}

func checkCount(count int, db *mgo.Collection) bool {
	found, err := db.Find(bson.M{"count": count}).Count()
	if err != nil {
		fmt.Printf("Error found: %v\n", count)
		return false
	}
	if found <= 0 {
		return false
		fmt.Printf("Not Found: %v\n", count)
	}
	return true
}

func endLog(db *mgo.Collection, startTime time.Time) {
	var retval map[string]interface{}
	db.Find(nil).Sort("-count").One(&retval)
	fmt.Println("End: ", retval["count"])
	fmt.Println(time.Since(startTime))
}
