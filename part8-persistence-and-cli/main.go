// auth: kunlun
// date: 2019-01-12
// description:
package main

import (
	"blc4"
	"fmt"
)

func main() {
	blockChain := blc4.NewBlockChain()
	fmt.Printf(" blockChain: %v\n", blockChain)
	fmt.Printf(" tip: %x", blockChain.Tip)

}
