// auth: kunlun
// date: 2019-01-12
// description:
package main

import "blc7"

func main() {

	//创建区块链
	blockChain := blc7.NewBlockChain()
	//创建CLI 对象
	cli := blc7.CLI{blockChain}
	//调用CLI的Run
	cli.Run()

}
