/**
  @author:王铮
  @data:2023/6/7
  @note:POS算法模拟实现
**/
package main

import (
	"fmt"
)

func main() {
	// 初始化矿工和币池
	fmt.Println("Initializing Miners and Coin Pool:")
	fmt.Print("Enter the number of miners: ")
	var numMiners int
	fmt.Scanln(&numMiners)

	coinPool := initializeMiners(numMiners)

	blockchain := &Blockchain{
		Blocks: []Block{},
	}

	// 打印并更新矿工数组信息
	printMinersInfo(coinPool)

	// 竞争挖矿并生成区块
	fmt.Println("Mining Phase and Generating New Coin and Block:")
	//这里其实应该不会有判断的,因为有创世区块的生成(这一步省略了,为了主要突出POS算法的逻辑)
	if len(blockchain.Blocks) > 0 {
		newBlock := generateBlock(blockchain.Blocks[len(blockchain.Blocks)-1], coinPool, "New Transaction Data", 2, 1)
		blockchain.Blocks = append(blockchain.Blocks, newBlock)
		fmt.Printf("Mining Successful! and New Block Generated! Block Hash: %s\n", newBlock.Hash)
		fmt.Printf("The winner of stakeholder: %s\n", newBlock.Stakeholder)
		blockchain.Blocks = append(blockchain.Blocks, newBlock)
		fmt.Println("------------------------------")
	} else {
		newBlock := generateBlock(Block{}, coinPool, "New Transaction Data", 2, 1)
		blockchain.Blocks = append(blockchain.Blocks, newBlock)
		fmt.Printf("Mining Successful! and New Block Generated! Block Hash: %s\n", newBlock.Hash)
		fmt.Printf("The winner of stakeholder: %s\n", newBlock.Stakeholder)
		blockchain.Blocks = append(blockchain.Blocks, newBlock)
		fmt.Println("------------------------------")
	}

	// 打印并更新矿工数组信息
	printMinersInfo(coinPool)

}

//如果还想要实验更加的复杂并更加贴合POS算法的本质原理,则:
//  1.设置difficulty为全局变量参与计算后,每次更新迭代
//  2.将Block数据结构改为链表形式,并有创世区块的生成
//  3.复杂化选出竞争者的逻辑,不单单是只看币数
