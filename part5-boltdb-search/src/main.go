// auth: kunlun
// date: 2019-01-11
// description:
package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

//数据库名称
const dbFile = "blockchain.db"

//仓库
const blockBucket = "blocks"

func main() {

	// ----------创建数据库---------
	// 如果数据库存在，打开，否则，创建
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//查询数据
	err = db.View(func(tx *bolt.Tx) error {
		//获取表
		b := tx.Bucket([]byte(blockBucket))
		valueByte := b.Get([]byte("kunlun"))
		fmt.Printf("%s", valueByte)
		return nil
	})

	if err != nil {
		panic(err)
	}

}
