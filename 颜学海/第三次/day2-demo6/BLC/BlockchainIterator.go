package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//迭代器用于区块的遍历
type BlockchainIterator struct {
	CurrentHash []byte //存储最新区块的hash
	DB *bolt.DB //数据库对象 结构里的结构一般要用*指针
}

//开始迭代
func (blcIterator *BlockchainIterator) Next() *Block {
	var block *Block
	err := blcIterator.DB.View(func(tx *bolt.Tx) error {
		table := tx.Bucket([]byte(tableName))
		if table != nil {
			blockBytes := table.Get(blcIterator.CurrentHash)
			block = DeserializeBlock(blockBytes)
			blcIterator.CurrentHash = block.PrevBlockHash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block
}
