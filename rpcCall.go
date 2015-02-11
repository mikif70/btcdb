package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Call(cmd string, params interface{}) interface{} {

	//	fmt.Println(params)

	request := Request{
		Jsonrpc: "1.0",
		Id:      time.Now().Unix(),
		Method:  cmd,
		Params:  params,
	}

	jrequest, _ := json.Marshal(request)

	//	fmt.Println(jrequest)

	req, err := http.NewRequest("POST", "http://10.39.81.85:8332", bytes.NewBuffer(jrequest))
	req.SetBasicAuth(User, Password)
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(req)
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var dat map[string]interface{}

	if err := json.Unmarshal(body, &dat); err != nil {
		fmt.Println(body)
		panic(err)
	}

	return dat["result"]

	return body
}
