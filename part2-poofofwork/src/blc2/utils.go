// auth: kunlun
// date: 2019-01-09
// description:
package blc2

import (
	"bytes"
	"encoding/binary"
	"log"
)

//int64转字节数组
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
