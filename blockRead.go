package main

import (
	//	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
)

func main() {

	session, err := mgo.Dial("10.39.81.90")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	paramsInt := make([]int, 0)

	var blockCount = Call("getblockcount", paramsInt)

	db := session.DB("btc").C("block")

	var retval map[string]interface{}
	db.Find(nil).Sort("-count").One(&retval)

	paramsInt = make([]int, 1)
	paramsString := make([]string, 1)
	for i := retval["count"].(int); i < int(blockCount.(float64)); i++ {
		paramsInt[0] = int(i)
		var hash = Call("getblockhash", paramsInt)
		paramsString[0] = hash.(string)
		var block = Call("getblock", paramsString)
		block.(map[string]interface{})["count"] = int(i)
		//		fmt.Println(block.(map[string]interface{}))
		err = db.Insert(block.(map[string]interface{}))
		if err != nil {
			fmt.Println(i)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("")
	db.Find(nil).Sort("-count").One(&retval)
	fmt.Println("End: ", retval["count"])
}
