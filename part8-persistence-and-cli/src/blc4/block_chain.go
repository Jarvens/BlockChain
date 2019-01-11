// auth: kunlun
// date: 2019-01-09
// description:
package blc4

import (
	"fmt"
	"github.com/boltdb/bolt"
)

//数据库名称
const dbFile = "blockchain.db"

//仓库
const blockBucket = "blocks"

type BlockChain struct {
	Tip []byte   //区块链里面最后一个区块的哈希
	DB  *bolt.DB //数据库
}

// 新增区块
//func (blockChain *BlockChain) AddBlock(data string) {
//	//1.创建新区块
//	preBlock := blockChain.Block[len(blockChain.Block)-1]
//	newBlock := NewBlock(data, preBlock.Hash)
//	//2.将区块添加到Block
//	blockChain.Block = append(blockChain.Block, newBlock)
//}

// 创建一个带有创世区块的区块链
func NewBlockChain() *BlockChain {

	var tip []byte //获取最后一个区块的哈希值
	// 1.尝试打开或者是创建数据库

	// 2.db.update 更新数据
	// 2.1表是否存在，如果不存在，需要创建表
	// 2.2创建区块
	// 2.3需要将区块序列化
	// 2.4把区块的哈希值作为key block序列化数据作为value存储
	// 2.5设置一个key 将哈希作为value再次存储到数据库里面
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		if b == nil {
			fmt.Println("Not Exist")
			block := NewGenesisBlock()
			b, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				panic(err)
			}
			// 将区块序列化后的数据存储到表中
			err = b.Put(block.Hash, block.Serialize())
			if err != nil {
				panic(err)
			}
			err = b.Put([]byte("l"), block.Hash)
			if err != nil {
				panic(err)
			}
			tip = block.Hash
		} else {
			// key属于自定义值
			// 获取最后一个区块的哈希
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
	return &BlockChain{tip, db}
}

// int64类型时间数据转换为Unix时间  2006xxx为固定格式
// time.Unix(block.Timestamp,0).Format("2006-01-02 03:04:05 PM")
