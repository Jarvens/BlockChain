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