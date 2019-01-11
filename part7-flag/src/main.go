// auth: kunlun
// date: 2019-01-11
// description:
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "block data", "usage")
	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}

	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println("No addblock and printchain")

	}
}
