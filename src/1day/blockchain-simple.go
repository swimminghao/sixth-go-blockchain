package main

import "fmt"

// 1. 定义结构
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
		Hash:    []byte{}, //先填空，后面再计算//TODO
		Date:    []byte(date),
	}

	return &block
}

//3. ⽣成哈希

//4. 引⼊区块链

//5. 添加区块
//6. 重构代码

func main() {
	block := NewBlock("老师转班长一枚比特币！", []byte{})
	fmt.Printf("前区块哈希值：%x\n", block.PreHash)
	fmt.Printf("当前区块哈希值：%x\n", block.Hash)
	fmt.Printf("区块数据：%s\n", block.Date)
	fmt.Println("hello")
}
