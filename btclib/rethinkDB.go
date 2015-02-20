// rethinkDB
package btclib

import (
	"fmt"
	"log"

	rt "github.com/dancannon/gorethink"
)

func TestRT() {
	var session *rt.Session
	session, err := rt.Connect(rt.ConnectOpts{
		Address:  Rethink[0],
		Database: "test",
		MaxIdle:  10,
		MaxOpen:  10,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	session.SetMaxOpenConns(5)

	resp, err := rt.Table("test").Insert("test").RunWrite(session)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(resp)
}

func openRT() *rt.Session {
	//	var session *rt.Session
	session, err := rt.Connect(rt.ConnectOpts{
		Address:  Rethink[0],
		Database: "btc",
		MaxIdle:  10,
		MaxOpen:  20,
	})
	if err != nil {
		log.Fatalln("Open error: ", err.Error())
	}

	session.SetMaxOpenConns(10)

	return session
}

func insertRT(session *rt.Session, data []interface{}) {

	_, err := rt.Table("block").Insert(data).RunWrite(session)
	if err != nil {
		fmt.Println()
		fmt.Println("Write Error: ", err.Error())
	}

	//	fmt.Println(resp)
}

func InsertRt() {

	session := openRT()
	defer session.Close()

	var blockCount, _ = callCmd("getblockcount", make([]int, 0))
	start := 0
	stop := int(blockCount.(float64))

	params := make([]interface{}, 1)
	blocks := make([]interface{}, 0)
	for i := start; i < stop; i++ {
		params[0] = i
		var hash, _ = callCmd("getblockhash", params)
		if hash == nil {
			fmt.Println()
			fmt.Println(i)
			continue
		}
		params[0] = hash.(string)
		var block, _ = callCmd("getblock", params)
		if block == nil {
			fmt.Println()
			fmt.Println(i)
			continue
		}
		var newBlock = block.(map[string]interface{})
		newBlock["count"] = int(i)
		blocks = append(blocks, newBlock)
		if len(blocks) >= maxBulk {
			go insertRT(session, blocks)
			blocks = make([]interface{}, 0)
		}
		fmt.Print(".")
	}
	fmt.Println("")
}
