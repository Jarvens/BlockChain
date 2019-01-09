// auth: kunlun
// date: 2019-01-09
// description:
package blc2

type BlockChain struct {
	Block []*Block
}

// 新增区块
func (blockChain *BlockChain) AddBlock(data string) {
	//1.创建新区块
	preBlock := blockChain.Block[len(blockChain.Block)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	//2.将区块添加到Block
	blockChain.Block = append(blockChain.Block, newBlock)
}

// 创建一个带有创世区块的区块链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

// int64类型时间数据转换为Unix时间  2006xxx为固定格式
// time.Unix(block.Timestamp,0).Format("2006-01-02 03:04:05 PM")
