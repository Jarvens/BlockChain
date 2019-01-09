// auth: kunlun
// date: 2019-01-09
// description:
package main

import (
	"blc"
	"fmt"
	"time"
)

func main() {

	blockChain := blc.NewBlockChain()
	blockChain.AddBlock("send 10 btc to zhangsan")
	blockChain.AddBlock("send 10 btc to lisi")
	blockChain.AddBlock("send 10 btc to zhaoliu")

	for _, block := range blockChain.Block {
		fmt.Printf("Data: %s \n", string(block.Data))
		fmt.Printf("PreBlockHash: %x \n", string(block.PrevBlockHash))
		fmt.Printf("TimeStamp: %s \n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash: %x \n", block.Hash)
	}
}
