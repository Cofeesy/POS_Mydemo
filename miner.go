/**
  @author:王铮
  @data:2023/6/7
  @note:矿工相关操作
**/
package main

import (
	"fmt"
	"strconv"
	"time"
)

//
//  Stakeholder
//  @Description:矿工数据结构
//
type Stakeholder struct {
	Address   string
	CoinAge   int
	CoinCount int
}

//
//  CoinPool
//  @Description: 币池数据结构
//
type CoinPool struct {
	Stakeholders map[string]*Stakeholder
}

//
// printMinersInfo
//  @Description: 打印矿工的信息
//  @param coinPool
//
func printMinersInfo(coinPool *CoinPool) {
	// 打印矿工数组信息
	fmt.Println("Miners Information:")
	for index, stakeholder := range coinPool.Stakeholders {
		fmt.Printf("Stakeholder: %s, tokenAmout: %d\n,coinAge: %d\n", index, stakeholder.CoinCount, stakeholder.CoinAge)
	}
	fmt.Println()
}

//
// selectStakeholder
//  @Description: 选择币池中的权益者 => POS的具体简易实现
//  @param coinPool
//  @return string
//
func selectStakeholder(coinPool *CoinPool) string {
	// 根据权益持有者的抵押数量选择共识节点
	var maxTokens int
	var maxStakeholder string
	for index, stakeholder := range coinPool.Stakeholders {
		if stakeholder.CoinCount > maxTokens {
			maxTokens = stakeholder.CoinCount
			maxStakeholder = index
		}
	}
	return maxStakeholder
}

//
// extendCoinAge
//  @Description:延长币龄
//  @param coinPool
//
func extendCoinAge(coinPool *CoinPool) {
	// 延长权益持有者的币龄，每个区块生成后币龄加1
	for stakeholder := range coinPool.Stakeholders {
		coinPool.Stakeholders[stakeholder].CoinAge++
	}
}

//
// isHashValid
//  @Description:验证哈希是否满足条件
//  @param hash
//  @param difficulty
//  @param coinAge
//  @return bool
//
func isHashValid(hash string, difficulty int, coinAge int) bool {
	// 简易实现：SHA256(SHA256(tradeData|timeCounter))<D×coinAge
	target := difficulty * coinAge
	return hash < strconv.Itoa(target)
}

func (coinPool *CoinPool) resetToZero(stakeholder string) bool {
	coinPool.Stakeholders[stakeholder].CoinAge = 0
	return true
}

//
// initialize
//  @Description:初始化币池和区块链
//  @param miners
//  @return *CoinPool
//  @return *Blockchain
//
func initializeMiners(numMiners int) *CoinPool {
	coinPool := &CoinPool{
		Stakeholders: make(map[string]*Stakeholder),
	}

	for i := 1; i <= numMiners; i++ {
		address := fmt.Sprintf("Miner%d", i)
		coinPool.Stakeholders[address] = &Stakeholder{
			Address:   address,
			CoinAge:   5,
			CoinCount: 5,
		}
	}

	return coinPool
}

//
// generateBlock
//  @Description: 根据区块信息生成新区块
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
	}

	// 选择权益持有者作为共识节点
	stakeholder := selectStakeholder(coinPool)
	//fmt.Println("竞争胜利者：", stakeholder)
	newBlock.Stakeholder = stakeholder

	// 延长权益持有者的币龄
	extendCoinAge(coinPool)

	//获胜者的币龄清零
	coinPool.resetToZero(stakeholder)

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
