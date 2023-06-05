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
	Index       int
	Timestamp   time.Time
	Stakeholder string
	TimeCounter int
	TradeData   string
	PrevHash    string
	Hash        string
	Difficulty  int
}

//
//  Blockchain
//  @Description: 区块链数据结构
//
type Blockchain struct {
	Blocks []Block
}

//
// calculateHash
//  @Description: 计算区块hash
//  @param block
//  @return string
//
func calculateHash(block Block) string {
	data := strconv.Itoa(block.Index) + block.Timestamp.String() + block.Stakeholder +
		strconv.Itoa(block.TimeCounter) + block.TradeData + block.PrevHash + strconv.Itoa(block.Difficulty)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
