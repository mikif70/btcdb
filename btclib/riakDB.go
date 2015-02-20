// riakDB
package btclib

import (
	//	"bytes"
	//	"encoding/json"
	"fmt"
	//	"io/ioutil"
	"log"
	//	"net/http"

	"github.com/riaken/riaken-core"
)

func OpenRiak() {
	client := riaken_core.NewClient(Riak, 5)
	if err := client.Dial(); err != nil {
		log.Fatal(err.Error()) // all nodes are down
	}
	defer client.Close()
	session := client.Session()
	defer session.Release()

	if !session.Ping() {
		log.Fatal("no ping response")
	}

	buckets, err := session.ListBuckets()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(buckets)
}

func riakTxInsert() {

}

func riakBlockInsert() {

}

/*
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
*/
