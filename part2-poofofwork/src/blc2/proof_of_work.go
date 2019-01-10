// auth: kunlun
// date: 2019-01-09
// description:  POW 工作量证明
package blc2

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

//位移操作
//000000001
//如果将第三位为1 那么需要左移  8-3=5  00100000
//<< 5
const targetBits = 16 //挖矿难度

//最大值
var maxNonce = math.MaxInt64

type ProofOfWork struct {
	block  *Block   //当前需要验证的区块
	target *big.Int //大数存储，大数据使用，用来存储挖矿难度，也就是区块难度值
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	//左移  该值应该是 2^256-24 次方
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{block, target}
	return pow
}

func (pow *ProofOfWork) Validate() bool {

	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	//fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		// \r 表示只打印一条数据，并且将上一条数据从控制台擦除
		fmt.Printf("\r Mining: hash= %x", hash)
		hashInt.SetBytes(hash[:])
		// pow.target compare hash
		// hashInt < target = -1  证明找到了可用hash
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

//准备数据
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce))}, []byte{})
	return data
}
