// auth: kunlun
// date: 2019-01-12
// description:
package main

import (
	"blc6"
	"fmt"
	"github.com/boltdb/bolt"
	"math/big"
	"time"
)

func main() {
	blockChain := blc6.NewBlockChain()
	fmt.Println(blockChain)
	fmt.Printf("tip: %x\n", blockChain.Tip)
	blockChain.AddBlock("send 10 btc to zhangsan")
	blockChain.AddBlock("send 20 btc to lisi")
	blockChain.AddBlock("send 30 btc to wangwu")
	fmt.Printf("tip: %x\n", blockChain.Tip)

	var blockChainIterator *blc6.BlockChainIterator
	blockChainIterator = blockChain.Iterator()
	var hashInt big.Int
	for {
		fmt.Printf("%x\n", blockChainIterator.CurrentHash)

		err := blockChainIterator.DB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("blocks"))
			blockBytes := b.Get(blockChainIterator.CurrentHash)
			block := blc6.Deserialize(blockBytes)
			fmt.Printf("Data: %s \n", string(block.Data))
			fmt.Printf("PreBlockHash: %x \n", string(block.PrevBlockHash))
			fmt.Printf("TimeStamp: %s \n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
			fmt.Printf("Hash: %x \n", block.Hash)
			fmt.Printf("Nonce: %d \n", block.Nonce)
			return nil
		})
		if err != nil {
			panic(err)
		}

		//获取下一个迭代器
		blockChainIterator = blockChainIterator.Next()
		//将迭代器中的哈希存储到hashInt中
		hashInt.SetBytes(blockChainIterator.CurrentHash)
		// 如果是创世区块 则退出  创世区块的哈希为0000000000000000
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
}
