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

	//插入或者更新数据
	err = db.Update(func(tx *bolt.Tx) error {
		//判断表是否存在于数据库中
		b := tx.Bucket([]byte(blockBucket))
		if b == nil {

			fmt.Println("No existing blockchain found")
			//创建表
			b, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				panic(err)
			}
			//存储数据
			err = b.Put([]byte("kunlun"), []byte("www.baidu.com"))
			if err != nil {
				panic(err)
			}

			err = b.Put([]byte("canghai"), []byte("www.baidu.com"))
			if err != nil {
				panic(err)
			}
		}
		return nil
	})

}
