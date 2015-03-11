// rpcCall.go
package btclib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func callCmd(cmd string, params interface{}) (interface{}, error) {

	//	fmt.Println(params)

	request := Request{
		Jsonrpc: "1.0",
		Id:      time.Now().Unix(),
		Method:  cmd,
		Params:  params,
	}

	jrequest, _ := json.Marshal(request)

	//	fmt.Println(jrequest)

	req, err := http.NewRequest("POST", BitcoinD, bytes.NewBuffer(jrequest))
	req.SetBasicAuth(User, Password)
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error bitcoind call: ", req)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		fmt.Println("Error unmarshal: ", body, err)
		return nil, err
	}

	return dat["result"], nil
}

// GetTransaction get a transaction with hash = tx
func GetTransaction(tx interface{}) map[string]interface{} {
	params := make([]string, 1)
	params[0] = tx.(string)
	var rettx, err = callCmd("getrawtransaction", params)
	if rettx == nil {
		fmt.Println("Error getraw: ", params, err)
		return nil
	}
	params[0] = rettx.(string)
	var strtx, _ = callCmd("decoderawtransaction", params)
	//	fmt.Println(strtx)
	if strtx == nil {
		fmt.Println("Error decoderaw: ", params, err)
		return nil
	}

	return strtx.(map[string]interface{})
	//	return strtx
}
