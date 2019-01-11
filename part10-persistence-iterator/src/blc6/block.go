// auth: kunlun
// date: 2019-01-09
// description:
package blc6

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  //区块创建时间
	PrevBlockHash []byte //上一区块 hash
	Data          []byte //交易数据
	Hash          []byte //当前区块hash
	Nonce         int    //随机数
}

// 16进制
// 64个数字
// 32 字节
// 256 bit  32 * 8
// factory method create block
func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), preBlockHash, []byte(data), []byte{}, 0}
	//block.SetHash()

	pow := NewProofOfWork(block)
	// 执行一次工作量证明
	nonce, hash := pow.Run()
	//设置区块hash
	block.Hash = hash[:]
	//设置nonce
	block.Nonce = nonce

	// 校验区块有效性
	isValid := pow.Validate()
	fmt.Println("\n ", isValid)
	fmt.Println("\n")
	//返回区块
	return block
}

// 1.将时间戳转化为字节数组
// 2.将除了Hash以外的其它属性以字节数组的形式全拼接起来
// 3.将拼接起来的数据进行256Hash
// 4.将Hash赋给当前Hash属性字段
func (block *Block) SetHash() {

	//base: 2~36 表示进制
	timeString := strconv.FormatInt(block.Timestamp, 2)
	timestamp := []byte(timeString)
	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	// 将hash 完全拷贝一份给 block.hash属性
	block.Hash = hash[:]
}

// 初始化创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genenis Block", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

// serialize
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encode := gob.NewEncoder(&result)
	err := encode.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()

}

// deserialize
func Deserialize(b []byte) *Block {

	var block Block
	decode := gob.NewDecoder(bytes.NewReader(b))
	err := decode.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
