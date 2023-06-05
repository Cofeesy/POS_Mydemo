/**
  @author:BOEN
  @data:2023/6/5
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
	CoinAge     int
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
		strconv.Itoa(block.TimeCounter) + block.TradeData + block.PrevHash + strconv.Itoa(block.Difficulty) + strconv.Itoa(block.CoinAge)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

//
// generateBlock
//  @Description: 根据区块信息生成区块
//  @param prevBlock
//  @param coinPool
//  @param tradeData
//  @param difficulty
//  @param coinAge
//  @return Block
//
func generateBlock(prevBlock Block, coinPool *CoinPool, tradeData string, difficulty int, coinAge int) Block {
	newBlock := Block{
		Index:       prevBlock.Index + 1,
		Timestamp:   time.Now(),
		Stakeholder: "",
		TimeCounter: 0,
		TradeData:   tradeData,
		PrevHash:    prevBlock.Hash,
		Difficulty:  difficulty,
		CoinAge:     coinAge,
	}

	// 选择权益持有者作为共识节点
	stakeholder := selectStakeholder(coinPool)
	newBlock.Stakeholder = stakeholder

	// 从币池中移除权益持有者的代币
	coinPool.Stakeholders[stakeholder]--

	// 延长权益持有者的币龄
	extendCoinAge(coinPool)

	// 寻找满足条件的 timeCounter
	for {
		hash := calculateHash(newBlock)
		if isHashValid(hash, difficulty, coinAge) {
			newBlock.Hash = hash
			break
		}
		newBlock.TimeCounter++
	}

	return newBlock
}

//
// initialize
//  @Description:初始化币池和区块链
//  @param miners
//  @return *CoinPool
//  @return *Blockchain
//
func initialize(miners int) (*CoinPool, *Blockchain) {
	// 创建初始币池和区块链
	coinPool := &CoinPool{
		Stakeholders: make(map[string]int),
	}

	blockchain := &Blockchain{
		Blocks: []Block{},
	}

	// 初始化矿工数组和币池
	for i := 1; i <= miners; i++ {
		stakeholder := "Stakeholder" + strconv.Itoa(i)
		coinPool.Stakeholders[stakeholder] = 5 // 初始抵押数量为5
	}

	return coinPool, blockchain
}
