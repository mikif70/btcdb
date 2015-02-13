// riakDB
package btclib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func riakTxInsert() {

}

func riakBlockInsert() {

}

func riackInsert(data interface{}) (interface{}, error) {

	jrequest, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", Riak, bytes.NewBuffer(jrequest))
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error riak call: ", req)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		fmt.Println("Error unmarshal: ", body)
		return nil, err
	}

	return dat["result"], nil
}

func main() {
	fmt.Println("Hello World!")
}
