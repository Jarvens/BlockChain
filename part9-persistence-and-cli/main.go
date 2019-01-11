// auth: kunlun
// date: 2019-01-12
// description:
package main

import (
	"blc5"
	"fmt"
)

func main() {

	blockChain := blc5.NewBlockChain()
	fmt.Println(blockChain)
	fmt.Printf("tip: %x\n", blockChain.Tip)
	blockChain.AddBlock("send 10 btc to zhangsan")
	fmt.Printf("tip: %x\n", blockChain.Tip)
}
