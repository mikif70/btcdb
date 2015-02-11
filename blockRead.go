package main

import (
	//	"encoding/json"
	"fmt"
)

func main() {

	paramsInt := make([]int, 0)

	var blockCount = Call("getblockcount", paramsInt)

	paramsInt = make([]int, 1)
	paramsString := make([]string, 1)
	for i := 0.0; i < blockCount.(float64); i++ {
		paramsInt[0] = int(i)
		var hash = Call("getblockhash", paramsInt)
		paramsString[0] = hash.(string)
		var block = Call("getblock", paramsString)
		fmt.Println(block.(map[string]interface{}))
	}
}
