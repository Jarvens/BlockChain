// auth: kunlun
// date: 2019-01-12
// description:
package main

import (
	"blc6"
	"fmt"
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

	for {
		blockChainIterator = blockChain.Iterator()
		fmt.Printf("%x\n", blockChainIterator.CurrentHash)

		blockChainIterator = blockChainIterator.Next()

	}
}
