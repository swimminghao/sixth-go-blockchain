package main

// 4. 引⼊区块链
type BlockChain struct {
	//定义一个区块链数组
	blocks []*Block
}

// 5. 定义一个区块链
func NewBLockChain() *BlockChain {
	//创建一个创世块，并作为第一个区块添加到区块链中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

// 创世块
func GenesisBlock() *Block {
	return NewBlock("创世块!!!!", []byte{})
}

//5. 添加区块

func (bc *BlockChain) AddBlock(data string) {
	//获取前区块的哈希
	lastBlock := bc.blocks[len(bc.blocks)-1]
	preHash := lastBlock.Hash
	//	a.创建新的区块
	block := NewBlock(data, preHash)
	//	b.添加到区块链数组中
	bc.blocks = append(bc.blocks, block)
}

//6. 重构代码
