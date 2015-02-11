package main

import (
	//	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
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
	paramsInt = make([]int, 1)
	paramsString := make([]string, 1)
	for i := 0.0; i < blockCount.(float64); i++ {
		paramsInt[0] = int(i)
		var hash = Call("getblockhash", paramsInt)
		paramsString[0] = hash.(string)
		var block = Call("getblock", paramsString)
		//		fmt.Println(block.(map[string]interface{}))
		err = db.Insert(block.(map[string]interface{}))
		if err != nil {
			fmt.Println(i)
		} else {
			fmt.Print(".")
		}
	}
}
