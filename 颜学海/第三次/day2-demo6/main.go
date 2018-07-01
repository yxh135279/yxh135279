package main

import (
	"gostudy/blockchain/day2-demo6/BLC"
	"fmt"
	"encoding/hex"
)

/*
Bolt就是这么一个纯粹的Go语言版的嵌入式key/value的数据库，而且在Go的应用中很方便地去用作持久化
Bolt类似于LMDB，这个被认为是在现代kye/value存储中最好的。但是又不同于LevelDB，BoltDB支持完全可序列化的ACID事务，也不同于SQLlite，BoltDB没有查询语句，对于用户而言，更加易用
BoltDB将数据保存在一个单独的内存映射的文件里。它没有wal、线程压缩和垃圾回收；它仅仅安全地处理一个文件

LevelDB和BoltDB的不同:
LevelDB是Google开发的，也是一个k/v的存储数据库，和BoltDB比起起来有很大的不同。对于使用者而言，最大的不同就是LevelDB没有事务
LevelDB实现了一个日志结构化的merge tree,有序的key/value存储在不同文件的之中，并通过“层级”把它们分开并且周期性地将小的文件merge为更大的文件,这让其在随机写的时候会很快，但是读的时候却很慢。这也让LevelDB的性能不可预知
但数据量很小的时候，它可能性能很好，但是当随着数据量的增加，性能只会越来越糟糕

BoltDB使用一个单独的内存映射的文件，实现一个写入时拷贝的B+树，这能让读取更快。而且，BoltDB的载入时间很快，特别是在从crash恢复的时候，因为它不需要去通过读log（其实它压根也没有）去找到上次成功的事务，它仅仅从两个B+树的根节点读取ID
Bolt将所有数据都存储在一个文件中，这让它很容易使用和部署，你可以get/set数据和处理error
读-写方式开始于db.Update方法
只读事务在db.View函数之中：

安装：go get github.com/boltdb/bolt/

 */

//func main() {
//	pwd, _ := os.Getwd()
//	fmt.Println("in day2-demo3中",pwd)
//
//	//打开数据库 注意生成数据库文件的路径，在这里执行会生成在根目录下，如果要生成在本目录下，可在命令行编译执行main
//	//注意，此处的路径除了当前路径，或只能通过系统函数获取，不能写成死路径，否则报错
//	//db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
//	//db, err := bolt.Open("my.db", 0666, &bolt.Options{ReadOnly: true})
//	db,err := bolt.Open(pwd+"/blockchain/day2-demo3/BLC/blcok.db",0600,nil)
//	if err != nil {
//		log.Panic(err)
//	}
//	//数据库打开后一定要关闭
//	defer db.Close()
//
//
//	//执行数据库的事务更新
//	//err = db.Update(func(tx *bolt.Tx) error {
//	//	table := tx.Bucket([]byte("block"))
//	//	if table == nil {
//	//		table,err = tx.CreateBucket([]byte("block"))
//	//		if err != nil {
//	//			log.Panic(err)
//	//			//log.Fatal(err)
//	//		}
//	//	}
//	//
//	//	block := BLC.CreateGenesisBlock("Genesis Block ...")
//	//	blockBytes := block.Serialize()
//	//
//	//	fmt.Println(block.Nonce)
//	//
//	//	err = table.Put([]byte("yxh135279"),blockBytes)
//	//	if err != nil {
//	//		log.Panic(err)
//	//		//log.Fatal(err)
//	//	}
//	//	return nil
//	//})
//
//	//查询数据
//	err = db.View(func(tx *bolt.Tx) error {
//		table := tx.Bucket([]byte("block"))
//		if table != nil {
//			blockBytes := table.Get([]byte("yxh135279"))
//			fmt.Println(blockBytes)
//			block := BLC.DeserializeBlock(blockBytes)
//			fmt.Println(block.Nonce)
//		}
//
//		return nil //返回nil则会提交数据，否则会回滚
//	})
//
//	fmt.Println("in day2-demo3 start end")
//
//
//	if err != nil {
//		log.Panic(err)
//	}
//
//}

func main() {
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(hex.EncodeToString(blockchain.Tip))


	blockchain.AddBlockToBlockchain("Send 100RMB To zhangshang")
	blockchain.AddBlockToBlockchain("Send 200RMB To zhangshang")
	blockchain.AddBlockToBlockchain("Send 300RMB To zhangshang")
	blockchain.AddBlockToBlockchain("Send 400RMB To zhangshang")

	fmt.Println(hex.EncodeToString(blockchain.Tip))

	blockchain.PrintChain()


}
