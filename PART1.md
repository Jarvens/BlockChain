#### Block 结构体

Block结构体是区块链最基础的数据机构，区块链由N个区块组成

```go
type Block struct{
    Timestamp     int64     //时间戳
    PrevBlockHash []byte    //上一个区块的Hash
    Data          []byte    //交易数据
    Hash          []byte    //当前区块Hash
}
```

当前使用的哈希算法为Sha256，也就是256bit位。

1byte = 8 bit

哈希数值占用的字节数为   256 / 8 = 32  byte 

二进制则表示为  32 * 2 = 64 个数字

> 创建一个区块

代码含义为：

- 定义一个`NewBlock`方法
- 交易数据`data`
- 上一个区块的哈希`prevHash`
- 返回区块Block的地址 `*Block`

```go
package blc
import "time"

func NewBlock(data string,prevHash []byte) *Block{
    block:=&Block{time.Now().Unix(), preBlockHash, []byte(data), []byte{}}
    return block
}
```

接下来打印该区块信息，看一下是不是我们想象的样子

```go
package main


//此处由于是第一个区块，那么我们传入的上一个区块的哈希就全部给0
func main(){
    block:=blc.NewBlock("this is block",[]{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
    
    fmt.Println("block info %v",block)
}
```

```go
{1547122543 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [116 104 105 115 32 105 115 32 98 108 111 99 107] []}
```

控制台打印的结果好像并不是我们希望的样子，而且当前区块的哈希为空，接下来我们再改造一下这个方法

```go
package blc
import "time"

func NewBlock(data string,prevHash []byte) *Block{
    block:=&Block{time.Now().Unix(), preBlockHash, []byte(data), []byte{}}
    block.SetHash()
    return block
}



func (block *Block) SetHash(){
    timeString:=strconv.FormatInt(block.Timestamp,2)
    timestamp:=[]byte(timeString)
    headers:=bytes.Join([][]byte{block.PreBlockHash,block.Data,timestamp},[]byte{})
    hash:=sha256.Sum256(headers)
    block.Hash=hash[:]
}
```

看到我们现在新增了一个`SetHash`方法，这个方法绑定给了Block对象，所以我们可以直接通过block.SetHash()进行调用。

哈希计算的步骤这里列举一下：

- 将时间戳转化为字节数组
- 将除了当前区块哈希以外的属性通过字节数组的方式拼接起来
- 将拼接起来的字节数组进行256哈希
- 将得到的哈希数值赋给当前区块Hash属性

在计算哈希的过程中，我们用到了一个strconv.FormatInt 方法。该方法的第二个参数是可变的，范围是2~36

看到`2~36`的时候我想大家应该知道是什么意思了。2标识的就是2进制…。然后我么能通过bytes.Join方法将字节数组以二维数组的方式拼接起来，进行哈希，并且将hash赋值给当前区块hash

接下来我们打印一下此时的Block数据

```go
{1547123223 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [116 104 105 115 32 105 115 32 98 108 111 99 107] [9 213 183 30 160 73 175 91 225 220 225 250 175 210 234 199 250 120 203 99 134 201 111 227 51 252 145 158 141 220 63 145]}
```

我们看到打印的哈希属性已经有值了，但是都是数字。这是因为我们输出的是二进制编码，我们想要输出正确的哈希值应该这样打印

```go
fmt.Printf("block hash: %x",block.Hash)
//block hash: 2102a25333c4b671557f83eacfb3e2b215c23c905e04efbac8cf81b6e49c7f8c
```

到这里我们已经生成了一个区块，但是好像还有一点不方便的地方。哪里呢？

我们每次生成区块的时候都需要手动传入两个参数，一个是交易数据，另一个就是上一个区块的哈希，那我们就给区块一个方法我们直接调用就好了。

```go
// 初始化区块
func NewInitialBlock()*Block{
    return NewBlock("this is first block",[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}
```

这样的话我们就可以直接调用了。

#### BlockChain

有了区块，我们怎么生成区块链呢？区块链就是一个区块的数组，我们先来看下`BlockChain`的数据结构

```go
type BlockChain struct{
    Block []*Block
}
```

创建我们的第一个区块链

```go
package blc

//区块链结构体
type BlockChain struct{
    Block *[]Block
}

//生成区块链
func NewBlockChain()*BlockChain{
    return &BlockChain{[]*Block{NewInitialBlock()}}
}
```

打印一下我们的第一个BlockChain

```go
package main

func main(){
    blockChain:=NewBlockChain()
    for _,block:=range blockChain{
        fmt.Printf("Data: %s \n", string(block.Data))
		fmt.Printf("PreBlockHash: %x \n", string(block.PrevBlockHash))
		fmt.Printf("TimeStamp: %s \n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash: %x \n", block.Hash)
    }
}
//Data: Genenis Block 
//PreBlockHash: 00000000000000000000000000000000 
//TimeStamp: 2019-01-10 08:52:58 PM 
//Hash: f1df7a893026d2e5ab1afb46ec3597627878b032fa7b019a61061648f0661edc 
```

可以看到我们的第一个区块链的雏形已经出来了，但是仅有一个不够啊。我们还需要像该链加入区块

```go
func (blockChain *BlockChain)AddBlock(data string){
    preBlock:=blockChain.Block[len(blockChain)-1]
    newBlock:=NewBlock(data,preBlock.Hash)
    blockChain.Block=append(blockChain.Block,newBlock)
}
```

- 首先我们给BlockChain结构体绑定AddBlock方法，并且将交易数据作为参数传递
- 拿到上一个区块信息
- 创建新区块，并且将上一个区块的哈希传入
- 将新创建的区块加入当前链中
