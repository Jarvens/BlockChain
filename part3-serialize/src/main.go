// auth: kunlun
// date: 2019-01-10
// description:
package main

import (
	"blc3"
	"fmt"
	"time"
)

func main() {

	block := blc3.Block{time.Now().Unix(), []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []byte("send data"), []byte{}, 0}
	fmt.Printf("%s\n", block.Data)
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("\n\n")
	bytes := block.Serialize()
	fmt.Println(bytes)
	fmt.Println("\n\n")
	blc := blc3.Deserialize(bytes)
	fmt.Printf("%s\n", blc.Data)
	fmt.Printf("%d\n", blc.Nonce)
	fmt.Printf("\n\n")
}
