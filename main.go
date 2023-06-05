/**
  @author:Cofeesy
  @data:2023/6/5
  @note:POS算法模拟实现
**/
package main

import (
	"fmt"
)

func main() {
	// 初始化阶段
	miners := 3
	coinPool, blockchain := initialize(miners)

	// 竞争挖矿阶段
	fmt.Println("Mining Phase:")
	newBlock := generateBlock(Block{}, coinPool, "Transaction Data", 2, 1)
	fmt.Printf("Mining Successful! Block Hash: %s\n", newBlock.Hash)
	fmt.Println()

	// 打印并更新矿工数组信息
	printMinersInfo(coinPool)

	// 生成新币和新区块
	fmt.Println("Generating New Coin and Block:")
	newBlock = generateBlock(blockchain.Blocks[len(blockchain.Blocks)-1], coinPool, "New Transaction Data", 2, 1)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
	fmt.Printf("New Block Generated! Block Hash: %s\n", newBlock.Hash)
	fmt.Println()

	// 打印并更新矿工数组信息
	printMinersInfo(coinPool)
}
