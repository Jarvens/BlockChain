// auth: kunlun
// date: 2019-01-10
// description:
package blc3

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	Data          []byte
	Hash          []byte
	Nonce         int
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
