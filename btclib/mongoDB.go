// mongoDB.go
package btclib

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

func openDB() (*mgo.Session, *mgo.Collection, *mgo.Collection) {
	session, err := mgo.Dial(MongoD)
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
		start = retval["count"].(int)
	} else {
		start = 0
	}

	var stop = int(blockCount.(float64))
	//	stop = start + 1000

	return start, stop
}

func endLog(db *mgo.Collection, startTime time.Time) {
	var retval map[string]interface{}
	db.Find(nil).Sort("-count").One(&retval)
	fmt.Println("End: ", retval["count"])
	fmt.Println(time.Since(startTime))
}
