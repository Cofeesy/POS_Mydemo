/**
  @author:王铮
  @data:2023/6/7
  @note:区块相关操作
**/
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

//
//  Block
//  @Description: 区块数据结构
//
type Block struct {
	Index     int
	Timestamp time.Time
	//当一个矿工竞争成功的时候,记录该矿工
	Stakeholder string
	TimeCounter int
	TradeData   string
	PrevHash    string
	Hash        string
	//简化了运算，没有全局定义和全局初始化difficulty,于是初始使用的difficulty是零,而生成一个区块后,difficulty变为2
	Difficulty int
}

//
//  Blockchain
//  @Description: 区块链数据结构 => 简化的数组,没有使用链表的定义
//
type Blockchain struct {
	Blocks []Block
}

//
// calculateHash
//  @Description: 计算区块hash => 用于区块生成时生成block中的Hash和Prevhash字段
//  @param block
//  @return string
//
func calculateHash(block Block) string {
	data := strconv.Itoa(block.Index) + block.Timestamp.String() + block.Stakeholder +
		strconv.Itoa(block.TimeCounter) + block.TradeData + block.PrevHash + strconv.Itoa(block.Difficulty)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
