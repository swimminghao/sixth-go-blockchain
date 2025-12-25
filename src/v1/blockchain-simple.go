package main

import "fmt"

// import (
//
//	"crypto/sha256"
//	"fmt"
//
// )
//
// // 1. 定义结构
//
//	type Block struct {
//		//i. 前区块哈希
//		PreHash []byte
//		//ii. 当前区块哈希
//		Hash []byte
//		//iii. 数据
//		Date []byte
//	}
//
// // 2. 创建区块
//
//	func NewBlock(date string, preBlockHash []byte) *Block {
//		block := Block{
//			PreHash: preBlockHash,
//			Hash:    []byte{}, //先填空，后面再计算
//			Date:    []byte(date),
//		}
//		block.SetHash()
//		return &block
//	}
//
// // 3. ⽣成哈希
// func (block *Block) SetHash() {
//
//		//	1.拼装数据
//		blockInfo := append(block.PreHash, block.Date...)
//		//	2.shah256
//		hash := sha256.Sum256(blockInfo)
//		block.Hash = hash[:]
//	}
//
// // 4. 引⼊区块链
//
//	type BlockChain struct {
//		//定义一个区块链数组
//		blocks []*Block
//	}
//
// // 5. 定义一个区块链
//
//	func NewBLockChain() *BlockChain {
//		//创建一个创世块，并作为第一个区块添加到区块链中
//		genesisBlock := GenesisBlock()
//		return &BlockChain{
//			blocks: []*Block{genesisBlock},
//		}
//	}
//
// // 创世块
//
//	func GenesisBlock() *Block {
//		return NewBlock("创世块!!!!", []byte{})
//	}
//
// //5. 添加区块
//
//	func (bc *BlockChain) AddBlock(data string) {
//		//获取前区块的哈希
//		lastBlock := bc.blocks[len(bc.blocks)-1]
//		preHash := lastBlock.Hash
//		//	a.创建新的区块
//		block := NewBlock(data, preHash)
//		//	b.添加到区块链数组中
//		bc.blocks = append(bc.blocks, block)
//	}
//
// //6. 重构代码
func main() {
	bc := NewBlockChain()
	bc.AddBlock("班长向班花转了50枚比特币！")
	bc.AddBlock("班长又向班花转了50枚比特币！")
	bc.AddBlock("班长又向班花转了40枚比特币！")
	for i, block := range bc.blocks {
		fmt.Printf("=====当前区块高度：%d\n", i)
		fmt.Printf("前区块哈希值：%x\n", block.PreHash)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("区块数据：%s\n", block.Date)
	}
}
