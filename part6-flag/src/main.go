// auth: kunlun
// date: 2019-01-11
// description:
package main

import (
	"flag"
	"fmt"
)

// need 查询 flag详细用法
func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numberPtr := flag.Int("numb", 42, "an int")

	//所有指令声明完成之后调用 flag.parse() 来执行命令解析
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("number:", *numberPtr)
	fmt.Println("tail", flag.Args())
	//
	flag.Usage()
}
