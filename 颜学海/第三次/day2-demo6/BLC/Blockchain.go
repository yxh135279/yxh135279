package BLC

import (
	"github.com/boltdb/bolt"
	"fmt"
	"log"
	"os"
	"encoding/hex"
	"math/big"
	"time"
)

//数据库库
const dbName = "blockchain.db"
//表名
const tableName = "blocks"

type Blockchain struct {
	Tip []byte //存储最新区块的hash
	DB *bolt.DB //数据库对象 结构里的结构一般要用*指针
}



//生成迭代器
func (blockchain *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{CurrentHash:blockchain.Tip,DB:blockchain.DB}
}




//遍历数据库
func (blockchain *Blockchain) PrintChain() {
	blcIterator := blockchain.Iterator()
	for {
			block := blcIterator.Next()
			fmt.Printf("Height:%d\n",block.Height)
			fmt.Printf("PrevBlockHash : %x\n",block.PrevBlockHash)
			fmt.Printf("Data:%s\n",block.Data)
			fmt.Printf("Hash:%x\n",block.Hash)
			fmt.Printf("Nonce:%d\n",block.Nonce)
			fmt.Printf("Timestamp：%s\n",time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
			var hashInt big.Int
			hashInt.SetBytes(block.PrevBlockHash)

			if big.NewInt(0).Cmp(&hashInt) == 0 {
				break
			}
			//比较高度也可以
			//if block.Height == 1 {
			//	break
			//}
		}
}


//1.创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	pwd, _ := os.Getwd()
	fmt.Println("in day2-demo3中",pwd)

	//打开数据库 注意生成数据库文件的路径，在这里执行会生成在根目录下，如果要生成在本目录下，可在命令行编译执行main
	//注意，此处的路径除了当前路径，或只能通过系统函数获取，不能写成死路径，否则报错
	//db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	//db, err := bolt.Open("my.db", 0666, &bolt.Options{ReadOnly: true})
	//db,err := bolt.Open(pwd+"/blockchain/day2-demo3/BLC/"+dbName,0600,nil)
	db,err := bolt.Open(dbName,0600,nil)
	if err != nil {
		log.Panic(err)
	}
	//数据库打开后一定要关闭
	//defer db.Close()

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		//不存在则创建表
		table,err := tx.CreateBucketIfNotExists([]byte(tableName))
		if err != nil && table != nil{
			log.Panic(err)
		}
		//创建区块
		genesisBlock := CreateGenesisBlock("Genesis Block ...")
		blockBytes := genesisBlock.Serialize()
		blockHash = genesisBlock.Hash
		//存储新的区块hash值到数据库表中
		err = table.Put(blockHash, blockBytes)
		//存储最新的区块hash值到blockchain中
		err = table.Put([]byte("l"), genesisBlock.Hash)
		if err != nil {
			log.Panic(err)
		}


		fmt.Println(genesisBlock.Nonce)

		err = table.Put([]byte("yxh135279"),blockBytes)
		if err != nil {
			log.Panic(err)
		}
		return nil
	})

	//返回区块链对象
	return &Blockchain{Tip:blockHash, DB:db}
}

//2.新增普通区块并添加到区块链
func (blc *Blockchain) AddBlockToBlockchain(data string) {
	err := blc.DB.Update(func(tx *bolt.Tx) error {
			//获取表名
			table := tx.Bucket([]byte(tableName))
			if table != nil {
				//获取数据库里最新的区块数据
				blockBytes := table.Get(blc.Tip)
				//反序列化
				block := DeserializeBlock(blockBytes)
				//产生新的区块
				newBlock := NewBlock(data,block.Height + 1,block.Hash)
				//新区块序列化
				newBlockBytes := newBlock.Serialize()
				//存储新区块
				err := table.Put(newBlock.Hash,newBlockBytes)
				if err != nil {
					log.Panic(err)
				}
				//更新最新的区块hash
				err = table.Put([]byte("l"),newBlock.Hash)
				if err != nil {
					log.Panic(err)
				}
				//更新区块的最新的hash
				fmt.Println("新区块的Hash:",hex.EncodeToString(newBlock.Hash))
				blc.Tip = newBlock.Hash
			}
			return nil

	})
	if err != nil {
		log.Panic(err)
	}
}