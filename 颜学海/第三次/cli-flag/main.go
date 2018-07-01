package main

import (
	"fmt"
	"os"
	"flag"
	"log"
)

//func main() {
//
//	arr := os.Args
//	isValideArgs(len(arr))
//
//	flagString := flag.String("printchain","","输出所有的区块信息")
//	flagInt := flag.Int("number",6,"请输入一个整数")
//	flagBool:= flag.Bool("open",false,"判断真假")
//	flag.Parse()
//	fmt.Printf("%s\n",*flagString)
//	fmt.Printf("%d\n",*flagInt)
//	fmt.Printf("%v\n",*flagBool)
//}

func main() {
	arr := os.Args
	isValideArgs(len(arr))
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("prntchain",flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "http://yxh135279","交易数据。。。")

	switch arr[1] {
	case "addBlock":
		err := addBlockCmd.Parse(arr[2:])
		if err != nil {
			log.Panic(err)
		}
	case "prntchain":
		err := printChainCmd.Parse(arr[2:])
		if err != nil {
			log.Panic(err)
		}
	default :
		printUsage()
		os.Exit(1)
	}
	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		fmt.Println(*flagAddBlockData)
	}
	if printChainCmd.Parsed() {
		fmt.Println("输出所有区块的数据........")
	}


}

func isValideArgs(len int) {
	if len < 2 {
		printUsage()
		os.Exit(1)
	}

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\taddblock -data DATA -- 交易数据.")
	fmt.Println("\tprintchain -- 输出区块信息.")
}