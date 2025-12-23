package main

import "crypto/sha256"

type Block struct {
	//i. 前区块哈希
	PreHash []byte
	//ii. 当前区块哈希
	Hash []byte
	//iii. 数据
	Date []byte
}

// 2. 创建区块
func NewBlock(date string, preBlockHash []byte) *Block {
	block := Block{
		PreHash: preBlockHash,
		Hash:    []byte{}, //先填空，后面再计算
		Date:    []byte(date),
	}
	block.SetHash()
	return &block
}

// 3. ⽣成哈希
func (block *Block) SetHash() {

	//	1.拼装数据
	blockInfo := append(block.PreHash, block.Date...)
	//	2.shah256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
