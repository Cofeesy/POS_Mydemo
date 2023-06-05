/**
  @author:BOEN
  @data:2023/6/5
  @note:矿工相关操作
**/
package main

import (
	"fmt"
	"strconv"
)

//
//  CoinPool
//  @Description: 币池数据结构
//
type CoinPool struct {
	Stakeholders map[string]int
}

//
// printMinersInfo
//  @Description: 打印矿工的信息
//  @param coinPool
//
func printMinersInfo(coinPool *CoinPool) {
	// 打印矿工数组信息
	fmt.Println("Miners Information:")
	for stakeholder, tokens := range coinPool.Stakeholders {
		fmt.Printf("Stakeholder: %s, Tokens: %d\n", stakeholder, tokens)
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
	for stakeholder, tokens := range coinPool.Stakeholders {
		if tokens > maxTokens {
			maxTokens = tokens
			maxStakeholder = stakeholder
		}
	}
	return maxStakeholder
}

//
// extendCoinAge
//  @Description:
//  @param coinPool
//
func extendCoinAge(coinPool *CoinPool) {
	// 延长权益持有者的币龄，每个区块生成后币龄加1
	for stakeholder := range coinPool.Stakeholders {
		coinPool.Stakeholders[stakeholder]++
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
