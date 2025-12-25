package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 定义⼀个⼯作量证明的结构ProofOfWork
type ProofOfWOrrk struct {
	//a. block
	block *Block
	//b. ⽬标值
	//一个非常大的数，有丰富的方法：比较，赋值方法
	target *big.Int
}

// 2. 提供创建POW的函数
//
// NewProofOfWork(参数)
func NewProofOfWork(block *Block) *ProofOfWOrrk {
	pow := ProofOfWOrrk{
		block: block,
	}
	//我们指定的难度值
	targetStr := "00000100000000000000000000000000000000000000000000000000000000000"
	bigInt := big.Int{}
	bigInt.SetString(targetStr, 16)
	pow.target = &bigInt
	return &pow
}

// 3. 提供计算不断计算hash的哈数
//
// Run()
func (pow *ProofOfWOrrk) Run() ([]byte, uint64) {

	var nonce uint64
	block := pow.block
	var hash [32]byte
	for {
		//1.拼装数据（区块的数据，还有不断变化的随机数）
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PreHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Date,
		}
		//将二维的切片数组连接起来，返回一个一维的切片
		blockInfo := bytes.Join(tmp, []byte{})
		//2.做hash运算
		hash = sha256.Sum256(blockInfo)
		//3.与pow中的target进行比较
		//将得到的hash数组，转换为bigInt
		tmpBig := big.Int{}
		tmpBig.SetBytes(hash[:])
		//比较当前hash值，小于则完成
		// Cmp compares x and y and returns:
		//
		//	-1 if x <  y
		//	 0 if x == y
		//	+1 if x >  y
		if tmpBig.Cmp(pow.target) == -1 {
			//	a.找到了，退出返回
			fmt.Printf("找到了，挖矿成功！hash：%x, nonce: %d \n", hash, nonce)
			break
		} else {
			//	b.没找到继续找，随机数+1
			nonce++
		}

	}
	return hash[:], nonce
}

//4. 提供⼀个校验函数
//
//IsValid()
