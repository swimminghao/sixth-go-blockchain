package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

type Block struct {
	//1.版本号
	Version uint64
	//2.前区块哈希
	PreHash []byte
	//3.Merkel根(梅克尔根，这就是一个哈希值，我们先不管，我们后面v4再介绍）
	MerkelRoot []byte
	//4.时间戳
	TimeStamp uint64
	//5.难度值
	Difficulty uint64
	//6.随机数，也就是挖矿要找的数据
	Nonce uint64

	//a. 当前区块哈希，正常比特币区块中没有当前区块的哈希，我们为了方便做了简化！
	Hash []byte
	//b. 数据
	Date []byte
}

//1. 补充区块字段
//2. 更新计算哈希函数
//3. 优化代码

// 2. 创建区块
func NewBlock(date string, preBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PreHash:    preBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,        //随便填写的无效值
		Nonce:      0,        //同上
		Hash:       []byte{}, //先填空，后面再计算
		Date:       []byte(date),
	}
	block.SetHash()
	return &block
}

// 实现一个辅助函数，功能是将uint64转成[]byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

// 3. ⽣成哈希
func (block *Block) SetHash() {

	var blockInfo []byte
	//	1.拼装数据
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PreHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	blockInfo = append(block.PreHash, block.Date...)

	//	2.shah256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
